package main

import (
	"encoding/json"
	"log"

	"github.com/chenxiaoli/crawler/base"
	"github.com/streadway/amqp"
)

/*
SendPageCrawlTask 发起一个网页抓取任务
*/
func SendPageCrawlTask(url URL) {
	conn, err := amqp.Dial("amqp://findata:fax123@localhost:5672/")
	base.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"page-crawler", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	log.Println(q)
	base.FailOnError(err, "Failed to declare a queue")

	body, _ := json.Marshal(url)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	base.FailOnError(err, "Failed to publish a message")
}
