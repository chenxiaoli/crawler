package models

import (
	"fmt"
	"time"

	"github.com/chenxiaoli/crawler/utils"

	"github.com/chenxiaoli/crawler/storage"

	"log"

	"gopkg.in/mgo.v2/bson"
)

/*
URL ...

{"url":"http://www.jjmmw.com/"}
*/
type URL struct {
	ID              string    `json:"_id"`
	Hash            string    `json:"hash"`
	URL             string    `json:"url"`
	Usages          []string  `json:"usages"` //页面的主要用途
	Domain          string    `json:"domain"`
	Code            string    `json:"code"`
	Method          string    `json:"method"`
	PostData        string    `json:"post_data"`
	Status          string    //new,in,out
	StatusCreatedAt time.Time `json:"status_created_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

/*
AddUsage 增加usage
*/
func (f *URL) AddUsage(usage string) {
	exit := false
	for i := 0; i < len(f.Usages); i++ {
		if usage == f.Usages[i] {
			exit = true
		}
	}
	if exit == false {
		f.Usages = append(f.Usages, usage)
	}
}

func (f *URL) SetHash() {
	if f.URL == "" {
		panic("func (f *Url) GetID f.ID should not be null")
	}
	if f.Method == "" {
		panic("func (f *Url) GetID f.Method should not be null")
	}
	idstr := fmt.Sprintf("%s/%s/%s", f.URL, f.Method, f.PostData)
	f.Hash = utils.StringToHash(idstr)
}

/*
URLPattern url模式下的页面用途, parser的值就是worker 队列的名称，例如：parser-jjmmw.com-fund-detail
*/
type URLPattern struct {
	Parsers []string
	Pattern string
	Domain  string
}

/*
PagePaser 页面解析的情况
*/
type PagePaser struct {
	parser    string
	status    string //new,done,err
	message   string //出错信息，或者其他有用信息
	createdAt time.Time
	updatedAt time.Time
}

/*
Page ...
*/
type Page struct {
	//ID          string   `json:"_id"`
	Hash        string   `json:"hash"` //URL.hash=md5(url+method+PostData),unique
	URL         string   `json:"url"`
	Usages      []string `json:"usages"` //页面的主要用途
	Code        string   `json:"code"`   //标识一个页面的实体ID，比如基金详情页，code=000001,可空
	Domain      string   `json:"domain"`
	Method      string   `json:"method"` //GET or POST
	PostData    string   `json:"post_data"`
	Data        []byte
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/*
Host 允许爬虫抓取的网站
*/
type Host struct {
	Host    string
	Allowed bool
}

/*
PageSaveNote 页面保存成功后，发往队列的通知
*/
type PageSaveNote struct {
	PageHash  string    `json:"page_hash"`
	CreatedAt time.Time `json:"created_at"`
}

/*
GetPage 获取Page
*/
func (f *PageSaveNote) GetPage() (Page, error) {

	session := storage.GetSession()
	dbPage := Page{}
	c := session.DB("findata").C("page")
	err := c.Find(bson.M{"hash": &f.PageHash}).One(&dbPage)
	if err != nil {
		log.Printf("func (f *PageSaveNote) GetPage():%s", err)
	}
	return dbPage, nil
}
