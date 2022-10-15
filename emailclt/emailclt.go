package emailclt

import (
	"HC_WJ/conf"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
)

var mailConn conf.EmailSenderSettings

func InitEmailCtl(settings conf.EmailSenderSettings) {
	mailConn = settings
}

//定义邮箱
func SendMail(mailTo []string, subject string, body string) error {
 
    port := mailConn.Port //转换端口类型为int
 
    m := gomail.NewMessage()
 
    m.SetHeader("From",  m.FormatAddress(mailConn.User, "班长")) //这种方式可以添加别名，即“XX官方”
    m.SetHeader("To", mailTo...)    //发送给多个用户
    m.SetHeader("Subject", subject) //设置邮件主题
    m.SetBody("text/html", body)    //设置邮件正文
 
    d := gomail.NewDialer(mailConn.Host, port, mailConn.User, mailConn.Pass)
 
    err := d.DialAndSend(m)
    return err
}
