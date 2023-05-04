package user

import (
	"encoding/json"

	"github.com/producer_rabbitmq/model/user"
	"github.com/producer_rabbitmq/rabbitmq"
	rabbit "github.com/wagslane/go-rabbitmq"
)

type EventUserContract interface {
	UserCreated(user user.User) error
}

type EventUser struct {
	rabbitMQ  *rabbitmq.RabbitMQ
	Exchanges map[string]*rabbit.Publisher
}

func NewEvent(rabbitMQ *rabbitmq.RabbitMQ, exchanges map[string]*rabbit.Publisher) EventUserContract {
	return &EventUser{
		Exchanges: exchanges,
		rabbitMQ:  rabbitMQ,
	}
}

func (e *EventUser) UserCreated(user user.User) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = e.rabbitMQ.SendMessage(e.Exchanges["user_created"], "user_created", []string{""}, jsonData)
	if err != nil {
		return err
	}
	return nil
}
