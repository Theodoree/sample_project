package models

import (
	_ `github.com/jinzhu/gorm/dialects/mysql`
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"time"
)


var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}

func Setup() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Info(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10) //最多操作句柄
	db.DB().SetMaxOpenConns(100) //最多打开连接数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}

}