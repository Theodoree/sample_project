package setting

import (
	"time"
	"gopkg.in/ini.v1"
	"log"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
	PrefixUrl   string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	User     string
	Password string
	Host     string
	Db       string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}




func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting) //把数据解析给app对象
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting errs: %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting) //把数据解析给Server对象
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting errs: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting) //解析给Database对象
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting errs: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting) //解析给Database对象
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting errs: %v", err)
	}
}

