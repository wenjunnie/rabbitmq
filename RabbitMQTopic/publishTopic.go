package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

//话题模式：一个消息被多个消费者获取，并且消息的目标队列可用BindingKey以通配符(#:一个或多个词，*:一个词)的方式指定
func main()  {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.topic.one")
	imoocTwo := RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.topic.two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishTopic("Hello imooc topic one!" + strconv.Itoa(i))
		imoocTwo.PublishTopic("Hello imooc topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
