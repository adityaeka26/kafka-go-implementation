package rest

import (
	"fmt"

	"test-kraft/consumer/config"

	"github.com/gofiber/fiber/v2"
)

func ServeREST(config *config.EnvConfig) error {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

	return nil
}
