package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"HH_LHY/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchPost(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		//c.HTML(http.StatusOK,"search.html",nil)
		c.JSON(http.StatusOK, gin.H{
			"message": c.Request.URL.Path,
			"result":  "you search nothing",
		})
		return
	}
	var strs []string = services.CutStr(q)
	var post []model.Post
	var relatedPosts []model.Post
	client.DB.Find(&post)
	for _, m := range post {
		if services.IsInSearchList(m.Title, strs) {
			relatedPosts = append(relatedPosts, m)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"relatedPosts": relatedPosts,
	})
}
