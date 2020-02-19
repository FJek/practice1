package direct

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"practice1/util"
)



func main() {
	// 1 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2 建立信道
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 3_工厂方法 声明 exchange
	err = ch.ExchangeDeclare(
		"logs_direct",   // name
		"direct", // fanout 类型 会广播消息给给所有订阅这个队列的消费者
		true,     // 持久化
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")
	// 4 声明队列
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	// 5 exchange 和 queue 绑定
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,        // queue name
			s,             // routing key
			"logs_direct", // exchange
			false,
			nil)
		util.FailOnError(err, "Failed to bind a queue")
	}
	util.FailOnError(err, "Failed to bind a queue")
	// 6 消费者订阅消息，制定 queue 名称
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	// 6 读取消息并且打印，
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	// forever 一直没有值所以会一直阻塞
	<-forever
}
