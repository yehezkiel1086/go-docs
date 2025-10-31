package main

import (
	"fmt"
	"log"
	"rabbitmq-hello/consumer-api/config"
	"rabbitmq-hello/consumer-api/messaging/rabbitmq"
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

	// consume
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
