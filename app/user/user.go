package user

import (
	"context"

	"github.com/producer_rabbitmq/events/user"
	userModel "github.com/producer_rabbitmq/model/user"
)

type UserApp interface {
	UserCreate(ctx context.Context, user userModel.User) error
}

type appImpl struct {
	eventUser user.EventUserContract
}

func NewUserApp(eventUser user.EventUserContract) UserApp {
	return &appImpl{
		eventUser: eventUser,
	}
}

func (a *appImpl) UserCreate(ctx context.Context, user userModel.User) error {
	err := a.eventUser.UserCreated(user)
	if err != nil {
		return err
	}
	return nil
}
