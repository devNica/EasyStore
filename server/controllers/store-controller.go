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
	app.Post("/easystore/v1/store", middlewares.AuthenticateJWT("customers"), controller.RegisterStore)
	app.Get("/easystore/v1/store", middlewares.AuthenticateJWT("owners"), controller.GetStoreByOwnerId)
	app.Put("/easystore/v1/store/:storeId", middlewares.AuthenticateJWT("owners"), controller.UpdateStore)
}

func (controller storeController) RegisterStore(c *fiber.Ctx) error {
	var request request.StoreRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	userId := c.Locals("userId").(string)

	result := controller.StoreService.RegisterStore(c.Context(), request, userId)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    201,
		Message: "successfull store registration",
		Data:    result,
	})
}

func (controller storeController) GetStoreByOwnerId(c *fiber.Ctx) error {

	ownerId := c.Locals("userId").(string)

	// fmt.Println(c.Locals("user"))

	result := controller.StoreService.GetStoresByOwnerId(c.Context(), ownerId)
	return c.Status(fiber.StatusOK).JSON(models.GeneralHttpResponseModel{
		Code:    200,
		Message: "request success",
		Data:    result,
	})
}

func (controller storeController) UpdateStore(c *fiber.Ctx) error {

	var request request.UpdateStoreRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)
	storeId := c.Params("storeId")

	// fmt.Println(c.Locals("user"))

	controller.StoreService.UpdateStoreInfoByStoreId(c.Context(), storeId, request)
	return c.Status(fiber.StatusAccepted).JSON(models.GeneralHttpResponseModel{
		Code:    202,
		Message: "Store has been updated successfull",
		Data:    "",
	})
}
