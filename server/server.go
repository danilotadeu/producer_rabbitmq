package server

import (
	"os"
	"os/signal"

	"github.com/producer_rabbitmq/api"
	"github.com/producer_rabbitmq/app"
	"github.com/producer_rabbitmq/events/user"
	"github.com/producer_rabbitmq/rabbitmq"
)

type Server interface {
	Start()
}

func Start() {
	urlRabbitMQ := os.Getenv("URL_RABBITMQ")
	rabbitMQ := rabbitmq.NewRabbitMQ(urlRabbitMQ)

	rabbitMQConnection := rabbitMQ.Connect()
	eventUser := user.Register(rabbitMQConnection)
	apps := app.Register(eventUser)

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	go func() {
		<-gracefulShutdown
		_ = rabbitMQConnection.Conn.Close()
		_ = rabbitMQConnection.Channel.Close()
	}()

	api.Register(apps, "6000")
}
