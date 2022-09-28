package util

import (
	"HH_LHY/consts"
	"HH_LHY/model"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"os"
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

func ReadConfigFromFile(filepath string) (settings model.DbSettings, err error) {
	config, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer config.Close()
	var data []byte
	data, err = io.ReadAll(config)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
