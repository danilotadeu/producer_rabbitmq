package rabbitmq

import (
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type RabbitMQQueue interface {
	Connect() *RabbitMQ
	SendMessage(exchangePublisher *rabbitmq.Publisher, exchange string, routingKey []string, message []byte) error
}

type RabbitMQ struct {
	Url  string
	Conn *rabbitmq.Conn
}

func NewRabbitMQ(url string) RabbitMQQueue {
	return &RabbitMQ{
		Url: url,
	}
}

func (r *RabbitMQ) Connect() *RabbitMQ {
	conn, err := rabbitmq.NewConn(
		r.Url,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		panic(err)
	}

	return &RabbitMQ{
		Conn: conn,
	}
}

func (r *RabbitMQ) RegisterExchanges(exchanges map[string]string) map[string]*rabbitmq.Publisher {
	exchangesPublisher := make(map[string]*rabbitmq.Publisher)
	for exchange, exchangeKind := range exchanges {
		publisher, err := rabbitmq.NewPublisher(
			r.Conn,
			rabbitmq.WithPublisherOptionsLogging,
			rabbitmq.WithPublisherOptionsExchangeName(exchange),
			rabbitmq.WithPublisherOptionsExchangeDeclare,
			rabbitmq.WithPublisherOptionsExchangeKind(exchangeKind),
			rabbitmq.WithPublisherOptionsExchangeDurable,
		)
		if err != nil {
			panic(err)
		}
		exchangesPublisher[exchange] = publisher
	}

	return exchangesPublisher
}

func (r *RabbitMQ) SendMessage(exchangePublisher *rabbitmq.Publisher, exchange string, routingKey []string, message []byte) error {
	err := exchangePublisher.Publish(
		message,
		routingKey,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(exchange),
	)
	if err != nil {
		return err
	}
	return nil
}
