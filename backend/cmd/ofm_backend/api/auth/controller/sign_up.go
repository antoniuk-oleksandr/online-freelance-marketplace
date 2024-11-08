package controller

import (
	"errors"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var user body.SignUpBody

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	accessToken, refreshToken, err := service.SignUp(&user)
	if err != nil {
		if errors.Is(err, utils.ErrUserAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "User already exists with this username",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error adding user",
		})
	}

	return c.JSON(fiber.Map{
		"accessToken":   accessToken,
		"refreshToken": refreshToken,
	})
}
