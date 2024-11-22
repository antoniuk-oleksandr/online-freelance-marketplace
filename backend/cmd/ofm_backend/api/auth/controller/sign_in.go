package controller

import (
	"ofm_backend/cmd/ofm_backend/api/auth/service"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	accessToken, refreshToken, err := service.SignIn(c)

	if err != nil {
		if err.Error() == fiber.ErrUnprocessableEntity.Error() {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(
				fiber.Map{"error": "Invalid request body"},
			)
		}

		if err.Error() == fiber.ErrUnauthorized.Error() {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid credentials"},
			)
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "An unexpected error occurred"},
		)
	}

	return c.JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
