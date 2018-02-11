package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

/*
广播模式
发布方
*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 默认模式有默认交换机，广播自己定义一个交换机，交换机可与队列进行绑定
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type 广播模式
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args)

	// 发布
	err = ch.Publish(
		"logs", // exchange 消息发送到交换机，这个时候没队列绑定交换机，消息会丢弃
		"",     // routing key  广播模式不需要这个，它会把所有消息路由到绑定的所有队列
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
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