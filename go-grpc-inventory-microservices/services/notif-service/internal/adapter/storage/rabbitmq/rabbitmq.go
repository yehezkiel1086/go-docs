package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/config"
)

type Rabbitmq struct {
	q    *amqp.Queue
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(conf *config.Rabbitmq) (*Rabbitmq, error) {
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

	return &Rabbitmq{&q, conn, ch}, nil
}

func (r *Rabbitmq) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.ch.Consume(
		r.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *Rabbitmq) Close() error {
	if err := r.ch.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}
