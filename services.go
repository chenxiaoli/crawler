package main

import (
	"findata/funddata/storage"
	"io/ioutil"
	"log"
	"time"

	"github.com/henrylee2cn/surfer"

	"gopkg.in/mgo.v2/bson"
)

func crawlPage(aURL URL) {
	resp, err := surfer.Download(&surfer.DefaultRequest{Url: aURL.URL})
	if err != nil {
		log.Println(err)
	} else {
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println(err)
		} else {
			page := Page{}
			page.URL = aURL.URL
			page.Data = b
			page.ContentType = resp.Header.Get("Content-Type")
			page.Updated = time.Now()
			page.Created = time.Now()
			page.Usage = aURL.Usage
			page.Domain = aURL.Domain
			page.Code = aURL.Code
			savePage(&page)
		}
	}
}

func savePage(p *Page) {
	session := storage.GetSession()
	log.Println(p.URL)
	dbPage := Page{}
	c := session.DB("findata").C("page")

	err := c.Find(bson.M{"url": &p.URL}).One(&dbPage)
	if err == nil {
		p.Created = dbPage.Created
		c.Update(bson.M{"url": &p.URL}, p)
		log.Println("update page:" + p.URL)
	} else {
		err = c.Insert(p)
		if err != nil {
			log.Fatal(err)

		}
	}

}
