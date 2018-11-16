package GoMysql

import (
	_ "github.com/ziutek/mymysql/godrv"
	"github.com/ziutek/mymysql/autorc"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/models"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/pkg/setting"
	"github.com/Theodoree/sample_project/sample/Boos-zhiping/utils"
	"log"
	"fmt"
)

type Mysql struct {
	Db              *autorc.Conn
	ConsumerChannel <-chan *models.Recruitment
}

func (Mysql) New(chanel <-chan *models.Recruitment) *Mysql {
	Db := autorc.New("tcp", "", setting.DatabaseSetting.Host, setting.DatabaseSetting.User, setting.DatabaseSetting.Passwd, setting.DatabaseSetting.DB)
	Db.Register("set names utf8")
	return &Mysql{
		Db:              Db,
		ConsumerChannel: chanel,
	}
}

func (Mysql *Mysql) Run() {
	for {
		v, ok := <-Mysql.ConsumerChannel
		if ok {
			Mysql.Insert(v)
		}
	}
}

func (Mysql *Mysql) Insert(a *models.Recruitment) {
	db := Mysql.Db
	Query := `INSERT INTO demo.Boos (Position,Url,Desc,CompanyInfo,Location,Company,Slary,InsertTime) VALUES(
	'%s','%s','%s','%s','%s','%s','%s','%s')`
	_, _, err := db.Query(Query,db.Escape(a.Position),db.Escape(a.Url),db.Escape(a.Desc),
		db.Escape(a.CompanyInfo),db.Escape(a.Location),db.Escape(a.Company),db.Escape(a.Slary),
		db.Escape(utils.Get_Now()))
	if utils.CheckErr(err)  {
		fmt.Printf(Query,db.Escape(a.Position),db.Escape(a.Url),db.Escape(a.Desc),
			db.Escape(a.CompanyInfo),db.Escape(a.Location),db.Escape(a.Company),db.Escape(a.Slary),
			db.Escape(utils.Get_Now()))
		log.Println(`Insert `,err)
		return
	}
	fmt.Println(`插入成功`,a.Position)
}

func (Mysql *Mysql)Clear(){
	Mysql.Db.Clone()

}