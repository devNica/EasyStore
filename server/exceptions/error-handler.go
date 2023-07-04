package exceptions

import (
	"encoding/json"

	"github.com/devnica/EasyStore/models"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)

	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)

		return ctx.Status(fiber.StatusBadRequest).JSON(models.GeneralHttpResponseModel{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, badRequestError := err.(BadRequestError)

	if badRequestError {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.GeneralHttpResponseModel{
			Code:    400,
			Message: "Bad Request",
			Data:    err.Error(),
		})
	}

	_, notFoundError := err.(NotFoundError)

	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(models.GeneralHttpResponseModel{
			Code:    404,
			Message: "Not Found Request Error",
			Data:    err.Error(),
		})
	}

	_, unAuthorizedError := err.(UnAuthorizedError)

	if unAuthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.GeneralHttpResponseModel{
			Code:    401,
			Message: "UnAuthorized Request Error",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(models.GeneralHttpResponseModel{
		Code:    500,
		Message: "Internal Server Error",
		Data:    err.Error(),
	})
}
