package common

import (
	"github.com/Theodoree/sample_project/sample/weixin/utils"
	"github.com/parnurzeal/gorequest"
	"github.com/Theodoree/sample_project/sample/weixin/config"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type List struct {
	List []Msg `json:"list"`
}

type Msg struct {
	AppMsg  `json:"app_msg_ext_info"`
	CommMsg `json:"comm_msg_info"`
}

type CommMsg struct {
	Id       int `json:"id"`
	DateTime int `json:"datetime"`
}

type AppMsg struct {
	Author     string `json:"author"`
	ContentUrl string `json:"content_url"`
	SourceUrl  string `json:"source_url"`
	Cover      string `json:"cover"`
	Digest     string `json:"digest"`
	Fileid     int    `json:"fileid"`
	Title      string `json:"title"`
}

type Article struct {
	Fileid       int    `json:"fileid"`
	Aid          int    `json:"aid"`
	Author       string `json:"author"`
	Title        string `json:"title"`
	Link         string `json:"link"`
	Source_url   string `json:"source_url"`
	Cover        string `json:"cover"`
	Pubilshed_at string `json:"pubilshed_at"`
	Digest       string `json:"digest"`
	Content      string `json:"content"`
	Sortid       int    `json:"sortid"`
	Insertd_at   string `json:"insertd_at"`
	Pubilsh      int    `json:"pubilsh"`
}




const(
	viedo = `<iframe height=450 width=800 src="%s" frameborder=0 allowfullscreen></iframe>`
	pic = `![图片](%s)`
	center = `### **<center>%s</center>**`
)

func (Article *Article)GetContent()bool {
	try := 0
	request := gorequest.New()
	request.Header.Set(`User-Agent`, `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
	if !strings.HasPrefix(Article.Link, config.Base_Url) {
		Article.Link = config.Base_Url + Article.Link
	}
	for try < config.Max_Try {
		proxy := utils.GetProxy()
		res, _, errs := request.Get(Article.Link).Proxy(proxy).End()
		if utils.CheckErrs(errs) || res.StatusCode != 200 {
			try += 1
			continue
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if utils.CheckErr(err) {
			try += 1
			continue
		}
		content := ""
		doc.Find(`#img-content div.rich_media_content p `).Each(func(i int, s *goquery.Selection) {
			if value, ok := s.ChildrenFiltered(`iframe`).Attr(`data-src`); ok {
				content += "\n" + fmt.Sprintf(viedo, value) + "\n"
				return
			} else if value, ok := s.ChildrenFiltered(`img`).Attr(`data-src`); ok {
				content += "\n" + fmt.Sprintf(pic, value) + "\n"
				return
			} else if len(s.Text()) == 1 {
				content += fmt.Sprintf(center, s.Text()) + "\n"
			} else {
				content += "\n"+ s.Text() + "\n"
			}
		})
		Article.Content = content
		return true
	}
	return false
	}







