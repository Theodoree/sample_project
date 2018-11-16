package utils

import (
	"github.com/Theodoree/sample_project/sample/weixin/config"
	"fmt"
	"time"
	"github.com/parnurzeal/gorequest"
)

func CheckErr(err error)bool{
	er:=err !=nil
	if er{
		return er
	}
	return er

}

func CheckErrs(errs []error)bool{
	for _,value:=range errs{
		if value !=nil{
			return true
		}
	}
	return false
}

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

func Get_Now()string{
	return time.Now().Format("2006-01-02 15:04:05")
}
