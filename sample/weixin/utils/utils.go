package utils

import (
	"time"
	"fmt"
	"github.com/Theodoree/sample_project/sample/weixin/config"
	"github.com/parnurzeal/gorequest"
)

func GetProxy()string{
	_,Proxy,errs:=gorequest.New().Get(config.Proxy).End()
	for _,value:=range errs{
		CheckErr(value)
	}

	return "http://"+Proxy
}

func TimeOut(ip string){
	reqeuest :=gorequest.New()
	Send := fmt.Sprintf(`{"ip":"%s"}`,ip)
	reqeuest.Post(config.TimeOut_url).Type(`form`).Send(Send).End()
}


func CheckErr(err error)bool{
	re0 := err !=nil
	if re0{
		return re0
	}
	return re0
}


func Get(time_ int64) string{
	return time.Unix(time_, 0).Format("2006-01-02 15:04:05")
}
func Get_Now()string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func CheckErrs(errs []error)bool{
	for _,value:=range errs{
		if value !=nil{
			return true
	}
	}
	return false
}
