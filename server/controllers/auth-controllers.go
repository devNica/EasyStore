package controllers

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/dto/requests"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/models"
	"github.com/devnica/EasyStore/services"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	services.UserAccountService
	configurations.Config
}

func NewAuthController(service *services.UserAccountService, config configurations.Config) *authController {
	return &authController{UserAccountService: *service, Config: config}
}

func (controller authController) Route(app *fiber.App) {
	app.Post("/easystore/v1/auth/register", controller.UserRegister)
}

func (controller authController) UserRegister(c *fiber.Ctx) error {
	var request requests.UserAccountRegisterRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserAccountService.UserRegister(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    201,
		Message: "User has been created successfully",
		Data:    "",
	})
}
