package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func musicQuery(c *gin.Context) {
	musicName := c.Query("name")
	searchMusic(musicName, "1")

}

func lyricsQuery(c *gin.Context) {
	name := c.Request.FormValue("Name")
	musiclist := searchMusic(name, "1")
	lyc := downloadlyrics(musiclist[0])
	c.String(http.StatusOK, "%v", lyc)
}
