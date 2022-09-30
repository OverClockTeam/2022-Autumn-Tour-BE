package main

import "github.com/gin-gonic/gin"

func routerInit() {
	router := gin.Default()
	v1 := router.Group("admin")
	{
		v1.GET("/register", register)
		v1.GET("/login", login)
	}
	v2 := router.Group("music")
	{
		v2.GET("/getmusic", musicQuery)
		v2.GET("/getlyrics", lyricsQuery)
	}
	v3 := router.Group("list")
	{
		v3.POST("musiclist")
		v3.GET("searchlist", searchQuery)
	}
	router.Run()
}
