package main

import (
	"HH_LHY/client"
	"HH_LHY/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init() {
	logfile, err := os.OpenFile("./log/log.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logfile))
	err = client.InitDb("./config.json")
	if err != nil {
		log.Println(err)
	}
	//model.Db, err = gorm.Open(mysql.Open("root:fzxfzxfzx1102@tcp(localhost:49155)/go_db"), &gorm.Config{})
	//if err != nil {
	//	log.Println(err)
	//}
}
func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":9090")
}
