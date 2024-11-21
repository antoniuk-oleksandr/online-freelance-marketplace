package controller

import (
	"ofm_backend/cmd/ofm_backend/api/auth/service"

	"github.com/gofiber/fiber/v2"
)

func RefreshToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong token",
		})
	}

	accessToken, err := service.RefreshToken(c, tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"accessToken": accessToken,
	})
}
