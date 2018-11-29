package main

import (
	"github.com/Theodoree/sample_project/sample/mvc/pkg/setting"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/gredis"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/gomysql"
	"github.com/Theodoree/sample_project/sample/mvc/router"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	setting.Setup()
	gredis.Setup()
	gomysql.Setup()
	defer gredis.RedisConn.Close()
	defer gomysql.Db.Clone()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 2 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, router.Setup())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server errs: %v", err)
	}
}
