package common

import (
	"fmt"
	"github.com/Theodoree/sample_project/sample/weixin/config"
	"github.com/parnurzeal/gorequest"
	"regexp"
	"encoding/json"
	"log"
	"github.com/ziutek/mymysql/autorc"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/Theodoree/sample_project/sample/weixin/utils"
	"time"
	"math/rand"
)

type WebSpider struct {
	WorkChanel chan<- *Article
	Exit       chan struct{}
	Db         *autorc.Conn
}

func (WebSpider) New(chanel chan<- *Article, exit, Work chan struct{}) *WebSpider {
	Db := autorc.New("tcp", "", config.Host, config.User, config.Passwd, config.DB)
	Db.Register("set names utf8")
	return &WebSpider{
		WorkChanel: chanel,
		Exit:       exit,
		Db:         Db,
	}
}

func (Spider *WebSpider) Run(KeyWord string, Soreid int) {
	WeChat_url := Spider.GetWeChat(KeyWord)
	if WeChat_url == ""{
		log.Println(KeyWord,` GetWeChat 超时 `)
		return
	}
	ArticleList := Spider.GetArticleUrl(WeChat_url, Spider.Db, Soreid)
	for _, Article := range ArticleList {
		Spider.WorkChanel <- Article
	}
	Spider.Exit <- struct{}{}

}

func (Spider *WebSpider) GetWeChat(KeyWord string) (string) {
	try := 0
	for try < config.Max_Try {
		time.Sleep(time.Duration(rand.Float64()*5))
		Url := fmt.Sprintf(`https://weixin.sogou.com/weixin?type=1&s_from=input&query=%s&ie=utf8&_sug_=n&_sug_type_=`, KeyWord)
		proxy := utils.GetProxy()
		request := gorequest.New()
		request.Header.Set(`Host`, `weixin.sogou.com`)
		request.Header.Set(`User-Agent`, `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
		res, body, errs := request.Get(Url).Proxy(proxy).End()
		if utils.CheckErrs(errs) {
			try+=1
			continue
		}
		if res.StatusCode != 200 {
			fmt.Println(res.StatusCode)
			try+=1
			continue
		}

		find, _ := regexp.Compile(`您的访问过于频繁，为确认本次访问为正常用户行为，需要您协助验证`)
		if find.FindAll([]byte(body), -1) != nil {
			fmt.Println(`代理被限制`)
			proxy := strings.Split(proxy, `http://`)[1]
			utils.TimeOut(proxy)
			try+=1
			continue
		}
		doc, _ := goquery.NewDocumentFromReader(res.Body)
		url := ``
		doc.Find(`.txt-box a `).Each(func(i int, s *goquery.Selection) {
			if s.Text() == KeyWord {
				val, _ := s.Attr(`href`)
				url = val
			}

		})
		return url
	}
	return ""
}

func (Spider *WebSpider) GetArticleUrl(WechatUrl string, Db *autorc.Conn, Soreid int) ([]*Article) {
	try := 0
	for try < config.Max_Try {
		time.Sleep(time.Duration(rand.Float64()*5))
		ArticleList := []*Article{}
		proxy := utils.GetProxy()
		request := gorequest.New()
		request.Header.Set(`User-Agent`, `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
		res, body, errs := request.Get(WechatUrl).Proxy(proxy).End()
		if utils.CheckErrs(errs) {
			try+=1
			continue
		}
		if res.StatusCode != 200 {
			try+=1
			continue
		}
		find, _ := regexp.Compile(`<div class="weui_cells_tips">为了保护你的网络安全，请输入验证码</div>`)

		if find.Find([]byte(body)) != nil {
			fmt.Println(`代理被限制`)
			proxy := strings.Split(proxy, `http://`)[1]
			utils.TimeOut(proxy)
			try+=1
			continue
		}
		RegexpJson, _ := regexp.Compile(`var msgList = ({.*?});`)

		msgList := RegexpJson.FindSubmatch([]byte(body))
		list := List{}
		if msgList == nil {
			try+=1
			continue
		}
		json.Unmarshal(msgList[1], &list)
		for _, obj := range list.List {
			if Is_Have(Soreid, obj.AppMsg.Title, Db) {
				continue
			} else {
				Articletest := &Article{
					Fileid:       obj.AppMsg.Fileid,
					Aid:          obj.CommMsg.Id,
					Author:       obj.AppMsg.Author,
					Title:        obj.AppMsg.Title,
					Link:         obj.AppMsg.ContentUrl,
					Source_url:   obj.AppMsg.SourceUrl,
					Cover:        obj.AppMsg.Cover,
					Pubilshed_at: utils.Get(int64(obj.CommMsg.DateTime)),
					Digest:       obj.AppMsg.Digest,
					Sortid:       Soreid,
					Insertd_at:   utils.Get_Now(),
					Pubilsh:      1,
				}
				ok :=Articletest.GetContent()
				if ok {
					ArticleList = append(ArticleList, Articletest)

				}
			}
		}
		return ArticleList
	}
	return nil
	}


func Is_Have(Soreid int, title string, Db *autorc.Conn) bool {
	res, _, err := Db.Query(`select * from weixin WHERE title= '%s' AND sortid=%d`, Db.Escape(title), Soreid)
	if utils.CheckErr(err) {
		log.Println(err)
		return true
	}
	if len(res) > 0 {
		return true
	}
	return false
}
