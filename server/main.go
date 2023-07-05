package main

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/controllers"
	"github.com/devnica/EasyStore/exceptions"
	repository "github.com/devnica/EasyStore/repositories/impl"
	service "github.com/devnica/EasyStore/services/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configurations.New()
	conn := configurations.DatabaseConnect(config)
	argon := configurations.NewArgonConfig()

	// repositories
	userAccountRepository := repository.NewUserAccountRepositoryImpl(conn)
	adminCommitRepository := repository.NewAdminCommitRepositoryIMpl(conn)

	//services
	userAccountService := service.NewUserAccountServiceImpl(&userAccountRepository, &adminCommitRepository, &argon)

	// controllers
	authController := controllers.NewAuthController(&userAccountService, config)

	app := fiber.New(configurations.NewFiber())

	//middlewares
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	// app.Use(csrf.New())

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	//routing
	authController.Route(app)

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)
}
