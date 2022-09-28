package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPublish(c *gin.Context) {
	value, exists := c.Get("user")
	u, ok := value.(model.User)
	if exists && ok && !u.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"message": c.Request.URL.Path,
			"user":    u,
		})
		//c.HTML(http.StatusOK, "publish.html", nil)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/api/auth/login")
}

func PostPublish(c *gin.Context) {
	post := &model.Post{
		Title:   c.PostForm("title"),
		Content: c.PostForm("content"),
		Author:  c.PostForm("author"),
		Tag:     c.PostForm("tag"),
	}
	if client.DB.Create(post).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "post failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "post succeed",
			"info":    post,
		})
	}
}
