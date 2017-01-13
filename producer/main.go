package main

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

func main() {
	configFile := flag.String("configfile", os.Args[1], "General configuration file")
	config.InitCnf(configFile)
	storage.StartUp()
	StartWorker()
}

/*
StartPageCrawlWorker 启动一个抓取worker
*/
func StartWorker() {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ["username"], config.RabbitMQ["password"], config.RabbitMQ["host"], config.RabbitMQ["port"])
	log.Printf("amqp dial:%s", url)
	conn, err := amqp.Dial(url)
	base.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"page-crawl-done", // name
		true,              // durable
		false,             // delete when usused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	base.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	base.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var note models.PageSaveNote
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &note)
			URLProducer(note)
			log.Printf("Done")
			d.Ack(false)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


