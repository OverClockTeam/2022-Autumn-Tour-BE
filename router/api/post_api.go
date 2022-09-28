package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetPostsIndex(c *gin.Context) {
	var post []model.Post
	err := client.DB.Find(&post).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.URL.Path,
		"posts":   post,
	})
	//c.HTML(http.StatusOK, "index.html", nil)
}

func GetTagList(c *gin.Context) {
	tag := c.Param("tag")
	var posts []model.Post
	err := client.DB.Where("tag = ?", tag).Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.URL.Path,
		"tag":     tag,
		"posts":   posts,
	})
}

func GetPostsDetail(c *gin.Context) {
	//c.HTML(http.StatusOK,"post.html",nil)
	postID := c.Query("postID")
	var p model.Post
	result := client.DB.Where("id = ?", postID).First(&p)
	err := result.Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	//if p.Title == "" && p.Content == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "no such post",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"detail": p,
	})
}
