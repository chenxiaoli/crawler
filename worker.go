package main

import (
	"encoding/json"
	"log"

	"fmt"

	"github.com/streadway/amqp"
)

/*
StartPageCrawlWorker 启动一个抓取worker
*/
func StartPageCrawlWorker() {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", RabbitMQ["username"], RabbitMQ["password"], RabbitMQ["host"], RabbitMQ["port"])
	log.Printf("amqp dial:%s", url)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"page-crawler", // name
		true,           // durable
		false,          // delete when usused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var aURL URL
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &aURL)
			crawlPage(aURL)
			log.Printf("Done")
			d.Ack(false)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
