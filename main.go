package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)
type User struct {
  username   string
  email      string
  passward   string
}
func main() {}
