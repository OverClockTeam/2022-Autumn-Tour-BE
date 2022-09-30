package router

import (
	v1 "OverClock/api/v1"
	"OverClock/middleware"
	"OverClock/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JWTToken())
	{

		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		auth.POST("category/add",v1.AddCate)
		auth.PUT("category/:id",v1.EditCate)
		auth.DELETE("category/:id",v1.DeleteCate)
		auth.POST("article/add",v1.AddArticle)
		auth.PUT("article/:id",v1.EditArticle)
		auth.DELETE("article/:id",v1.DeleteArticle)
	}
	router := r.Group("api/v1")
	{
		//用户模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		//分类模块的路由接口
		router.GET("category",v1.GetCate)
		//帖子模块的路由接口
		router.GET("article",v1.GetArticle)
		router.GET("article/content/:id",v1.GetArticleContent)
		router.GET("article/list/:id",v1.GetCate_Article)
		router.POST("login", v1.Login)
	}

	_ = r.Run(utils.HttpPort)


}