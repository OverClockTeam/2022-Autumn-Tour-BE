package model

import (
	"OverClock/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error
func InitDb(){
	db,err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		fmt.Print("连接数据库失败，请检查参数：", err)
	}

	db.SingularTable(true)//禁用默认表名的复数形式
	db.DB().SetMaxIdleConns(10) //设置连接池最大闲置连接数
	db.DB().SetMaxOpenConns(100) //设置数据库最大连接数量
	db.DB().SetConnMaxLifetime(10 * time.Second) //设置连接的最大可复用时间

}
