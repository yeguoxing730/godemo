package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/streadway/amqp"
	"time"
)

/*
默认点对点模式
*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// 连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 打开一个并发服务器通道来处理消息
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 申明一个队列
	q, err := ch.QueueDeclare(
		"go_task_queue", // name
		true,         // durable  持久性的,如果事前已经声明了该队列，不能重复声明
		false,        // delete when unused
		false,        // exclusive 如果是真，连接一断开，队列删除
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)

	// 发布
	err = ch.Publish(
		"",     // exchange 默认模式，exchange为空
		q.Name,           // routing key 默认模式路由到同名队列，即是task_queue
		false,  // mandatory
		false,
		amqp.Publishing{
			// 持久性的发布，因为队列被声明为持久的，发布消息必须加上这个（可能不用），但消息还是可能会丢，如消息到缓存但MQ挂了来不及持久化。
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = fmt.Sprintf("%s-%v","hello", time.Now())
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}