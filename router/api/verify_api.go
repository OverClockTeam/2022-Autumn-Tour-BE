package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"HH_LHY/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetVerify(c *gin.Context) {
	email := c.Param("email")
	var u model.User
	result := client.DB.Where("email = ?", email)
	if err := result.Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	result.First(&u)
	if u.IsActivated {
		c.JSON(http.StatusOK, gin.H{
			"message": "user is already activated",
		})
		return
	}
	code := services.SendEmail(u.Email)
	client.DB.Model(&u).Update("verify_code", code)
	c.JSON(http.StatusOK, gin.H{
		"message":     c.Request.URL.Path,
		"verify code": code,
	})
	//c.HTML(http.StatusOK,"verify.html",nil)
}

func PostVerify(c *gin.Context) {
	code := c.PostForm("code")
	email := c.Param("email")
	var u model.User
	result := client.DB.Where("email = ?", email)
	if err := result.Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	result.First(&u)
	if u.VerifyCode == code {
		u.IsActivated = true
		client.DB.Save(&u)
		c.JSON(http.StatusOK, gin.H{
			"message": "user is activated",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "wrong code",
		})
	}
}
