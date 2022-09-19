package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Name               string `json:"name" `
	Password           string `json:"password" `
	Email              string `json:"email" `
	jwt.StandardClaims `gorm:"-"`
}

func NewUser(args ...string) *User {
	if len(args) == 0 {
		return new(User)
	}
	if len(args) > 3 {
		panic("wrong args")
	}
	var u *User
	for i, arg := range args {
		if i == 0 {
			u.Name = arg
		} else if i == 1 {
			u.Password = arg
		} else {
			u.Email = arg
		}
	}
	u.ExpiresAt = time.Now().Add(expiration).Unix()
	u.Issuer = "lhy"
	return u
}
