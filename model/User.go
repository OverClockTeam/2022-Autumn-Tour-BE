package model

import (
	"OverClock/utils/errmsg"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct{
	gorm.Model
	Username string   `gorm:"type:varchar(20);not null " json:"username"`
	Password string   `gorm:"type:varchar(20);not null " json:"password"`
	Email string  `gorm:"type:varchar(40);not null " json:"email"`
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
	data.Password = ScryptPassword(data.Password)
	error := db.Create(&data).Error
	if error != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询用户列表
func GetUsers(pageSize int,pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err!= nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

//编辑用户
func EditUser(id int,data *User)int{
	var maps = make(map[string]interface{})
	var user User
	maps["username"] = data.Username
	err = db.Model(&user).Where("id = ?",id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?",id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//密码加密
func ScryptPassword(password string)string{
	const KeyLen = 10
	salt := make([]byte,8)
	salt = []byte{10,20,30,40,50,60,80,70}
	HashPw,err := scrypt.Key([]byte(password),salt,1024,8,1,KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw
}