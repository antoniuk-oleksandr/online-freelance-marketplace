package routes

import (
	"ofm_backend/cmd/ofm_backend/api/order/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterOrderRoutes(apiGroup fiber.Router) {
	ordersGroup := apiGroup.Group("/orders")

	ordersGroup.Post("/", controller.CreateOrder)
	ordersGroup.Get("/public-payment-key", controller.GetPublicPaymentKey)
}
