package controllers

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/middlewares"
	"github.com/devnica/EasyStore/models"
	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/services"
	"github.com/gofiber/fiber/v2"
)

type storeController struct {
	services.StoreService
	configurations.Config
}

func NewStoreController(service *services.StoreService, config configurations.Config) *storeController {
	return &storeController{StoreService: *service, Config: config}
}

func (controller storeController) Route(app *fiber.App) {
	app.Post("/easystore/v1/store/:userId", controller.RegisterStore)
	app.Get("/easystore/v1/store", middlewares.AuthenticateJWT("owners"), controller.GetStoreByOwnerId)
}

func (controller storeController) RegisterStore(c *fiber.Ctx) error {
	var request request.StoreRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	userId := c.Params("userId")

	controller.StoreService.RegisterStore(c.Context(), request, userId)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    201,
		Message: "successfull store registration",
		Data:    "",
	})
}

func (controller storeController) GetStoreByOwnerId(c *fiber.Ctx) error {

	ownerId := c.Locals("userId").(string)

	result := controller.StoreService.GetStoresByOwnerId(c.Context(), ownerId)
	return c.Status(fiber.StatusOK).JSON(models.GeneralHttpResponseModel{
		Code:    200,
		Message: "request success",
		Data:    result,
	})
}
