package main

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var dbinfo = "@tcp(45.32.116.100)/MM_userinfo?charset=utf8&parseTime=True&loc=Local"
var db *gorm.DB

func initDB() (err error) {
	// DSN:Data Source Name
	db, err := gorm.Open("mysql", dbinfo)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	defer db.Close()
	return err
}
