package v1

import (
	"OverClock/middleware"
	"OverClock/model"
	"OverClock/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code := model.LoginPassword(data.Username,data.Password)
	if  code == errmsg.SUCCEED{
		token,_ = middleware.GetToken(data.Username,data.Password)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":  token,
	})
}
