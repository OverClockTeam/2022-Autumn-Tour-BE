package router

import (
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"msg":"ok",
			})
		})
	}

	_ = r.Run(utils.HttpPort)


}