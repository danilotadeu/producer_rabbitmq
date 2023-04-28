package user

import (
	"context"

	"github.com/producer_rabbitmq/events/user"
	userModel "github.com/producer_rabbitmq/model/user"
)

type UserApp interface {
	SendMessage(ctx context.Context, user userModel.User) error
}

type appImpl struct {
	eventUser *user.EventUser
}

func NewUserApp(eventUser *user.EventUser) UserApp {
	return &appImpl{
		eventUser: eventUser,
	}
}

func (a *appImpl) SendMessage(ctx context.Context, user userModel.User) error {
	err := a.eventUser.UserCreated(user)
	if err != nil {
		return err
	}
	return nil
}
