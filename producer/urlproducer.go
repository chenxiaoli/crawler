package producer

import (
	"bytes"
	"encoding/json"
	"time"

	"flag"
	"fmt"
	"log"
	"os"

	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/chenxiaoli/crawler/base"
	"github.com/chenxiaoli/crawler/config"
	"github.com/chenxiaoli/crawler/models"
	"github.com/chenxiaoli/crawler/storage"
	"github.com/chenxiaoli/crawler/utils"
	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2/bson"
)

/*
URLProducer url生成器，把以抓取的页面解析出URL，根据规则丢到队列。
*/

func URLProducer(p models.PageSaveNote) {

	session := storage.GetSession()
	//var url_patterns []URLPattern
	//session.DB("findata").C("url_pattern").Find(bson.M{}).All(&url_patterns)
	c := session.DB("findata").C("page")
	dbPage := models.Page{}
	log.Println(p.URL)
	_id := utils.ToMd5String(p.URL)
	err := c.FindId(_id).One(&dbPage)
	if err != nil {
		panic(err)
	}
	log.Println(dbPage.URL)
	br := bytes.NewReader(dbPage.Data)
	doc, err := goquery.NewDocumentFromReader(br)
	if err != nil {
		log.Fatal(err)
	}

	baseURLString, err := url.Parse(dbPage.URL)
	if err != nil {
		panic(err)
	}

	var urls = make(map[string]*url.URL)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		aurl, b := s.Attr("href")

		if b == true {

			urlStruct, err := url.Parse(aurl)

			if err == nil {
				urlStruct.Fragment = ""
				if urlStruct.IsAbs() == false {
					urlStruct.Host = baseURLString.Host
					urlStruct.Scheme = baseURLString.Scheme
				}
				urls[urlStruct.String()] = urlStruct
			} else {
			log.Printf("error:%s \n", err)
			}
		}

	})

	AllowedHosts := make(map[string]bool)
	var hosts []models.Host
	session.DB("findata").C("host").Find(bson.M{}).All(&hosts)
	for index := 0; index < len(hosts); index++ {
		AllowedHosts[hosts[index].Host] = hosts[index].Allowed
	}

	var ids []string
	var existUrls []models.URL
	for k, _ := range urls {
		ids = append(ids, utils.ToMd5String(k))
	}
	session.DB("findata").C("url").Find(bson.M{"_id": bson.M{"$in": ids}}).All(&existUrls)
	for index := 0; index < len(ids); index++ {
		var notIn = true
		for j := 0; j < len(existUrls); j++ {
			if existUrls[j].URL == ids[index] {
				notIn = false
			}
		}
		if notIn == true {
			delete(urls, ids[index])
		}

	}

	for k, v := range urls { //URL不再数据库中，添加到数据库，如果是需要爬数据，发到爬虫队列。
		log.Println("ff:" + k)
		var aURL models.URL
		_id := utils.ToMd5String(k)

		allowed, ok := AllowedHosts[v.Host]

		aURL.URL = k
		aURL.Domain = v.Host
		aURL.Method = "GET"
		aURL.Status = "new"
		aURL.StatusCreatedAt = time.Now()

		if ok && allowed {
			b, err := json.Marshal(aURL)
			if err == nil {
				base.NewTask("page-crawler", string(b))
				aURL.Status = "in"
				aURL.StatusCreatedAt = time.Now()
			} else {
				log.Println(err)
			}
		}
		err := session.DB("findata").C("url").Insert(bson.M{"_id": _id}, aURL)
		if err != nil {
			panic(err)
		}

	}
}