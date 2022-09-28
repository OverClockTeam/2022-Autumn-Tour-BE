package router

import (
	v1 "OverClock/ api/v1"
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user:/id",v1.DeleteUser)

		//分类模块的路由接口
		router.POST("category/add",v1.AddCate)
		router.GET("category",v1.GetCate)
		router.PUT("category/:id",v1.EditCate)
		router.DELETE("category/:id",v1.DeleteCate)
		//帖子模块的路由接口
		router.POST("article/add",v1.AddArticle)
		router.GET("article",v1.GetArticle)
		router.GET("article/content/:id",v1.GetArticleContent)
		router.PUT("article/:id",v1.EditArticle)
		router.DELETE("article/:id",v1.DeleteArticle)
		router.GET("article/list/:id",v1.GetCate_Article)

	}

	_ = r.Run(utils.HttpPort)


}