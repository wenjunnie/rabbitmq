package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

//工作模式：一个消息只能被一个消费者获取
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("Simple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello World!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
