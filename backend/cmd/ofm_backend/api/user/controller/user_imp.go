package controller

import (
	"errors"
	"ofm_backend/cmd/ofm_backend/api/user/service"
	"ofm_backend/cmd/ofm_backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (us *userController) GetReviewsByUserId(ctx *fiber.Ctx) error {
	cursor := ctx.Query("cursor", "")
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	reviews, err := us.userService.GetReviewsByUserId(id, cursor)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(reviews)
}

func (us *userController) GetServicesByUserId(ctx *fiber.Ctx) error {
	cursor := ctx.Query("cursor", "")
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	services, err := us.userService.GetServicesByUserId(id, cursor)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(services)
}

func (us *userController) GetUserById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	userByIDResponse, err := us.userService.GetUserById(id)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(userByIDResponse)
}
