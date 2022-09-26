package email

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
)

func Reminder(c *gin.Context) {
	eAddress := c.PostForm("email")
	mailHeader := map[string][]string{
		"From":    {"allow202208@126.com"},
		"To":      {eAddress},
		"Subject": {"作业提醒"},
	}

	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	m.SetBody("text/html", "该交作业啦") /*这里改动过一次，最初写的是text/plain，然后运行失败，改成html就成功了*/

	d := gomail.NewDialer("smtp.126.com", 25, "allow202208@126.com", "XGGSMHWHOXUSKDPM")

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}
