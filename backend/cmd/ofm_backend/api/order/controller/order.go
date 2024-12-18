package controller

import (
	"ofm_backend/cmd/ofm_backend/api/order/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(ctx *fiber.Ctx) error {
	err := service.CreateOrder()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created",
	})
}
