package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/config"
)

const (
	maxRetries    = 5
	retryInterval = 3 * time.Second
)

type Rabbitmq struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(conf *config.RabbitMQ) (*Rabbitmq, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)

	var (
		conn *amqp.Connection
		err  error
	)
	for i := range maxRetries {
		conn, err = amqp.Dial(uri)
		if err == nil {
			break
		}
		log.Printf("rabbitmq: connection attempt %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: failed to connect after %d attempts: %w", maxRetries, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("rabbitmq: failed to open channel: %w", err)
	}

	return &Rabbitmq{conn, ch}, nil
}

func (mq *Rabbitmq) DeclareQueue(name string) (*amqp.Queue, error) {
	q, err := mq.ch.QueueDeclare(
		name,
		true,  // durable: must match user-service declaration
		false, // auto-delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: failed to declare queue %q: %w", name, err)
	}

	return &q, nil
}

func (mq *Rabbitmq) Consume(q *amqp.Queue) (<-chan amqp.Delivery, error) {
	msgs, err := mq.ch.Consume(
		q.Name,          // queue
		"notif-service", // consumer tag — meaningful name for management UI/logs
		false,           // auto-ack: false — ack manually AFTER successful email send
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: failed to consume from queue %q: %w", q.Name, err)
	}

	return msgs, nil
}

// Close shuts down channel first, then connection.
// Closing connection first forcefully tears down the channel.
func (mq *Rabbitmq) Close() {
	if err := mq.ch.Close(); err != nil {
		log.Printf("rabbitmq: error closing channel: %v", err)
	}
	if err := mq.conn.Close(); err != nil {
		log.Printf("rabbitmq: error closing connection: %v", err)
	}
}
