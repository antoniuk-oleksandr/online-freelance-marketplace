package controller

import (
	"ofm_backend/cmd/ofm_backend/api/filter_params/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func FilterParamsGetAll(ctx *fiber.Ctx) error {
	filterParams, err := service.FilterParamsGetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(filterParams)
}
