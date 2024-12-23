package controller

import (
	"ofm_backend/cmd/ofm_backend/api/order/service"

	"github.com/gofiber/fiber/v2"
)

func GetPublicPaymentKey(ctx *fiber.Ctx) error {
	key, err := service.GetPaymentPublicKey() 
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"key": key,
	})
}