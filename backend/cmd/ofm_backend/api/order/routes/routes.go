package routes

import (
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/cmd/ofm_backend/api/order/controller"
	order_repo "ofm_backend/cmd/ofm_backend/api/order/repository"
	order_service "ofm_backend/cmd/ofm_backend/api/order/service"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterOrderRoutes(apiGroup fiber.Router, db *sqlx.DB, fileService file_service.FileService) {
	ordersGroup := apiGroup.Group("/orders")

	orderRepository := order_repo.NewOrderRepository(db)
	orderService := order_service.NewOrderService(orderRepository, fileService)
	orderController := controller.NewOrderController(orderService)

	ordersGroup.Get("/:id/freelance-questions", middleware.ProcessRegularJWT(), orderController.GetOrderById)
	ordersGroup.Post("/:id/requirements", middleware.ProcessRegularJWT(), orderController.SubmitOrderRequirements)
}
