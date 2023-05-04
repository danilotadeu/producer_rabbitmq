package server

import (
	"os"
	"os/signal"

	"github.com/producer_rabbitmq/api"
	"github.com/producer_rabbitmq/app"
	eventUser "github.com/producer_rabbitmq/events/user"
	"github.com/producer_rabbitmq/rabbitmq"
)

type Server interface {
	Start()
}

var exchanges = map[string]string{
	"user_created": "fanout",
}

func Start() {
	urlRabbitMQ := os.Getenv("URL_RABBITMQ")
	rabbitMQ := rabbitmq.NewRabbitMQ(urlRabbitMQ)
	rabbitMQConnection := rabbitMQ.Connect()

	exchangesPublisher := rabbitMQConnection.RegisterExchanges(exchanges)
	userEvent := eventUser.NewEvent(rabbitMQConnection, exchangesPublisher)

	apps := app.Register(userEvent)

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	go func() {
		<-gracefulShutdown
		for _, publisher := range exchangesPublisher {
			publisher.Close()
		}
		_ = rabbitMQConnection.Conn.Close()
	}()

	api.Register(apps, "6000")
}
