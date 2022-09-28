package middleware

import (
	"HH_LHY/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuditPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		content := c.PostForm("content")
		title := c.PostForm("title")
		hasSensitiveWord1, result1 := services.Audit(content)
		hasSensitiveWord2, result2 := services.Audit(title)
		if hasSensitiveWord1 || hasSensitiveWord2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":        "detected sensitive word ",
				"sensitive word": append(result1, result2...),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
