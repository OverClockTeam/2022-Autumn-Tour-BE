package main

import (
	"OverClock/model"
	router "OverClock/routes"
)

func main(){

	//引用数据库
	model.InitDb()
	router.InitRouter()
}
