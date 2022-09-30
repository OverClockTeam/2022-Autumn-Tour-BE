package middleware

import (
	"OverClock/utils"
	"OverClock/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type UserClaims struct{
	Username string   `gorm:"type:varchar(20);not null " json:"username"`
	Password string   `gorm:"type:varchar(20);not null " json:"password"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

//生成Token
func GetToken(username string, password string)(string,int){
	c := UserClaims{
		"username",
		"password",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(1) * time.Hour).Unix(), // 过期时间
			Issuer: "OverClock", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	reqtoken := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	token,err := reqtoken.SignedString(JwtKey)
	if err != nil{
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCEED
}

//验证token
func ParseToken(token string)(*UserClaims,int){
	gettoken,_ := jwt.ParseWithClaims(token,&UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey,nil
	})
	if claims,_ := gettoken.Claims.(*UserClaims); gettoken.Valid  {
		return claims, errmsg.SUCCEED
	}else{
		return nil,errmsg.ERROR
	}
}
//中间件
func JWTToken() gin.HandlerFunc{
	return func(c *gin.Context){

		token := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCEED
		if token == ""{
			code = errmsg.ERROR_TOKEN_EXIST
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		pariseToken := strings.SplitN(token," ",2)
		if len(pariseToken) != 2 && pariseToken[0] != "Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		claims,key := ParseToken(pariseToken[1])
		if key == errmsg.ERROR{
			key = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix()>claims.ExpiresAt{
			code = errmsg.ERROT_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username",claims.Username)
		c.Next()
	}
}
