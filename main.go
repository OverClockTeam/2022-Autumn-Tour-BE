package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
)

func FileUpload(c *gin.Context) {
	subject := c.PostForm("subject")
	newDir := "./Homework_Collector/" + subject
	_, err := os.Stat(newDir) //返回文件属性
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"msg": "stat newDir error,maybe is not exit",
		})
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotImplemented, gin.H{
				"msg": "newDir is not exit",
			})
			err := os.Mkdir(newDir, 0666)
			if err != nil {
				c.JSON(http.StatusNotImplemented, gin.H{
					"msg": "make new dir failed",
				})
			}
		}
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"msg": "Upload failed!",
		})
		return
	}
	dst := "./Homework_Collector/" + subject + "/" + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"msg": "save file failed!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully!",
	})
}

func EmailRemainder(c *gin.Context) {
	sendAddress := c.PostForm("sendAddress")
	e := email.NewEmail()
	e.From = "Reminder <3143199752@qq.com>"
	e.To = []string{sendAddress}
	e.Subject = "homework reminder"
	e.Text = []byte("该交作业啦！！！")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "3143199752@qq.com", "reyykuddrtvfdcce", "smtp.qq.com"))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "send successfully",
	})
}

func InitRouter(r *gin.Engine) {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: false,
	}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	api := r.Group("/main")
	{
		user := api.Group("/user")
		user.Use(cors.Default())
		{
			user.POST("/FileUpload", FileUpload)
			user.POST("/EmailRemainder", EmailRemainder)
		}
	}
}

func main() {
	r := gin.Default()
	InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
