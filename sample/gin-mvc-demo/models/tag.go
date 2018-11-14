package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface {}) (count int){
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool{
	var tag Tag
	db.Select(`id`).Where(`name = ?`,name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool{
	var tag Tag
	db.Select(`id`).Where(`id = ?`,id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

//更新数据
func EditTag(id int, data interface {}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}


func AddTag(name string,state int,createBy string)bool{
	db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createBy,
		})
	return true
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope)error{
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("CreatedOn", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("ModifiedOn", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
/*
gorm所支持的回调方法：

创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind
*/
