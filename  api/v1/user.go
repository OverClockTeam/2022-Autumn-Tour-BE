package v1

import (
	"OverClock/model"
	"OverClock/utils/errmsg"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int
//查询用户是否存在
func UserExit(c *gin.Context){

}
//添加用户
func AddUser(c *gin.Context){
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCEED{
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询单个用户
func GetUser(c *gin.Context){

}
//查询用户列表
func GetUsers(c *gin.Context){
	pageSize,_ :=strconv.Atoi(c.Query("pagesize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCEED
	 c.JSON(http.StatusOK,gin.H{
	 	"status":code,
	 	"data":data,
	 	"message":errmsg.GetErrMsg(code),
	 })
}
//编辑用户
func EditUser(c *gin.Context){

}
//删除用户
func DeleteUser(c *gin.Context){

}