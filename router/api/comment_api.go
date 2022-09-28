package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostComment(c *gin.Context) {
	postID := c.Param("post")
	var p model.Post
	err := client.DB.Where("id = ?", postID).First(&p).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	p.NewComment(c.PostForm("author"), c.PostForm("content"))
	client.DB.Save(&p)
	c.JSON(http.StatusOK, gin.H{
		"message": "comment post successfully",
	})
}

func GetComment(c *gin.Context) {
	postID := c.Param("post")
	var p model.Post
	err := client.DB.Where("id = ?", postID).First(&p).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  c.Request.URL.Path,
		"comments": p.Comment,
	})
	//c.HTML(http.StatusOK,"comment.html",nil)
}
