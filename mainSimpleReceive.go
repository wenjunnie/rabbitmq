package main

import "rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("Simple")
	rabbitmq.ConsumerSimple()
}
