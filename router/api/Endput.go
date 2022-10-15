package api

import (
	"fmt"
	"net/http"
	"HC_WJ/util"
	"HC_WJ/dbclt"
	"github.com/gin-gonic/gin"
)

//清空本班的作业表
func Endput (c *gin.Context) {
	//获取登录者信息
	tokenstring := c.GetHeader("")
	u, err := util.ParseToken(tokenstring)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : "令牌过期",
		})
		return 
	}
	s := "delete from homework where class = ?"
	r, _ := dbclt.Db.Exec(s, u.Class)
	fmt.Printf("r: %v\n", r)
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/login/index",
		"Context" : "已清空本班作业提交记录",
	})
}