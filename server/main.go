package main

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configurations.New()

	app := fiber.New(configurations.NewFiber())

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)
}
