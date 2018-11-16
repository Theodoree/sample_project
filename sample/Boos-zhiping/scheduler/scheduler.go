package scheduler

import (
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/Spider"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/GoMysql"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/models"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/setting"
	"fmt"
)

type Scheduler struct {
	WebSpider *Spider.WebSpider
	Mysql     *GoMysql.Mysql
	exit      chan struct{}
	page      int
	work      chan *models.Recruitment
	query     string
}

func (Scheduler) New(page int, query string) *Scheduler {
	Work := make(chan *models.Recruitment)
	Exit := make(chan struct{})
	spider := Spider.WebSpider{}.New(Work, Exit)
	db := GoMysql.Mysql{}.New(Work)
	return &Scheduler{
		WebSpider: spider,
		exit:      Exit,
		page:      page,
		Mysql:     db,
		work:      Work,
		query:query,
	}
}

const (
	url = `/c101010100/?query=%s&page=%d&ka=page-%d`
)

func (S *Scheduler) Run() {
	go S.Mysql.Run()
	for i := 0; i < S.page; i++ {
		Rquest_Url := setting.ServerSetting.Base_Url + fmt.Sprintf(url, S.query,i, i)
		go S.WebSpider.Run(Rquest_Url)
	}

	for i := 0; i < S.page; i++ {
		<-S.exit
	}
	S.Clear()
}

func (S *Scheduler) Clear() {
	S.Mysql.Clear()
	S.WebSpider.Clear()
	close(S.work)
	close(S.exit)
}
