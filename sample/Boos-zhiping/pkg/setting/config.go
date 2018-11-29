package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type Server struct{
	Proxy string
	TimeOut_url string
	Max_Try int
	Base_Url string


}


type Database struct{
	DB string
	User string
	Passwd string
	Host string
}


var ServerSetting = &Server{}
var DatabaseSetting =&Database{}


func Setup(){
	Cfg, err :=ini.Load("config/config.ini")
	if err !=nil{
		log.Fatal(`Fail to parse ''config/config.ini`)
		return
	}

	err = Cfg.Section("Server").MapTo(ServerSetting) //把数据解析给app对象
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting errs: %v", err)
	}


	err = Cfg.Section("Database").MapTo(DatabaseSetting) //把数据解析给app对象
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting errs: %v", err)
	}
}



