package main

import (
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"fmt"
	"log"
	"syscall"
	"github.com/fvbock/endless"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}