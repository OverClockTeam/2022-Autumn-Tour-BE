package main

import (
	"fmt"
    	"io/ioutil"
    	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	uploadFileKey = "upload-key"
)

func main() {
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	r.Run()
}
 resp, err := http.Get("")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }

func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	if err != nil {
		//ignore
	}
	dst := header.Filename
  // gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		// ignore
	}
}

//基本的GET请求
 



