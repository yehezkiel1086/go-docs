package rabbitmq

import (
	"context"
	"fmt"
	"go-rabbitmq-work-queues/config"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	conn *amqp.Connection
	ch *amqp.Channel
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

	return &Rabbitmq{
		conn: conn,
		ch: ch,
	}, nil
}

func (mq *Rabbitmq) DeclareQueue(name string) (*amqp.Queue, error) {
	q, err := mq.ch.QueueDeclare(
		name, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err.Error())
	}

	return &q, nil
}

func (mq *Rabbitmq) PublishMessage(q *amqp.Queue, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mq.ch.PublishWithContext(ctx,
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
	}); err != nil {
		return fmt.Errorf("failed to publish a message: %v", err.Error())
	}
	log.Printf(" [x] Sent %s", body)

	return nil
}

func (mq *Rabbitmq) Close() {
	mq.conn.Close()
	mq.ch.Close()
}
