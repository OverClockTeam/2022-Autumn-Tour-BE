package main

import (
	"HC_WJ/conf"
	"HC_WJ/dbclt"
	"HC_WJ/emailclt"
	"HC_WJ/router"
	"github.com/gin-gonic/gin"
)
func SettingUpEnvironment() {
	//读入配置文件
	c := conf.ReadSettingsFromFile("Config.json")

	//配置数据库
	dbclt.InitDb(c.DbSettings)

	//配置邮箱
	emailclt.InitEmailCtl(c.EmailSenderSettings)
}
func main() {
	//设置初始环境
	SettingUpEnvironment()
	
	//创建路由对象
	f := gin.Default()

	//使用路由功能
	router.UseMyRouter(f)

	//启动
	f.Run(":8080");

}