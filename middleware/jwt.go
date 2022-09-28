package middleware

import (
	"OverClock/utils"
	"OverClock/utils/errmsg"
	"github.com/dgrijalva/jwt-go/v4"
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

//生成Token
func GetToken(username string, password string)(string,int){
	expiredTime := time.Now().Add(2 * time.Hour)
	GetClaims := UserClaims{
		Username:       username,
		Password:       password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:"OverClock",
		},
	}
	reqtoken := jwt.NewWithClaims(jwt.SigningMethodHS256,GetClaims)
	token,err := reqtoken.SignedString(JwtKey)
	if err != nil{
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCEED
}

//验证token
func ParseToken(token string)(*UserClaims,int){
	gettoken,err := jwt.ParseWithClaims(token,&UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey,nil
	})
	if claims,code := gettoken.Claims.(*UserClaims); gettoken.Valid  {
		return claims, errmsg.SUCCEED
	}else{
		return nil,errmsg.ERROR
	}
}
//类似过滤器
func JWTAuth() gin.HandlerFunc{
	return func(c *gin.Context){
		if strings.Contains(c.Request.RequestURI,"login"){
			return
		}

		token := c.Request.Header.Get("token")
		code := errmsg.SUCCEED
		if token == ""{
			code = errmsg.ERROR_TOKEN_EXIST}
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
			reutrn
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
