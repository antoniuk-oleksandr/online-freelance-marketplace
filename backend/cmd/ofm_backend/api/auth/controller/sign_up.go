package controller

import (
	"errors"
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"

	"github.com/gofiber/fiber/v2"
)

func SignUp(ctx *fiber.Ctx) error {
	var user body.SignUpBody

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := service.SignUp(&user)
	if err != nil {
		if errors.Is(err, fiber.ErrConflict) {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "User already exists",
			})
		}

		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create account",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The email was sent successfully",
	})
}
