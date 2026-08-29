package rabbitmq

import (
	"bytes"
	"fmt"
	"log"
	"rabbitmq-work-queues_worker-1/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	conn *amqp.Connection
	ch *amqp.Channel
	q *amqp.Queue
}

func InitRabbitmq(conf *config.RabbitMQ) (*Rabbitmq, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Pass, conf.Host, conf.Port)
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %v", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err.Error())
	}

	q, err := ch.QueueDeclare(
		"message", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err.Error())
	}

	if err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		return nil, fmt.Errorf("failed to set QoS: %v", err.Error())
	}

	return &Rabbitmq{
		conn: conn,
		ch: ch,
		q: &q,
	}, nil
}

func (mq *Rabbitmq) Consume() (error) {
	msgs, err := mq.ch.Consume(
		mq.q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message (worker 1): %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

func (mq *Rabbitmq) Close() {
	mq.conn.Close()
	mq.ch.Close()
}
