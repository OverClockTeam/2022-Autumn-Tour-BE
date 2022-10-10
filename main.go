package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

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

func main() {
	f := gin.Default()
	//解析模板
	f.LoadHTMLGlob("./index.html")
	//在根目录渲染模板
	f.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./index.html", nil)
	})
	//接受文件
	f.POST("/upload", Upload)
	f.Run()
}
