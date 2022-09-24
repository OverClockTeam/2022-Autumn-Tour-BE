package model

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Username string   `gorm:"type:varchar(20);not null " json:"username"`
	Password string   `gorm:"type:varchar(20);not null " json:"password"`
	Email string  `gorm:"type:varchar(30);not null " json:"email"`
	Role int  `gorm:"type:int;not null " json:"role"`
}


