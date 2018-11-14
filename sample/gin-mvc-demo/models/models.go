package models

import (
	_ `github.com/jinzhu/gorm/dialects/mysql`
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
)


var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn string `json:"created_on"`
	//ModifiedOn string `json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

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
}

func CloseDB() {
	defer db.Close()
}

