package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
)

//定义全局数据
var db *sql.DB
var f *gin.Engine
var u User 
var users [] Homework
var len int

//定义结构体
type User struct {
	Status int	//0是未登录，1是普通用户，2是班长
	Username string
	Password string
	Class int
	Email string
	Share string
}

type Homework struct {
	Username string
	Subject string
	Email string
}


//定义邮箱
func SendMail(mailTo []string, subject string, body string) error {
 
    mailConn := map[string]string{
        "user": "wj13985781016@163.com",
        "pass": "NPGZEFRFHVRMRUHO",
        "host": "smtp.163.com",
        "port": "465",
    }
 
    port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
 
    m := gomail.NewMessage()
 
    m.SetHeader("From",  m.FormatAddress(mailConn["user"], "班长")) //这种方式可以添加别名，即“XX官方”
    m.SetHeader("To", mailTo...)    //发送给多个用户
    m.SetHeader("Subject", subject) //设置邮件主题
    m.SetBody("text/html", body)    //设置邮件正文
 
    d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
 
    err := d.DialAndSend(m)
    return err
}

//发送邮件
func Send(c *gin.Context)() {
	//定义收件人
	var mailTo []string
	for i := 0; i < len; i++ {
		mailTo = append(mailTo, users[i].Email)
	}
	//邮件主题为"Hello"
	subject := "交作业"
	// 邮件正文
	body := "请及时交" + subject + "作业！" + "来自" + u.Username + "班长"
 
	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}
 
	fmt.Println("send successfully")
}
//连接数据库
func InitDB() (err error) {
	dsn := "root:wj2581320495@tcp(127.0.0.1:3306)/user"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//产生随机数
func Randstr() (string) {
	var str string
	for i := 0; i < 18; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		str += string(rune(rand.Intn(74) + 48))
	}
	return str
}
//接收文件
func Upload(c *gin.Context) {
	//接受文件
	file, err := c.FormFile("file")
	subject := c.PostForm("subject")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)

	//上传文件到指定目录,格式为 班级_用户名_文件名
	file_name := strconv.Itoa(u.Class) + "_" + u.Username + "_" + file.Filename
	dst := fmt.Sprintf("./subject_file/%s/%s", subject, file_name)
	c.SaveUploadedFile(file, dst)
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/login/index",
		"Context" : "上传文件成功！请稍后",
	})

	//登记上传信息
	u.Share = Randstr()
	//清除之前上传的信息,再上传新信息
	s := "delete from homework where username = ? and subject = ?"
	r, _ := db.Exec(s, u.Username, subject)
	fmt.Printf("r: %v\n", r);
	s = "insert into homework (username, subject, share, file_name, class) values(?,?,?,?,?)"
	r, _ = db.Exec(s, u.Username, subject, u.Share, file_name, u.Class)
	fmt.Printf("r: %v\n", r);
}

//注册页面
func Register(c *gin.Context) {
	//渲染注册页面模版
	c.HTML(http.StatusOK, "html/register.html", nil)

	//处理提交注册信息后的情况
	f.POST("/register/get", func(c *gin.Context) {
		//获取上传的信息
		var u User
		var err error
		u.Username = c.PostForm("username")
		u.Password = c.PostForm("password1")
		u.Email = c.PostForm("email")
		u.Class, _ = strconv.Atoi(c.PostForm("class"))
		password1 := c.PostForm("password2")

		//重复密码不相等时重新进入网页
		if u.Password != password1 {
			c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
				"Url" : "/",
				"Context" : "重复密码不相同",
			})
			return
		}

		//用户名存在时重新进入页面
		s := "select * from users where username = ?"
		err = db.QueryRow(s, u.Username).Scan(&u.Username, &u.Username, &u.Username)
		if err == nil {
			c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
				"Url" : "/",
				"Context" : "用户名已存在",
			})
			return
		}

		//将数据上传到数据库
		u.Status, _ = strconv.Atoi(c.PostForm("monitor"))
		s = "insert into users (username, password, email, class, monitor) values(?,?,?,?,?)"
		r, err := db.Exec(s, u.Username, u.Password, u.Email, u.Class, u.Status)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		} else {
			i, _ := r.LastInsertId()
			fmt.Printf("i: %v\n", i)
		}

		//提示注册成功
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/",
			"Context" : "注册成功，请稍后",
		})
	})
}

