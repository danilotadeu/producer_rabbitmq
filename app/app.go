package app

import (
	userApp "github.com/producer_rabbitmq/app/user"
	userEvent "github.com/producer_rabbitmq/events/user"
)

type Container struct {
	User userApp.UserApp
}

func Register(userEvent *userEvent.EventUser) *Container {
	return &Container{
		User: userApp.NewUserApp(userEvent),
	}
}
