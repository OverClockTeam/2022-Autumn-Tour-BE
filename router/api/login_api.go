package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"HH_LHY/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetLogin(c *gin.Context) {
	value, exists := c.Get("user")
	u, ok := value.(model.User)
	if exists && ok && !u.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome" + u.Name,
		})
		return
	} else {
		//c.HTML(http.StatusOK,"login.html",nil)
		c.JSON(http.StatusOK, gin.H{
			"message": c.Request.URL.Path,
		})
	}
}

func PostLogin(c *gin.Context) {
	//var u = model.User{
	//	Name:     c.PostForm("name"),
	//	Password: c.PostForm("password"),
	//}
	//if client.DB.First(&u).Error != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "用户不存在",
	//	})
	//	return
	//}
	var u model.User
	name := c.PostForm("username")
	err := client.DB.Where("name = ?", name).First(&u).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	password := c.PostForm("password")
	if u.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"message": "wrong password",
		})
		return
	}
	token, err := util.GenerateToken(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "欢迎" + u.Name,
		"token":   token,
	})
	return
}
