package main

import "rabbitmq/RabbitMQ"

func main()  {
	imoocAll := RabbitMQ.NewRabbitMQTopic("exImoocTopic","#")
	imoocAll.ReceiveTopic()
}
