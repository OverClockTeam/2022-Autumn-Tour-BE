package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//定义全局数据
var db *sql.DB
var f *gin.Engine
var u User 

//定义结构体
type User struct {
	Status bool
	Username string
	Password string
	Email string
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
	//上传文件到指定目录
	dst := fmt.Sprintf("./subject_file/%s/%s", subject, file.Filename)
	c.SaveUploadedFile(file, dst)
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/",
		"Context" : "上传文件成功！请稍后",
	})
	
}

//注册页面
func Register(c *gin.Context) {
	//渲染注册页面模版
	c.HTML(http.StatusOK, "html/register.html", nil)

	//处理提交注册信息后的情况
	f.POST("/register/get", func(c *gin.Context) {
		//获取上传的信息
		var u User
		u.Username = c.PostForm("username")
		u.Password = c.PostForm("password1")
		u.Email = c.PostForm("email")
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
		err := db.QueryRow(s, u.Username).Scan(&u.Username, &u.Username, &u.Username)
		if err == nil {
			c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
				"Url" : "/",
				"Context" : "用户名已存在",
			})
			return
		}

		//将数据上传到数据库
		s = "insert into users (username, password, email) values(?,?,?)"
		r, err := db.Exec(s, u.Username, u.Password, u.Email)
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

func Login(c *gin.Context) {

	//接收数据
	u.Username = c.PostForm("username")
	u.Password = c.PostForm("password")
	
	//判断是否存在账号
	s := "select * from users where username = ? and password = ?"
	err := db.QueryRow(s, u.Username, u.Password).Scan(&u.Username, &u.Password, &u.Email)
	if err != nil {
		u.Status = false
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/",
			"Context" : "用户名或密码不正确",
		})
	} else {
		u.Status = true
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/index",
			"Context" : "登录成功！请稍后",
		})
	}
}

func Index(c *gin.Context) {
	if u.Status {
		c.HTML(http.StatusOK, "html/index.tmpl", u.Username)
	} else {
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/",
			"Context" : "请先登录！",
		})
	}
}

func Logout(c *gin.Context) {
	u.Status = false
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/",
		"Context" : "正在登出",
	})
}

func main() {
	err := InitDB()	//连接数据库
	if err != nil {
		panic(err)
	}
	u.Status = false //初始化用户状态

	//获取路由对象
	f = gin.Default()

	//解析模板
	f.LoadHTMLGlob("./html/*")

	//根目录渲染登录模板
	f.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html/login.html", nil)
	})

	//处理点击登录后的情况
	f.POST("/login", Login)

	//处理点击注册后的情况
	f.POST("/register", Register)

	//接受文件
	f.POST("/upload", Upload)
	
	//登录成功后的主界面
	f.Any("/index", Index)

	//退出登录
	f.POST("/logout", Logout)

	f.Run()
}
