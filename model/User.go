package model

import (
	"OverClock/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model
	Username string   `gorm:"type:varchar(20);not null " json:"username"`
	Password string   `gorm:"type:varchar(20);not null " json:"password"`
	Email string  `gorm:"type:varchar(30);not null " json:"email"`
	Role int  `gorm:"type:int;not null " json:"role"`
}
//查询用户是否存在
func CheckUser(name string)(code int){
	var users User
	db.Select("id").Where("username = ?",name).First(&users) // 在数据表中查询
	if users.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}
//添加用户
func CreateUser(data *User)(code int){
	error := db.Create(&data).Error
	if error != nil{
		return errmsg.ERRO
	}
	return errmsg.SUCCEED
}

//查询用户列表
func GetUsers(pageSize int,pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

//编辑用户


//删除用户
