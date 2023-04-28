package user

import (
	"encoding/json"
	"log"

	"github.com/producer_rabbitmq/model/user"
	"github.com/producer_rabbitmq/rabbitmq"
)

type EventUserContract interface {
	UserCreated(user user.User) error
}

type EventUser struct {
	rabbitMQ *rabbitmq.RabbitMQ
}

func Register(rabbitMQ *rabbitmq.RabbitMQ) *EventUser {
	exchanges := map[string]bool{
		"user_created": true,
	}

	for exchange := range exchanges {
		err := rabbitMQ.Rabbit.ExchangeDeclare(
			exchange,
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Failed to declare exchange: %v", err)
			panic(err)
		}
	}

	return &EventUser{
		rabbitMQ: rabbitMQ,
	}
}

func (e *EventUser) UserCreated(user user.User) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = e.rabbitMQ.SendMessage("user_created", jsonData)
	if err != nil {
		return err
	}
	return nil
}
