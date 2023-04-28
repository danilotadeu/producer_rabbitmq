package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQQueue interface {
	Connect() *RabbitMQ
	SendMessage(exchange string, message []byte) error
}

type RabbitMQ struct {
	Url     string
	Rabbit  *amqp.Channel
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) RabbitMQQueue {
	return &RabbitMQ{
		Url: url,
	}
}

func (r *RabbitMQ) Connect() *RabbitMQ {
	conn, err := amqp.Dial(r.Url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		panic(err)
	}

	return &RabbitMQ{
		Rabbit:  ch,
		Conn:    conn,
		Channel: ch,
	}
}

func (r *RabbitMQ) SendMessage(exchange string, message []byte) error {
	err := r.Rabbit.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
		return err
	}
	return nil
}
