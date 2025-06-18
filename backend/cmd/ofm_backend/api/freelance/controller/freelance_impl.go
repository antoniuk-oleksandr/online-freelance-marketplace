package controller

import (
	"errors"
	"ofm_backend/cmd/ofm_backend/api/freelance/service"
	"ofm_backend/cmd/ofm_backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type freelanceController struct {
	service service.FreelanceService
}

func NewFreelanceController(
	service service.FreelanceService,
) FreelanceController {
	return &freelanceController{
		service: service,
	}
}

func (fc *freelanceController) GetResrictedFreelanceById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	freelanceByID, err := fc.service.GetResrictedFreelanceById(id)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(freelanceByID)
}

func (fc *freelanceController) GetFreelanceById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	freelanceByID, err := fc.service.GetFreelanceById(id)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(freelanceByID)
}

func (fc *freelanceController) GetReviewsByFreelanceId(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	reviewsCursor := ctx.Query("reviewsCursor", "")

	reviews, err := fc.service.GetReviewsByFreelanceID(id, reviewsCursor)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(reviews)
}
