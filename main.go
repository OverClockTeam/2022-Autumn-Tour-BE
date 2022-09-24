package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// url: 127.0.0.1:8080/main/user/postFile/?subject=物理
func PostFile(c *gin.Context) {
	subject := c.PostForm("subject")
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
		}
	}
}
func main() {
	r := gin.Default()
	InitRouter(r)
	r.Run(":8080")
}
