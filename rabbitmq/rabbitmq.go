package rabbitmq

import (
	"log"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type RabbitMQQueue interface {
	Connect() *RabbitMQ
	SendMessage(exchangePublisher *rabbitmq.Publisher, exchange string, routingKey []string, message []byte) error
}

type RabbitMQ struct {
	Url       string
	Conn      *rabbitmq.Conn
	Exchanges map[string]*rabbitmq.Publisher
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
		log.Fatal(err)
	}

	return &RabbitMQ{
		Conn: conn,
	}
}

func (r *RabbitMQ) RegisterExchanges(exchanges map[string]string) map[string]*rabbitmq.Publisher {
	exchangesPublisher := make(map[string]*rabbitmq.Publisher)
	for exchange := range exchanges {
		publisher, err := rabbitmq.NewPublisher(
			r.Conn,
			rabbitmq.WithPublisherOptionsLogging,
			rabbitmq.WithPublisherOptionsExchangeName(exchange),
			rabbitmq.WithPublisherOptionsExchangeDeclare,
			rabbitmq.WithPublisherOptionsExchangeKind("fanout"),
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
