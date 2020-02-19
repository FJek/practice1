package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {
	// 2 连接 amqp server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"fail to connect rabbitmq")
	defer conn.Close()
	// 3_工厂方法 创建channel
	ch, err := conn.Channel()
	failOnError(err,"fail to create channel")
	defer ch.Close()
	// 4 创建队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)

	err = ch.Publish(
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}


