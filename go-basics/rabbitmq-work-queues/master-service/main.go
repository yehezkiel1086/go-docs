package main

import (
	"fmt"
	"go-rabbitmq-work-queues/config"
	"go-rabbitmq-work-queues/controller"
	"go-rabbitmq-work-queues/messaging/rabbitmq"
	"go-rabbitmq-work-queues/router"
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

	// controllers dependency injections
	ctl := controller.InitController(*mq)

	// init router
	r := router.InitRouter(
		ctl,
	)
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}

	// serve api
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
