package models

import "time"

/*
URL ...

{"url":"http://www.jjmmw.com/"}
*/
type URL struct {
	ID              string            `json:"_id"`
	URL             string            `json:"url"`
	Usages          []string          `json:"usages"` //页面的主要用途
	Domain          string            `json:"domain"`
	Code            string            `json:"code"`
	Method          string            `json:"method"`
	PostData        map[string]string `json:"post_data"`
	Status          string            //new,in,out
	StatusCreatedAt time.Time         `json:"status_created_at"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
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
	ID       string            `json:"_id"` //_id=URL._id
	URL      string            `json:"url"`
	Usages   []string          `json:"usages"` //页面的主要用途
	Code     string            `json:"code"`   //标识一个页面的实体ID，比如基金详情页，code=000001,可空
	Domain   string            `json:"domain"`
	Method   string            `json:"method"` //GET or POST
	PostData map[string]string `json:"post_data"`

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
	URL         string    `json:"url"`
	Usages      []string  `json:"usages"`
	Code        string    `json:"code"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
