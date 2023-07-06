package middlewares

import (
	"os"

	"github.com/devnica/EasyStore/models"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(role string) func(*fiber.Ctx) error {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles := claims["roles"].([]interface{})

			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})
				if roleMap["rol"] == role {
					return ctx.Next()
				}
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(models.GeneralHttpResponseModel{
					Code:    401,
					Message: "Unathorized",
					Data:    "Invalid Role",
				})
		},

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(models.GeneralHttpResponseModel{
						Code:    400,
						Message: "Bad request",
						Data:    err.Error(),
					})
			} else {
				return c.
					Status(fiber.StatusForbidden).
					JSON(models.GeneralHttpResponseModel{
						Code:    403,
						Message: "Forbidden",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}
