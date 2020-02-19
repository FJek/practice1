package direct

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"practice1/util"
	"strings"
)



func main(){
	// 1 连接
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
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")
	// 4 从命令行读取信息
	body := bodyFrom(os.Args)
	// 5 发布消息
	err = ch.Publish(
		"logs_direct", // exchange
		severityFrom(os.Args),     // routing key,从命令行获取
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
// 从命令行读取信息
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
// 从命令行获取routing key
func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:]," ")
	}
	return s
}

