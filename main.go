package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

var db *gorm.DB

// Config 本地读取配置用  与项目无关
type Config struct {
	Account, Password, IP, Port string
}

func init() {
	var err error
	// 读取本地mysql配置  与本项目无关--------------------------------------------
	var c Config
	configFile, err := os.OpenFile("./config.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
	}
	defer configFile.Close()
	reader := bufio.NewReader(configFile)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&c); err != nil {
		log.Println(err)
	}
	//------------------------------------------------------------------------
	logFile, err := os.OpenFile("./log.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	log.SetPrefix("HUST_Helper")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile)) // 同时输出到后台和日志
	if err != nil {
		log.Println(err)
	}
	// 连接数据库
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/go_db?charset=utf8mb4&parseTime=True&loc=Local", c.Account, c.Password, c.IP, c.Port)), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	// 模型绑定
	db.AutoMigrate(&User{})
}
func main() {

}
