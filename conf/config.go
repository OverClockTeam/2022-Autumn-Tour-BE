package conf

type Config struct {
	DbSettings          DbSettings          `json:"DbSettings"`
	EmailSenderSettings EmailSenderSettings `json:"EmailSenderSettings"`
}

type DbSettings struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}

type EmailSenderSettings struct {
	Host string `json:"ServerAddress"`
	Port int    `json:"ServerPort"`
	Pass string `json:"ServerPassword"`
	User string `json:"ServerName"`
}