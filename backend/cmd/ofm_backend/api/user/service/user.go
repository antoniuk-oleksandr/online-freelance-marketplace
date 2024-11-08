package service

import (
	"ofm_backend/cmd/ofm_backend/api/user/mapper"
	"ofm_backend/cmd/ofm_backend/api/user/repository"
	"ofm_backend/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	db := database.GetDB()

	userModel, err := repository.GetUserById(id, db)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	userByIdReviewsModels, err := repository.GetUserByIdReviews(id, db)
	reviewsUsers := mapper.MapUserByIdReviewModelsToReviewUserDTOs(userByIdReviewsModels)
	reviewsServices := mapper.MapUserByIdReviewModelToReviewServiceDTOs(userByIdReviewsModels)
	reviews := mapper.MapReviewUsersServicesToUserByIdReviewDTOs(
		userByIdReviewsModels,
		reviewsUsers,
		reviewsServices,
	)

	services, err := repository.GetUserServicesByUserId(id, db)

	userDto := mapper.MapUserByIdModelToDTO(
		userModel,
		reviews,
		services,
	)

	return c.JSON(fiber.Map{
		"user": userDto,
	})
}
