package rest

import (
	"fmt"

	"kafka-go-implementation-consumer/config"

	"github.com/gofiber/fiber/v2"
)

func ServeREST(config *config.EnvConfig) error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

	return nil
}