//登陆页面
func Login(c *gin.Context) {

	//接收数据
	u.Username = c.PostForm("username")
	u.Password = c.PostForm("password")
	u.Share = ""
	
	//判断是否存在账号
	s := "select * from users where username = ? and password = ?"
	err := db.QueryRow(s, u.Username, u.Password).Scan(&u.Username, &u.Password, &u.Email, &u.Class, &u.Status)
	if err != nil {
		u.Status = 0
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/",
			"Context" : "用户名或密码不正确",
		})
	} else {
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/login/index",
			"Context" : "登录成功！请稍后",
		})
	}
}

//检查作业
func Check(c *gin.Context) {
	//接受数据
	subject := c.PostForm("subject")

	//找出自己班级没交某个作业的人并输出
	s := "select users.username, users.email, homework.username from users left join homework on users.username = homework.username and users.class = ? and homework.subject = ?"
	r, _ := db.Query(s, u.Class, subject)
	users = []Homework{}
	len = 0
	var u1 Homework
	for r.Next() {
		r.Scan(&u1.Username, &u1.Email, &u1.Subject)
		//若没有提交则左右表名字不匹配
		if u1.Subject != u1.Username  {
			u1.Subject = subject
			users = append(users, u1)
			len++
		}
	}
	
	//输出到html页面
	c.HTML(http.StatusOK, "html/index.tmpl", gin.H{
		"res": users,
		"name" : u.Username + "班长",
		"monitor" : 1,
	})

}

//下载作业

func Download(c *gin.Context) {
	var subject, file_name string
	share := c.PostForm("share")
	s := "select file_name, subject from homework where share = ?"
	err := db.QueryRow(s, share).Scan(&file_name, &subject)
	if err != nil {
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/login/index",
			"Context" : "没有找到分享码！",
		})
		return 
	}
	c.File("./subject_file/" + subject + "/" + file_name)
}
//主界面渲染
func Index(c *gin.Context) {
	//阻止没有登录的访问,并分别渲染普通用户和班长的页面
	switch u.Status {
	case 0:
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/",
			"Context" : "请先登录！",
		})
	case 1: //普通用户
		c.HTML(http.StatusOK, "html/index.tmpl", gin.H{
			"name" : u.Username,
			"share" : u.Share,
			"monitor" : 0,
		})
	case 2:	//班长
		//导出没有交作业的同班同学并制作成表格
		c.HTML(http.StatusOK, "html/index.tmpl", gin.H{
			"name" : u.Username + "班长",
			"share" : u.Share,
			"monitor" : 1,
			"res" : nil,
		})
	}

}

//登出
func Logout(c *gin.Context) {
	u.Status = 0
	u.Share = ""
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/",
		"Context" : "正在登出",
	})
}

//清空本班的作业表
func Endput (c *gin.Context) {
	s := "delete from homework where class = ?"
	r, _ := db.Exec(s, u.Class)
	fmt.Printf("r: %v\n", r)
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/login/index",
		"Context" : "已清空本班作业提交记录",
	})
}
func main() {
	err := InitDB()	//连接数据库
	if err != nil {
		panic(err)
	}
	u.Status = 0 //初始化用户状态

	//获取路由对象
	f = gin.Default()

	//解析模板
	f.LoadHTMLGlob("./html/*")

	//根目录渲染登录模板
	f.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html/login.html", nil)
	})

	//处理点击注册后的情况
	f.POST("/register", Register)

	//处理点击登录后的情况
	f.POST("/login", Login)

	//登录成功后的主界面
	f.Any("/login/index", Index)


	//下载作业
	f.POST("/login/index/download", Download)

	//接受作业
	f.POST("/login/index/upload", Upload)

	//检查作业
	f.POST("/login/index/check", Check)

	//退出登录
	f.POST("/login/index/logout", Logout)

	//发送邮件 
	f.POST("/login/index/check/sendmail", Send)

	//结束作业提交，清空本班的作业表
	f.POST("/login/index/check/endput", Endput)


	f.Run()
}
