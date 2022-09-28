package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"HH_LHY/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.URL.Path,
	})
	c.Redirect(http.StatusTemporaryRedirect, "/verify")
	//c.HTML(http.StatusOK, "register.html", nil)
}

func PostRegister(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if name == "" || password == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong arguments",
		})
		return
	}
	user := model.User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	err := client.DB.Create(&user).Error //创建成功将数据写入数据库
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "the email has been registered",
		})
		return
	}
	token, err := util.GenerateToken(user)
	if err != nil {
		log.Println(err)
		return
	}
	// 将生成的token返回给客户端
	c.JSON(http.StatusOK, gin.H{
		"message": "register succeeded",
		"token":   token,
	})
	return
}
