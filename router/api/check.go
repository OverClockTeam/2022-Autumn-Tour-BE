package api

import (
	"net/http"
	"HC_WJ/model"
	"HC_WJ/dbclt"
	"HC_WJ/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Users []model.Homework
var Len int
func Check(c *gin.Context) {
	//接受数据
	subject := c.PostForm("subject")
	tokenstring := c.GetHeader("token")
	u, err := util.ParseToken(tokenstring)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : "令牌过期",
		})
		return 
	}
	//找出自己班级没交某个作业的人并输出
	s := "select users.username, users.email, homework.username from users left join homework on users.username = homework.username and users.class = ? and homework.subject = ?"
	r, _ := dbclt.Db.Query(s, u.Class, subject)
	Len = 0
	var u1 model.Homework
	for r.Next() {
		r.Scan(&u1.Username, &u1.Email, &u1.Subject)
		//若没有提交则左右表名字不匹配
		if u1.Subject != u1.Username  {
			u1.Subject = subject
			Users = append(Users, u1)
			Len++
		}
	}
	
	//输出到html页面
	c.JSON(http.StatusOK, gin.H{
		"res": Users,
	})

}