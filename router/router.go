package router

import (
	"HH_LHY/middleware"
	"HH_LHY/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	auth := g.Group("/auth")
	{
		auth.GET("/login", middleware.JWTAuthMiddleHandler(), api.GetLogin)          //ok
		auth.POST("/login", api.PostLogin)                                           //ok
		auth.GET("/register", api.GetRegister)                                       //ok
		auth.POST("/register", middleware.RegisterMiddleHandler(), api.PostRegister) //ok
		auth.GET("/:email/verify", api.GetVerify)
		auth.POST("/:email/verify", api.PostVerify)
	}
	posts := g.Group("/post")
	{
		posts.GET("/index", api.GetPostsIndex)                                   //ok
		posts.GET("/publish", middleware.JWTAuthMiddleHandler(), api.GetPublish) //ok
		posts.POST("/publish", middleware.AuditPost(), api.PostPublish)          //ok
		posts.GET("/:tag", api.GetTagList)                                       //ok
		posts.POST("/comment/:post", middleware.AuditPost(), api.PostComment)    //ok
		posts.GET("/comment/:post", api.GetComment)                              //ok
		posts.GET("/detail", api.GetPostsDetail)                                 //ok
		posts.GET("/search", api.SearchPost)                                     //ok
	}
	users := g.Group("/user")
	{
		users.GET("/:name", api.GetUserIndex) //ok
	}
}
