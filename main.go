package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//http://localhost:8080/upload -F "file=@/Users/appleboy/test.zip" \
//-H "Content-Type: multipart/form-data"

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 4 << 20 // 4 MiB
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())//400
			return
		}

		filename := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())//400
			return
		}

		c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)//200
	})
	router.Run(":8080")
}
