package api

import (
	"HH_LHY/client"
	"HH_LHY/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserIndex(c *gin.Context) {
	name := c.Param("name")
	var u model.User
	err := client.DB.Where("name = ?", name).First(&u).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"info": u,
		})
	}
}
