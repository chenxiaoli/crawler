package main

import "time"

/*
URL ...
*/
type URL struct {
	URL      string            `json:"url"`
	Usage    string            `json:"usage"`
	Domain   string            `json:"domain"`
	Code     string            `json:"code"`
	Method   string            `json:"method"`
	PostData map[string]string `json:"post_data"`
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
	URL         string
	Data        []byte
	Usage       string
	Domain      string
	Code        string
	ContentType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
