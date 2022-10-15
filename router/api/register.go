package api

import (
	"fmt"
	"net/http"
	"strconv"
	"HC_WJ/dbclt"
	"HC_WJ/model"
	"github.com/gin-gonic/gin"
)

//注册页面
func Register(c *gin.Context) {
	//处理提交注册信息后的情况
	var u model.User
	var err error
	u.Username = c.PostForm("username")
	u.Password = c.PostForm("password1")          
	u.Email = c.PostForm("email")    
	u.Class, _ = strconv.Atoi(c.PostForm("class"))  
	u.Monitor, _ = strconv.Atoi(c.PostForm("monitor"))
	password1 := c.PostForm("password2")

	//重复密码不相等时重新进入网页
	if u.Password != password1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "重复密码不相同",
		})
		return 
	}

	//用户名存在时重新进入页面
	s := "select username from users where username = ?"
	err = dbclt.Db.QueryRow(s, u.Username).Scan(&u.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "用户名已存在",
		})
		return
	}

	//将数据上传到数据库
	s = "insert into users (username, password, email, class, monitor) values(?,?,?,?,?)"
	r, err := dbclt.Db.Exec(s, u.Username, u.Password, u.Email, u.Class, u.Monitor)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
	//提示注册成功
	c.JSON(http.StatusOK, gin.H{
		"message" : "注册成功",
	})
}
