package service

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/mapper"
	"ofm_backend/cmd/ofm_backend/api/freelance/repository"
	"ofm_backend/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetFreelanceById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid freelance service ID",
		})
	}
	db := database.GetDB()

	freelanceService, err := repository.GetFreelanceServiceById(id, db)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Freelance service not found",
		})
	}
	
	freelanceServiceReviews, err :=repository.GetFreelanceServiceByIdReviews(id, db)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Freelance service reviews not found",
		})
	}

	free := mapper.MapFreelanceModelToDTO(freelanceService, freelanceServiceReviews)
	
	return c.JSON(free)
}
