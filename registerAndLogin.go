package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Slice []User
var State = make(map[string]interface{})

func register(c *gin.Context) {
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	Bool := isExist(name)
	if Bool {
		State["state"] = 1
		State["text"] = "此用户已存在！"
	} else {
		addUser(name, passwd)
		State["state"] = 1
		State["text"] = "注册成功！"
	}
	c.String(http.StatusOK, "%v", State)
}

func login(c *gin.Context) {
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	Bool := isExist(name)
	if Bool {
		Bool_Pwd := isRight(name, passwd)
		if Bool_Pwd {
			State["state"] = 1
			State["text"] = "登录成功！"
		} else {
			State["state"] = 0
			State["text"] = "密码错误！"
		}
	} else {
		State["state"] = 2
		State["text"] = "登录失败！此用户尚未注册！"
	}

	c.String(http.StatusOK, "%v", State)
}

func addUser(name string, passwd string) {
	var user User
	user.Name = name
	user.Passwd = passwd
	user.Id = len(Slice) + 1
	Slice = append(Slice, user)
}

func isExist(user string) bool {
	if len(Slice) == 0 {
		return false
	} else {
		for _, v := range Slice {
			if v.Name == user {
				return true
			}
		}
	}
	return false
}

func isRight(user string, passwd string) bool {
	for _, v := range Slice {
		if v.Name == user {
			return v.Passwd == passwd
		}
	}
	return false
}
