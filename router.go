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
	//router.GET("/getMusic", musicQuery)
	router.Run()
}
