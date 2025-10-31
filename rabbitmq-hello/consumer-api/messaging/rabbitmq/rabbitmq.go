package rabbitmq

import (
	"fmt"
	"log"
	"rabbitmq-hello/consumer-api/config"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	conn *amqp.Connection
}

func InitRabbitmq(conf *config.Rabbitmq) (*Rabbitmq, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Pass, conf.Host, conf.Port)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %v", err.Error())
	}

	return &Rabbitmq{
		conn: conn,
	}, nil
}

func (r *Rabbitmq) Close() {
	r.conn.Close()
}

func (r *Rabbitmq) CreateChannel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %v", err.Error())
	}

	return ch, nil
}

func (r *Rabbitmq) DeclareQueue(ch *amqp.Channel, name string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return &amqp.Queue{}, fmt.Errorf("failed to declare queue: %v", err.Error())
	}

	return &q, nil
}

func (r *Rabbitmq) Consume(ch *amqp091.Channel, q *amqp091.Queue) error {
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
		return fmt.Errorf("failed to register a consumer: %v", err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
