package controller

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func ForgotPassword(ctx *fiber.Ctx) error {
	var forgotPasswordBody body.ForgotPassword

	if err := ctx.BodyParser(&forgotPasswordBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	if err := service.ForgotPassword(forgotPasswordBody.UsernameOrEmail); err != nil {
		if err == utils.ErrUsernameDoesNotExist {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ErrUsernameDoesNotExist.Error(),
			})
		}
		
		if err == utils.ErrEmailDoesNotExist {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ErrEmailDoesNotExist.Error(),
			})
		}
		
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The email was sent successfully",
	})
}
