package controller

import (
	"errors"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"

	"github.com/gofiber/fiber/v2"
)

func SignUp(ctx *fiber.Ctx) error {
	var user body.SignUpBody

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body.",
		})
	}

	token, err := service.SignUp(&user)
	if err != nil {
		if errors.Is(err, fiber.ErrConflict) {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Username already exists.",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create account.",
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
