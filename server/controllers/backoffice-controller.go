package controllers

import (
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/models"
	"github.com/devnica/EasyStore/services"
	"github.com/gofiber/fiber/v2"
)

type backOfficeController struct {
	services.BackofficeService
	configurations.Config
}

func NewBackofficeController(service *services.BackofficeService, config configurations.Config) *backOfficeController {
	return &backOfficeController{
		BackofficeService: *service,
		Config:            config,
	}
}

func (controller backOfficeController) Route(app *fiber.App) {
	app.Get("/easystore/v1/backoffice/kyc-review", controller.GetKYCReview)
}

func (controller backOfficeController) GetKYCReview(c *fiber.Ctx) error {

	reviewStatus := c.Query("status")

	data := controller.BackofficeService.GetKYCReview(c.Context(), reviewStatus)
	return c.Status(fiber.StatusCreated).JSON(models.GeneralHttpResponseModel{
		Code:    200,
		Message: "",
		Data:    data,
	})
}
