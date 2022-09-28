package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	IsActivated        bool   `json:"isActivated" gorm:"false"`
	Name               string `json:"name" gorm:"name"`
	Password           string `json:"password" gorm:"password"`
	Email              string `json:"email" gorm:"primaryKey;unique"`
	VerifyCode         string `json:"verifyCode" gorm:"null"`
	jwt.StandardClaims `json:"-" gorm:"-"`
}

func (u User) IsEmpty() bool {
	return !(u.Name != "" && u.Email != "")
}
