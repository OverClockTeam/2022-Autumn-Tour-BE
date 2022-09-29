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
func JWTAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.ERROR_TOKEN_EXIST,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.ERROR_TOKEN_TYPE_WRONG,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		mc, key := ParseToken(parts[1])
		if key != errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.ERROT_TOKEN_RUNTIME,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}

}
