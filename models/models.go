package models

import "time"

/*
URL ...

{"url":"http://www.jjmmw.com/"}
*/
type URL struct {
	ID              string            `json:"_id"`
	URL             string            `json:"url"`
	Usage           string            `json:"usage"` //页面的主要用途
	Domain          string            `json:"domain"`
	Code            string            `json:"code"`
	Method          string            `json:"method"`
	PostData        map[string]string `json:"post_data"`
	Status          string            //new,in,out
	StatusCreatedAt time.Time         `json:"status_created_at"`
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
	ID          string `json:"_id"`
	URL         string `json:"url"`
	Data        []byte
	Usage       string
	Domain      string
	Code        string
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
	Usage       string    `json:"usage"`
	Code        string    `json:"code"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
