package main

import "time"

/*
URL ...
*/
type URL struct {
	URL    string            `json:"url"`
	Usage  string            `json:"usage"`
	Domain string            `json:"domain"`
	Code   string            `json:"code"`
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
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
	Created     time.Time
	Updated     time.Time
	ParseError  bool
}
