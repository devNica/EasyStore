package configurations

import "github.com/gofiber/fiber/v2"

func NewFiber() fiber.Config {
	return fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // limit upload file 5mb
	}
}
