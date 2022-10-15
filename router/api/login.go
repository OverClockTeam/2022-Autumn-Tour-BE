package api

import (
	"log"
	"HC_WJ/model"
	"HC_WJ/dbclt"
	"HC_WJ/util"
	"net/http"
	"github.com/gin-gonic/gin"
)

//登陆页面
func Login(c *gin.Context) {

	//接收数据
	var u model.User
	var Password_t string
	u.Username = c.PostForm("username")
	Password_t = c.PostForm("password")
	
	//判断是否存在账号
	s := "select * from users where username = ?"
	err := dbclt.Db.QueryRow(s, u.Username).Scan(&u.Username, &u.Password, &u.Email, &u.Class, &u.Monitor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "用户名不存在",
		})
	} else if u.Password != Password_t {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "密码错误",
		})
	} else {
		token, err := util.GenerateToken(u)
		if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "欢迎" + u.Username,
			"token" : token,
		})
	}

}