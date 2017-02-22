package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/chenxiaoli/crawler/base"
	"github.com/chenxiaoli/crawler/config"
	"github.com/chenxiaoli/crawler/models"
	"github.com/chenxiaoli/crawler/storage"
	"github.com/streadway/amqp"

	"github.com/henrylee2cn/surfer"

	"gopkg.in/mgo.v2/bson"
)

func crawlPage(aURL models.URL) {
	resp, err := surfer.Download(&surfer.DefaultRequest{Url: aURL.URL})
	if err != nil {
		log.Println(err)
	} else {
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println(err)
		} else {
			urlStruct, _ := url.Parse(aURL.URL)
			page := models.Page{}
			page.URL = aURL.URL
			page.Data = b
			page.ContentType = resp.Header.Get("Content-Type")
			page.UpdatedAt = time.Now()
			page.CreatedAt = time.Now()
			page.Usages = aURL.Usages
			page.Domain = urlStruct.Host
			page.Code = aURL.Code
			savePage(&page)
		}
	}
}

func savePage(p *models.Page) {
	session := storage.GetSession()
	log.Println(p.URL)
	dbPage := models.Page{}
	c := session.DB("findata").C("page")
	urlc := session.DB("findata").C("url")
	urlc.Update(bson.M{"url": &p.URL}, bson.M{"status": "out", "status_created_at": time.Now()})
	err := c.Find(bson.M{"url": &p.URL}).One(&dbPage)
	if err == nil {
		p.CreatedAt = dbPage.CreatedAt
		c.Update(bson.M{"url": &p.URL}, p)
		log.Println("update page:" + p.URL)

		var note models.PageSaveNote
		note.Code = p.Code
		note.ContentType = p.ContentType
		note.CreatedAt = p.CreatedAt
		note.URL = p.URL
		note.Usages = p.Usages
		b, err := json.Marshal(note)
		if err == nil {
			if len(p.Usages) > 0 {
				for index := 0; index < len(p.Usages); index++ {
					name := fmt.Sprintf("page-crawl-done/%s", p.Usages[index])
					NewTask(name, string(b))
				}

			} else {
				NewTask("page-crawl-done", string(b))
			}

		}
	} else {
		err = c.Insert(p)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("insert page:" + p.URL)
		var note models.PageSaveNote
		note.Code = p.Code
		note.ContentType = p.ContentType
		note.CreatedAt = p.CreatedAt
		note.URL = p.URL
		note.Usages = p.Usages
		b, err := json.Marshal(note)
		if err == nil {
			NewTask("page-crawl-done", string(b))
		}
	}

}

/*
NewTask 创建一个新的任务
*/
func NewTask(queue string, payload string) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ["username"], config.RabbitMQ["password"], config.RabbitMQ["host"], config.RabbitMQ["port"])
	log.Printf("amqp dial:%s", url)
	conn, err := amqp.Dial(url)
	base.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	log.Println(q)
	base.FailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(payload),
		})
	log.Printf(" [x] Sent %s", payload)
	base.FailOnError(err, "Failed to publish a message")
}
