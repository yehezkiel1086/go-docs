package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
)

const (
	maxRetries     = 5
	retryInterval  = 3 * time.Second
	publishTimeout = 5 * time.Second
)

type Rabbitmq struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(conf *config.RabbitMQ) (*Rabbitmq, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)

	// retry loop — RabbitMQ frequently starts slower than the app
	// in Docker Compose, so a single Dial attempt will fail on startup.
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

	// create channel
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
		true,  // durable: survives broker restarts — critical for email confirmations
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

func (mq *Rabbitmq) SendJSON(q *amqp.Queue, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), publishTimeout)
	defer cancel()

	err := mq.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent, // message survives broker restart
			Body:         data,
		},
	)
	if err != nil {
		return fmt.Errorf("rabbitmq: failed to publish to queue %q: %w", q.Name, err)
	}

	return nil
}

// close shuts down the channel first, then the connection.
// closing the connection first would forcefully tear down the channel,
// potentially losing in-flight messages.
func (mq *Rabbitmq) Close() {
	if err := mq.ch.Close(); err != nil {
		log.Printf("rabbitmq: error closing channel: %v", err)
	}
	if err := mq.conn.Close(); err != nil {
		log.Printf("rabbitmq: error closing connection: %v", err)
	}
}
