package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./front/*.html") //加载html文件

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.GET("/login", JWTAuthMiddleHandler(), func(c *gin.Context) {
		value, _ := c.Get("user")
		user := value.(*User)
		_, exists := c.Get("isAccessible")
		// 验证成功   重定向至主页
		if exists {
			c.Redirect(http.StatusTemporaryRedirect, "/index/"+user.Name)
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/register", registerMiddleHandler(), func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")
		if isVerified, exists := c.Get("isVerified"); isVerified.(bool) && exists {
			email, _ := c.Get("email")
			user := NewUser(name, password, email.(string))
			db.Create(user) //创建成功将数据写入数据库
			token, err := CreateToken(*user)
			if err != nil {
				log.Println(err)
				return
			}
			// 将生成的token返回给客户端
			c.JSON(http.StatusOK, gin.H{
				"message": "注册成功",
				"token":   token,
			})
		}
		return
	})
	r.POST("/login", func(c *gin.Context) {

	})
	return r
}

// JWTAuthMiddleHandler JWTAuthMid jwt鉴权中间件
func JWTAuthMiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取header中的token
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"code":    401,
			})
			c.Next()
			return
		}
		parts := strings.Split(auth, " ")

		if len(parts) != 2 { // 2 or 3 ?
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "wrong auth format",
				"code":    401,
			})
			c.Next()
			return
		}
		user, err := ParseToken(parts[1]) //parts[1] == auth中的token
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err,
				"code":    401,
			})
			c.Next()
			return
		}
		c.Set("user", user)
		c.Set("isAccessible", true)
		c.Next()
	}
}

func registerMiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		// 默认只有学生账户能登录
		if strings.HasPrefix(email, "U") && strings.HasSuffix(email, "@hust.edu.cn") {
			// 判断位数  是否数字暂未判断
			if len(email[1:len(email)-12]) != 9 {
				c.JSON(http.StatusOK, gin.H{
					"message": "invalid email",
				})
				return
			}
			// 发送邮件验证
			//isVerified 为是否验证邮箱
			//c.Set("isVerified",isVerified)
			c.Set("email", email)
			c.Next()
		}
	}
}
