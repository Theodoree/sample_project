package main

import (
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/setting"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/scheduler"
)

func main() {
	setting.Setup()
	Sch:=scheduler.Scheduler{}.New(100,"go")
	Sch.Run()
}