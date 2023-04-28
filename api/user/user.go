package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/producer_rabbitmq/app"
	"github.com/producer_rabbitmq/model/user"
)

type apiImpl struct {
	userApp *app.Container
}

func NewAPI(g fiber.Router, t *app.Container) {
	api := apiImpl{
		userApp: t,
	}

	g.Post("/", api.createUser)
}

func (p *apiImpl) createUser(c *fiber.Ctx) error {
	request := user.User{}
	c.BodyParser(&request)
	ctx := c.Context()

	err := p.userApp.User.UserCreate(ctx, request)
	if err != nil {
		return c.JSON("Por favor tente novamente mais tarde")
	}

	return c.JSON("Usuário criado com sucesso!")
}
