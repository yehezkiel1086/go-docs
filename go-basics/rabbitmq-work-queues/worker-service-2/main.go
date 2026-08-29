package main

import (
	"fmt"
	"rabbitmq-work-queues_worker-2/config"
	"rabbitmq-work-queues_worker-2/messaging/rabbitmq"
)

func main() {
	// load .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs loaded successfully")

	// init rabbitmq
	mq, err := rabbitmq.InitRabbitmq(conf.RabbitMQ)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ RabbitMQ connected successfully")
	defer mq.Close()

	// consume messages
	if err := mq.Consume(); err != nil {
		panic(err)
	}
}
