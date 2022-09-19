package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	logFile, err := os.OpenFile("./log.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log.SetPrefix("HUST_Helper")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	db, err = gorm.Open(mysql.Open("root:fzxfzxfzx1102@tcp(localhost:49155)/go_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&User{})
}
func main() {

}
