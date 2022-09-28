package model

type DbSettings struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
}
