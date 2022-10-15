package api

import (
	"net/http"
	"HC_WJ/dbclt"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	var subject, file_name string
	share := c.PostForm("share")
	s := "select file_name, subject from homework where share = ?"
	err := dbclt.Db.QueryRow(s, share).Scan(&file_name, &subject)
	if err != nil {
		c.HTML(http.StatusOK, "html/jump.tmpl", gin.H{
			"Url" : "/login/index",
			"Context" : "没有找到分享码！",
		})
		return 
	}
	c.File("./subject_file/" + subject + "/" + file_name)
}