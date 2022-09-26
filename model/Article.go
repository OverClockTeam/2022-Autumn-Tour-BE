package model

import (
	"OverClock/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct{
	gorm.Model
	Category Category
	Title string `gorm:"type:varchar(100);not null " json:"title"`
	Content string `gorm:"type:longtext;not null " json:"string"`
}

//查询帖子是否存在
func CheckArticle(title string) int{
	var article Article
	db.Select("id").Where("title = ?",title).First(&article)
	if article.ID <= 0{
		return errmsg.SUCCEED
	}
	return errmsg.ERROR_TITLE_USED
}

//添加帖子
func CreateArticle(data *Article)int {
	error := db.Create(&data).Error
	if error !=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//编辑帖子
func EditArticle(id int, data *Article) int{
	
}
