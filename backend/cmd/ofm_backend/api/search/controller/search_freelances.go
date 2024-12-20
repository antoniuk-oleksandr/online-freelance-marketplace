package controller

import (
	"ofm_backend/cmd/ofm_backend/api/search/body"
	"ofm_backend/cmd/ofm_backend/api/search/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func SearchFreelances(ctx *fiber.Ctx) error {
	searchBody := body.Search{}

	if err := ctx.QueryParser(&searchBody); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	freelanceDTO, err := service.SearchFreelances(searchBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(freelanceDTO)
}
