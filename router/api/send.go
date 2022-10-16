package api

import (
	"log"
	"fmt"
	"net/http"
	"HC_WJ/util"
	"HC_WJ/emailclt"
	"github.com/gin-gonic/gin"
)

//发送邮件
func Send(c *gin.Context)() {
	//获取登录者信息
	tokenstring := c.GetHeader("token")
	u, err := util.ParseToken(tokenstring)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : "令牌过期",
		})
		return 
	}

	//定义收件人
	var mailTo []string
	for i := 0; i < Len; i++ {
		mailTo = append(mailTo, Users[i].Email)
	}
	//邮件主题为"Hello"
	subject := "交作业"
	// 邮件正文
	body := "请及时交" + subject + "作业！" + "来自" + u.Username + "班长"
 
	err = emailclt.SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}
 
	fmt.Println("send successfully")
}