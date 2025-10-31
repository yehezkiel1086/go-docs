package main

import (
	"fmt"
	"rabbitmq-hello/producer-api/config"
	"rabbitmq-hello/producer-api/controller"
	"rabbitmq-hello/producer-api/messaging/rabbitmq"
	"rabbitmq-hello/producer-api/router"
)

func main() {
	// load .env config
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)	
	}
	fmt.Println("✅ .env configs loaded successfully")

	// init rabbitmq
	mq, err := rabbitmq.InitRabbitmq(conf.Rabbitmq)
	if err != nil {
		panic(err)
	}
	defer mq.Close()

	// create channel
	ch, err := mq.CreateChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// declare queue
	q, err := mq.DeclareQueue(ch, "hello")
	if err != nil {
		panic(err)
	}

	// declare controllers
	helloController := controller.InitHelloController(ch, q)

	// declare routes
	r := router.InitRouter(
		helloController,
	)

	// run server
	r.Run(conf.HTTP)
}
