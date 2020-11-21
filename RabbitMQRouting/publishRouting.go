package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

//路由模式：一个消息被多个消费者获取，并且消息的目标队列可被生产者指定
func main()  {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc","imooc_one")
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc","imooc_two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishRouting("Hello imooc one!" + strconv.Itoa(i))
		imoocTwo.PublishRouting("Hello imooc Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
