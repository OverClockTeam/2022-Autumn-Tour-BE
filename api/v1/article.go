package v1

import (
	"OverClock/model"
	"OverClock/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加帖子
func AddArticle(c *gin.Context){
	var data model.Article
	model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//编辑帖子
func EditArticle(c *gin.Context){
	var data model.Article
	id ,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	model.EditArticle(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除帖子
func DeleteArticle(c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询帖子列表
func GetArticle(c *gin.Context){
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticle(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询帖子信息
func GetArticleContent(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticeContent(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询某分类下的帖子
func GetCate_Article(c *gin.Context){
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCate_Article(id,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}