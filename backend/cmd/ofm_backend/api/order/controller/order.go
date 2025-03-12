package controller

import (
	"ofm_backend/cmd/ofm_backend/api/order/body"

	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	GetOrderById(ctx *fiber.Ctx) error
	SubmitOrderRequirements(ctx *fiber.Ctx) error
	getRequiredFormValue(name string, ctx *fiber.Ctx) (string, error)
	parseSubmitOrderRequirementsRequestBody(ctx *fiber.Ctx) (*body.OrderRequirementsBody, error)
	getRequiredFormIntValue(name string, ctx *fiber.Ctx) (int, error)
}
