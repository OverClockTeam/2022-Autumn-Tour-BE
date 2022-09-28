package middleware

import (
	"HH_LHY/consts"
	"HH_LHY/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// JWTAuthMiddleHandler JWTAuthMid jwt鉴权中间件
func JWTAuthMiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取header中的token
		token := c.Request.Header.Get(consts.CookieName)
		if token == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{
			//	"message": "Unauthorized",
			//	"code":    401,
			//})
			c.Next()
			return
		}
		u, err := util.ParseToken(token)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"code":    401,
			})
			c.Next()
			return
		}
		if u.IsEmpty() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Next()
			return
		}
		c.Set("user", u)
		c.Next()
	}
}

func RegisterMiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		// 默认只有学生账户能登录
		if strings.HasPrefix(email, "U") && strings.HasSuffix(email, "@hust.edu.cn") && email != "" {
			// 判断位数  是否数字暂未判断
			if len(email[1:len(email)-12]) == 9 {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong email format",
		})
		c.Abort()
		return
	}
}
