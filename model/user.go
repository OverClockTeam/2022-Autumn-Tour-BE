package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username           string 
	Password           string 
	Email              string 
	Class              int 
	Monitor			   int
	jwt.StandardClaims `json:"-" gorm:"-"`
}
