package controllers

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/models"
	"github.com/devnica/EasyStore/models/request"
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
	app.Post("/easystore/v1/auth/login", controller.UserLogin)
	app.Put("/easystore/v1/user/:userId", controller.UpdatePersonalInfo)
}

func (controller authController) UserRegister(c *fiber.Ctx) error {
	var request request.UserAccountRegisterRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	controller.UserAccountService.UserRegister(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    201,
		Message: "User has been created successfully",
		Data:    "",
	})
}

func (controller authController) UserLogin(c *fiber.Ctx) error {
	var request request.UserAccountLoginRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	user := controller.UserAccountService.GetUserByEmail(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    200,
		Message: "Login successfully",
		Data:    user,
	})
}

func (controller authController) UpdatePersonalInfo(c *fiber.Ctx) error {
	var request request.UpdatePersonalInfoRequestModel
	err := c.BodyParser(&request)
	exceptions.PanicLogging(err)

	userId := c.Params("userId")

	controller.UserAccountService.UpdatePersonalInfo(c.Context(), request, userId)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    201,
		Message: "Update personal info",
		Data:    "",
	})
}
