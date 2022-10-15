package dbclt

import (
	"fmt"
	"HC_WJ/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//根据配置信息链接数据库
func InitDb(settings conf.DbSettings) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",settings.Username, settings.Password, settings.Hostname, settings.Dbname)
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}