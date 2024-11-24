package controller

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func GoogleAuth(ctx *fiber.Ctx) error {
	var googleBody body.Google

	if err := ctx.BodyParser(&googleBody); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	if googleBody.Code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	accessToken, refreshToken, err := service.GoogleAuth(ctx, googleBody.Code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
