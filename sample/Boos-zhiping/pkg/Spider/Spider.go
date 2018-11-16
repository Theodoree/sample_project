package Spider

import (
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/models"
	"github.com/ziutek/mymysql/autorc"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/setting"
	"log"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/utils"
	"github.com/parnurzeal/gorequest"
	"time"
	"math/rand"
	"strings"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"sync"
)

type WebSpider struct {
	WorkChanel chan<- *models.Recruitment
	Exit       chan struct{}
	Db         *autorc.Conn
}

func (Spider *WebSpider) Run(url string) {
	UrlList := Spider.GetRecruitmentUrl(url)
	if UrlList == nil {
		log.Println(url, ` Url`)
		return
	}
	fmt.Println(`Do`)
	RecruitmentList := Spider.Do(UrlList)
	fmt.Println(`Spider.WorkChanel <- Recruitment`)
	for _, Recruitment := range RecruitmentList {
		Spider.WorkChanel <- Recruitment
	}
	Spider.Exit <- struct{}{}

}

func (Spider *WebSpider) GetRecruitmentUrl(url string) (UrlList []string) {
	try := 0
	for try < setting.ServerSetting.Max_Try {
		time.Sleep(time.Duration(rand.Float64() * 5))
		proxy := utils.GetProxy()
		request := gorequest.New()
		request.Header.Set(`User-Agent`, `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
		request.Header.Set(`:authority`, `www.zhipin.com`)
		request.Header.Set(`:method`, `GET`)
		request.Header.Set(`:scheme`, `https`)
		request.Header.Set(`:path`, strings.Split(url, setting.ServerSetting.Base_Url)[1])
		res, body, errs := request.Proxy(proxy).Get(url).End()
		if utils.CheckErrs(errs) {
			try += 1
			continue
		}
		if res.StatusCode != 200 {
			fmt.Println(res.StatusCode)
			try += 1
			continue
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if utils.CheckErr(err) {
			log.Println(err)
			return
		}
		doc.Find(`div.job-list ul li div div.info-primary h3.name a`).Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr(`href`)
			Base_Url := setting.ServerSetting.Base_Url
			url := Base_Url + href
			if Spider.IsHave(url) {
				UrlList = append(UrlList, url)
			}

		})

		return UrlList
	}
	return nil
}

func (Spider *WebSpider) Do(UrlList []string, ) (RecruitmentList []*models.Recruitment) {

	for _, url := range UrlList {
		try := 0
		for try < setting.ServerSetting.Max_Try {
			time.Sleep(time.Duration(rand.Float64() * 10))
			proxy := utils.GetProxy()
			request := gorequest.New()
			request.Header.Set(`User-Agent`, `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36`)
			request.Header.Set(`:authority`, `www.zhipin.com`)
			request.Header.Set(`:method`, `GET`)
			request.Header.Set(`:scheme`, `https`)
			request.Header.Set(`:path`, strings.Split(url, setting.ServerSetting.Base_Url)[1])
			res, body, errs := request.Proxy(proxy).Get(url).End()
			if utils.CheckErrs(errs) {
				try += 1
				continue
			}
			if res.StatusCode != 200 {
				fmt.Println(res.StatusCode)
				try += 1
				continue
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
			if utils.CheckErr(err) {
				log.Println(err)
				return
			}
			doc.Find(`#main`).Each(func(i int, s *goquery.Selection) {
				Position := s.Find(` div.job-primary  div.info-primary div.name h1 `).Text()
				SalaryRege, _ := regexp.Compile(`\d*K-\d*K`)
				salary := fmt.Sprintf("%s", SalaryRege.Find([]byte(s.Find(` div.job-primary  div.info-primary span.badge `).Text())))
				Desc := strings.TrimSpace(s.Find(`div.detail-content div:nth-child(1) div.text  `).Text())
				CompanyInfo := strings.TrimSpace(s.Find(`div.detail-content div:nth-child(2) div.text  `).Text())
				Company := strings.TrimSpace(s.Find(`div.detail-content div:nth-child(4) div.name  `).Text())
				Location := strings.TrimSpace(s.Find(`div.detail-content div:nth-child(5) div.location-address  `).Text())
				RecruitmentList = append(RecruitmentList, &models.Recruitment{
					Position:    Position,
					Url:         url,
					Desc:        Desc,
					CompanyInfo: CompanyInfo,
					Location:    Location,
					Company:     Company,
					Slary:       salary,
				})
			fmt.Println(Position)
			})
			break
		}
	}
	return RecruitmentList
}

func (WebSpider) New(work chan<- *models.Recruitment, exit chan struct{}) *WebSpider {
	Db := autorc.New("tcp", "", setting.DatabaseSetting.Host, setting.DatabaseSetting.User, setting.DatabaseSetting.Passwd, setting.DatabaseSetting.DB)
	Db.Register("set names utf8")
	return &WebSpider{
		WorkChanel: work,
		Exit:       exit,
		Db:         Db,
	}

}

func (Web *WebSpider) Clear() {
	defer Web.Db.Clone()

}
var  Mx sync.Mutex

func (Web *WebSpider) IsHave(url string) bool {
	Mx.Lock()
	db := Web.Db
	query := ` SELECT id FROM demo.Boos  WHERE Url = "%s"`
	rows, _, err := db.Query(query, url)
	Mx.Unlock()
	if utils.CheckErr(err) {
		log.Println(` IsHave `,err)
		return false
	}
	if len(rows) >0 {
		return false
	}
	return true

}
