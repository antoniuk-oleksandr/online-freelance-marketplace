package controller

import (
	"errors"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func ConfirmEmail(ctx *fiber.Ctx) error {
	err := service.ConfirmEmail(ctx)
	if err != nil {
		if errors.Is(err, utils.ErrTempTokenExpired) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "The token has expired",
			})
		}
		
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}	
	
	return ctx.JSON(fiber.Map{
		"message":  "Email confirmed",
	})
}