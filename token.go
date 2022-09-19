package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var key = []byte("overclock")

const expiration = time.Hour * 2 // 设置token 2h内有效
func CreateToken(u User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, u).SignedString(key)
}

func ParseToken(token string) (*User, error) {
	claims, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user, ok := claims.Claims.(*User)
	if ok && claims.Valid {
		return user, nil
	}
	return nil, errors.New("invalid Token")
}
