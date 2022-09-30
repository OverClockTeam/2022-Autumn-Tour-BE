package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var(
	AppMode string
	HttpPort string
	JwtKey string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

func init(){
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：",err)
	}
	LoadServe(file)
	LoadData(file)
}

func LoadServe(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("serve").Key("HttpPort").MustString("80")
	JwtKey = file.Section("JwtKey").Key("JwtKey").MustString("12345678")
}

func LoadData(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("OverClock")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("031819lzs")
	DbName = file.Section("database").Key("DbName").MustString("OverClock")

}
