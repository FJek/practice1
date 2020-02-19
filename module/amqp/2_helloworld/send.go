package __helloworld

import (
	"github.com/streadway/amqp"
	"log"
)

// 1
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 2 连接 amqp server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"fail to connect rabbitmq")
	defer conn.Close()
	// 3_工厂方法 创建channel
	channel, err := conn.Channel()
	failOnError(err,"fail to create channel")
	defer channel.Close()
	// 4 创建队列
	queue, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// 5 创建消息,以字节数组方式传送
	body := "Hello World!"
	// 6 发布消息
	err = channel.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
