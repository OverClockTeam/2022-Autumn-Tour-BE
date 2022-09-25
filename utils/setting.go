package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone       int
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")//读入文件config.ini，且load里面的*files是结构体类型，下面需要再次实现（传值）一下
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")//“debug”为默认值，可转换
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("agiee")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("test")
}
