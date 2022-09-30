package model

import (
	"OverClock/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct{
	Id   uint   `gorm:"primary_key;auto_increment " json:"id"`
	Name string `gorm:"type:varchar(20);not null " json:"name"`
}

//查询分类是否存在
func CheckCate(name string)(code int){
	var cate Category
	db.Select("id").Where("name = ?",name).First(&cate) // 在数据表中查询
	if cate.Id > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCEED
}
//添加分类
func CreateCate(data *Category)(code int){
	//data.Password = ScryptPassword(data.Password)
	error := db.Create(&data).Error
	if error != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}

//查询分类列表
func GetCate(pageSize int,pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cate).Error
	if err!= nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return cate
}

//编辑分类
func EditCate(id int,data *Category)int{
	var maps = make(map[string]interface{})
	var cate Category
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?",id).Update(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
//查询分类下的所有文章


//删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?",id).Delete(&cate).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCEED
}
