package gomysql

import (
	_ "github.com/ziutek/mymysql/godrv"
	"github.com/ziutek/mymysql/autorc"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/setting"
)


var Db =  &autorc.Conn{}
func Setup()  {
	Db = autorc.New("tcp", "", setting.DatabaseSetting.Host, setting.DatabaseSetting.User, setting.DatabaseSetting.Passwd, setting.DatabaseSetting.DB)
	Db.Register("set names utf8")
}


