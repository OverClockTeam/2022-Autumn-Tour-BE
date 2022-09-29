package main

import (
	LYXemail "LYX/email"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

/* url: 127.0.0.1:8080/main/user/postFile/?subject=物理 */

func PostFile(c *gin.Context) {
	subject := c.Query("subject")
	//判断是否已经存在这个科目
	//如果是 直接存
	//如果不是 创建这个目录
	if subject == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "no subject",
		})
		return
	}

	err := os.Mkdir("./2022-Autumn-Tour-BE/upload/"+subject, 0666)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"msg": "err",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed send,try again!",
		})
		return
	}
	basePath := "./2022-Autumn-Tour-BE/upload/"
	savePath := basePath + subject + "/" + filepath.Base(file.Filename)
	//   upload/huaxue
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Success!",
	})
}

/*
命名规范问题
蛇形
大驼峰
小驼峰
下划线

函数命名  首字母小写 那只能在私有包下引用  首字母大写则可以在包外引用   类似于面对对象语言中的私有类  私有方法之类的设计
变量命名  小驼峰 美观     emailAddress   bool: isPasswordUsernameRight
*/

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "HEllO GO",
	})
	return
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
			user.GET("/hello", HelloWorld)
			user.POST("/postFile", PostFile)
			user.POST("/reminder", LYXemail.Reminder)
		}
	}
}

func IdLink(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	r := gin.Default()
	InitRouter(r)

	r.GET("/main/user/link", func(c *gin.Context) {
		id := c.Query("id")
		pswd := IdLink(id)
		c.JSON(http.StatusOK, gin.H{
			"link": "localhost:8080/main/user/" + pswd,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
