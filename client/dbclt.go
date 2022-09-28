package client

import (
	"HH_LHY/model"
	"HH_LHY/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDb(filepath string) (err error) {
	settings, err := util.ReadConfigFromFile(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	DB, err = gorm.Open(mysql.Open(dsn(settings)), &gorm.Config{})
	//db, err = gorm.Open(mysql.Open(dsn(settings)), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	DB.AutoMigrate(&model.User{}, &model.Post{})
	return
}

func dsn(settings model.DbSettings) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.UserName, settings.Password, settings.Host, settings.Port, settings.DbName)
}
