package controller

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func ResetPassword(ctx *fiber.Ctx) error {
	var resetPasswordBody body.ResetPassword
	if err := ctx.BodyParser(&resetPasswordBody); err != nil {
		return ctx.JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	email := ctx.Locals("email").(string)
	token := ctx.Locals("token").(string)
		
	err := service.ResetPassword(resetPasswordBody, email, token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Password reset successfully!",
	})
}
