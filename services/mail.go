package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func SendEmail(email string) (code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "2116545753@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "verification")
	code = createCode()
	m.SetBody("text/html", fmt.Sprintf("your verifity code is %s", code))
	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "2116545753@qq.com", "knecegtlqahlejch")
	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	return
}

func createCode() string {
	var code string
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
