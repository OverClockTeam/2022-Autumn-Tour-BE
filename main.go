package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
)
type User struct {
  username   string
  email      string
  passward   string
}
