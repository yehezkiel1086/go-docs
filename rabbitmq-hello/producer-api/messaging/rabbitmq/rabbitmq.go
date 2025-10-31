package rabbitmq

import (
	"fmt"
	"rabbitmq-hello/producer-api/config"

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