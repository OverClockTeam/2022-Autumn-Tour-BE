package router

import (
	"HC_WJ/router/api"
	"github.com/gin-gonic/gin"
)

func UseMyRouter(f *gin.Engine) {
	post := f.Group("/post")
	{
		//处理点击注册后的情况
		post.POST("/register", api.Register)

		//处理点击登录后的情况
		post.POST("/login", api.Login)

		//发送邮件 
		post.POST("/sendmail", api.Send)
		
		//上传作业
		post.POST("/upload", api.Upload)

		//结束作业提交，清空本班的作业表
		post.POST("/endput", api.Endput)

	}
	get := f.Group("/get")
	{
		//下载作业
		get.GET("/download", api.Download)

		//检查作业
		get.GET("/check", api.Check)

	}
}