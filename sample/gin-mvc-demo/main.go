package main

import (
	"fmt"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/gredis"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
