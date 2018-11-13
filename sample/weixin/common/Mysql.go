package common

import (
	"github.com/ziutek/mymysql/autorc"
	"github.com/Theodoree/sample_project/sample/weixin/config"
	"github.com/Theodoree/sample_project/sample/weixin/utils"
	_ "github.com/ziutek/mymysql/godrv"
)

type Mysql struct {
	Db              *autorc.Conn
	ConsumerChannel <-chan *Article
}

func (Mysql) New(chanel <-chan *Article) *Mysql {
	Db := autorc.New("tcp", "", config.Host, config.User, config.Passwd, config.DB)
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

func (Mysql *Mysql) Insert(a *Article) {
	db := Mysql.Db
	Query := `INSERT INTO weixin (fileid,aid,author,title,
   link,source_url,cover,pubilshed_at,
   digest,content,sortid,insertd_at,published) VALUES(%d,%d,'%s','%s','%s','%s','%s','%s','%s','%s',%d,'%s',%d)`
	_, _, err := db.Query(Query, a.Fileid, a.Aid, db.Escape(a.Author), db.Escape(a.Title), db.Escape(a.Link), db.Escape(a.Source_url), db.Escape(a.Cover), db.Escape(a.Pubilshed_at), db.Escape(a.Digest), db.Escape(a.Content), a.Sortid, db.Escape(a.Insertd_at), a.Pubilsh, )
	utils.CheckErr(err)
}

func (Mysql *Mysql) GetSortid(nickname string) (int, error) {
	db := Mysql.Db
	Query := `SELECT id FROM sort WHERE weixin_name = '%s'`
	Row, _, err := db.Query(Query, Mysql.Db.Escape(nickname))
	if utils.CheckErr(err) {
		return -1, err
	}
	return Row[0].Int(0), nil
}
