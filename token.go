package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type User struct {
	Name     string
	Password string
	jwt.StandardClaims
}

var key = []byte("overclock")

func createToken(name, password string) (string, error) {
	u := User{
		Name:     name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "lhy",
			NotBefore: time.Now().Unix(), // 立即生效
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return token.SignedString(key)
}

func parseToken(token string) (*User, error) {
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
