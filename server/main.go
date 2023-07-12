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
	storeRepository := repository.NewStoreRepositoryImpl(conn)
	fileRepository := repository.NewFileRepositoryImpl(conn)

	//services
	userAccountService := service.NewUserAccountServiceImpl(&userAccountRepository, &adminCommitRepository, &argon)
	storeService := service.NewStoreServiceImpl(&storeRepository, &userAccountRepository, &fileRepository)
	backofficeService := service.NewBackofficeServiceImpl(&adminCommitRepository)

	// controllers
	authController := controllers.NewAuthController(&userAccountService, config)
	storeController := controllers.NewStoreController(&storeService, config)
	backofficeController := controllers.NewBackofficeController(&backofficeService, config)

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
	storeController.Route(app)
	backofficeController.Route(app)

	// start app
	err := app.Listen(config.Get("SERVER_PORT"))
	exceptions.PanicLogging(err)
}
