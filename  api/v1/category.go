package v1

import (
	"OverClock/model"
	"OverClock/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

///添加分类
func AddCate(c *gin.Context){
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCEED{
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询分类列表
func GetCate(c *gin.Context){
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCate(pageSize,pageNum)
	code = errmsg.SUCCEED
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//编辑分类
func EditCate(c *gin.Context){
	var data model.Category
	id ,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Name)
	if code == errmsg.SUCCEED{
		model.EditCate(id,&data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除分类
func DeleteCate(c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
