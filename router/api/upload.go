package api

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"strconv"
	"net/http"
	"HC_WJ/util"
	"HC_WJ/dbclt"
	"github.com/gin-gonic/gin"
)

//产生随机数
func Randstr() (string) {
	var str string
	for i := 0; i < 18; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		str += string(rune(rand.Intn(74) + 48))
	}
	return str
}

//接收文件
func Upload(c *gin.Context) {
	//接受数据
	tokenstring := c.GetHeader("")
	u, err := util.ParseToken(tokenstring)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message" : "令牌过期",
		})
		return 
	}

	//接受文件
	file, err := c.FormFile("file")
	subject := c.PostForm("subject")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)

	//上传文件到指定目录,格式为 班级_用户名_文件名
	file_name := strconv.Itoa(u.Class) + "_" + u.Username + "_" + file.Filename
	dst := fmt.Sprintf("./subject_file/%s/%s", subject, file_name)
	c.SaveUploadedFile(file, dst)
	c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
		"Url" : "/login/index",
		"Context" : "上传文件成功！请稍后",
	})

	//登记上传信息
	Share := Randstr()
	//清除之前上传的信息,再上传新信息
	s := "delete from homework where username = ? and subject = ?"
	r, _ := dbclt.Db.Exec(s, u.Username, subject)
	fmt.Printf("r: %v\n", r);
	s = "insert into homework (username, subject, share, file_name, class) values(?,?,?,?,?)"
	r, _ = dbclt.Db.Exec(s, u.Username, subject, Share, file_name, u.Class)
	fmt.Printf("r: %v\n", r);
}