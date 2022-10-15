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

		//下载作业
		post.POST("/download", api.Download)

	}
	get := f.Group("/get")
	{
		//检查作业
		get.GET("/check", api.Check)

		//发送邮件 
		get.GET("/sendmail", api.Send)

		//结束作业提交，清空本班的作业表
		get.GET("/endput", api.Endput)

		//上传作业
		get.GET("/upload", api.Upload)
	}
}