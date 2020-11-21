package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("Simple")
	rabbitmq.PublishSimple("Hello World!")
	fmt.Println("发送成功！")
}
