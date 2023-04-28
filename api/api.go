package api

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/producer_rabbitmq/api/user"
	"github.com/producer_rabbitmq/app"
)

func Register(apps *app.Container, port string) {
	fiberRoute := fiber.New()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	go func() {
		<-gracefulShutdown
		fmt.Println("Gracefully shutting down...")
		_ = fiberRoute.Shutdown()
	}()

	baseAPI := fiberRoute.Group("/api")

	// Users
	user.NewAPI(baseAPI.Group("/users"), apps)

	fiberRoute.Get("/swagger/*", swagger.HandlerDefault)

	err := fiberRoute.Listen(":" + port)
	if err != nil {
		panic("error to listen api: " + err.Error())
	}
}
