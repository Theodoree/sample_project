package main

import (
	"github.com/Theodoree/sample_project/sample/weixin/scheduler"
	"github.com/Theodoree/sample_project/sample/weixin/config"
)

func main() {
	var CycleTime int64 =3600
	Schedulers:=scheduler.Scheduler{}.New(config.WeChatList,CycleTime)
	Schedulers.Run()

}


