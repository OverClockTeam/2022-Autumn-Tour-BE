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

//定义结构体
type User struct {
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
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
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
			c.HTML(http.StatusOK, "html/register_fail.html", nil)
		}

		//将数据上传到数据库

		s := "insert into users (username, password, email) values(?,?,?)"
		r, err := db.Exec(s, u.Username, u.Password, u.Email)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		} else {
			i, _ := r.LastInsertId()
			fmt.Printf("i: %v\n", i)
		}

		//提示注册成功
		c.HTML(http.StatusOK, "html/register_success.html", nil)
	})
}

func Login(c *gin.Context) {

	//接收数据
	var u User
	u.Username = c.PostForm("username")
	u.Password = c.PostForm("password")
	
	//判断是否存在账号
	s := "select * from users where username = ? and password = ?"
	err := db.QueryRow(s, u.Username, u.Password).Scan(&u.Username, &u.Password, &u.Email)
	if err != nil {
		c.HTML(http.StatusOK, "html/login_fail.html", nil)
	} else {
		c.HTML(http.StatusOK, "html/index.tmpl", u.Username)
	}
}

func main() {
	//连接数据库
	err := InitDB()
	if err != nil {
		panic(err)
	}

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
	f.Run()
}
