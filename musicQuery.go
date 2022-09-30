package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

func musicQuery(c *gin.Context) {
	musicName := c.Request.FormValue("Name")
	musiclist := searchMusic(musicName, "1")
	if musiclist == nil {
		c.String(http.StatusOK, "音乐查询失败")
	}
	{
		_, fpath := downLoad(musiclist[0])
		fileContent, err := os.ReadFile(fpath)
		if err != nil {
			c.String(http.StatusOK, "文件打开失败")
		}
		c.Header("Content-Type", "audio/mp3")
		c.Header("Content-Disposition", fpath)
		c.Data(http.StatusOK, "mp3", fileContent)
	}

}

func lyricsQuery(c *gin.Context) {
	name := c.Request.FormValue("Name")
	musiclist := searchMusic(name, "1")
	if musiclist == nil {
		c.String(http.StatusOK, "音乐查询失败")
	}
	{
		lyc := downloadlyrics(musiclist[0])
		c.String(http.StatusOK, "%v", lyc)
	}
}

func searchQuery(c *gin.Context) {
	searchName := c.Request.FormValue("Name")
	for i := 1; i < 10; i++ {
		musiclist := searchMusic(searchName, strconv.Itoa(i))
		if musiclist == nil {
			c.String(http.StatusOK, "音乐查询失败")
		}
		{
			for _, value := range musiclist {
				err, _ := downLoad(value)
				if err != nil {
					fmt.Println(err.Error())
				}
				_ = downloadlyrics(value)
				c.String(http.StatusOK, value.name+"下载已完成")
			}
			log.Println("Page" + strconv.Itoa(i) + "完成")
		}
	}

}
