package util

import (
	"github.com/dgrijalva/jwt-go"
	"HC_WJ/consts"
	"HC_WJ/model"
	"errors"
	"log"
	"time"
)

func GenerateToken(u model.User) (token string, err error) {
	u.ExpiresAt = time.Now().Unix() + consts.ExpireTime
	return jwt.NewWithClaims(jwt.SigningMethodHS256, u).SignedString([]byte(consts.SecretKey))
}

func ParseToken(token string) (model.User, error) {
	var u model.User
	claims, err := jwt.ParseWithClaims(token, &u, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SecretKey), nil
	})
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	if claims.Valid {
		return u, nil
	}
	return model.User{}, errors.New("invalid Token")
}