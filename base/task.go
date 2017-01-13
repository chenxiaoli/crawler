package base

import (
	"fmt"
	"log"

	"github.com/chenxiaoli/crawler/config"

	"github.com/streadway/amqp"
)

/*
NewTask 创建一个新的任务
*/
func NewTask(queue string, payload string) {

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQ["username"], config.RabbitMQ["password"], config.RabbitMQ["host"], config.RabbitMQ["port"])
	log.Printf("amqp dial:%s", url)
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
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
	FailOnError(err, "Failed to declare a queue")

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
	FailOnError(err, "Failed to publish a message")
}
