package main

import (
	"github.com/wejectchen/ginblog/model"
	"github.com/wejectchen/ginblog/routes"
)

func main() {
	// 数据库
	model.InitDb()
	// 初始化路由组件
	routes.InitRouter()

}
