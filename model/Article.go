package model

import (
	"OverClock/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct{
	gorm.Model `gorm:"foreignkey:Aid"`
	Category Category
	Title string `gorm:"type:varchar(100);not null " json:"title"`
	Content string `gorm:"type:longtext;not null " json:"string"`
	Aid uint `gorm:"primary_key;auto_increment " json:"id"`
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
	var maps = make(map[string]interface{})
	var article Article
	maps["title"] = data.Title
	maps["content"] = data.Content
	maps["cate"] = data.Category
	err = db.Model(&article).Where("id = ?",id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//删除帖子
func DeleteArticle(id int) int{
	var article Article
	err = db.Where("id = ?",id).Delete(&article).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询帖子
func GetArticle(pageSize int,pageNum int) ([]Article,int) {
	var article []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&article).Error
	if err!= nil && err != gorm.ErrRecordNotFound{
		return article,errmsg.ERROR
	}
	return article,errmsg.SUCCEED
}

//查询帖子的信息
func GetArticeContent(id int) (Article,int){
	var article Article
	err := db.Preload("Category").Where("id = ？",id).First(&article).Error
	if err != nil{
		return article,errmsg.ERROR_ARTICLE_NOT_EXSIT
	}
	return  article,errmsg.SUCCEED
}

//查询某个分类下的帖子
func GetCate_Article(id int,pageSize int,pageNum int) ([]Article,int){
	var article []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Aid = ?",id).Find(&article).Error
	if err != nil{
		return article,errmsg.ERROR_CATEGORY_NOT_EXSIT
	}
	return  article,errmsg.SUCCEED
}
