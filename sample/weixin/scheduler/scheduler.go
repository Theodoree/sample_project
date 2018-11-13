package scheduler

import (
	"time"
	"github.com/Theodoree/sample_project/sample/weixin/common"
	"fmt"
)

type Scheduler struct{
	CycleTime int64
	WebSpider *common.WebSpider
	Mysql     *common.Mysql
	WeChatList []string
	exit      chan struct{}


}



func (Scheduler)New(WechatList []string,CycleTime int64)*Scheduler{
	chanel := make(chan *common.Article)
	exit   := make(chan struct{})
	Work   := make(chan struct{})
	WebSpider:=common.WebSpider{}.New(chanel,exit,Work)
	Mysql:=common.Mysql{}.New(chanel)
	return &Scheduler{
		CycleTime:CycleTime,
		WebSpider:WebSpider,
		Mysql:Mysql,
		WeChatList:WechatList,
	}
}


func (s *Scheduler)Run(){
	go s.Mysql.Run()
	for {
		for _,nickname:=range s.WeChatList{
			fmt.Println(nickname)
			Soreid,err := s.Mysql.GetSortid(nickname)
			if err !=nil{
				return
			}
			go s.WebSpider.Run(nickname,Soreid)
		}
		for i:=0;i<len(s.WeChatList);i++{
			<-s.exit
		}
		time.Sleep(time.Duration(s.CycleTime))
	}




}