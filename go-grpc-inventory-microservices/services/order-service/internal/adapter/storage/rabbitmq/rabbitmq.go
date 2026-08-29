package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/config"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func New(conf *config.Rabbitmq) (*RabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"notification", // name
		true,           // durability
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn, ch, q}, nil
}

func (r *RabbitMQ) Publish(ctx context.Context, message []byte) error {
	return r.ch.PublishWithContext(ctx, "", r.q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})
}

func (r *RabbitMQ) Close() error {
	if err := r.ch.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}
