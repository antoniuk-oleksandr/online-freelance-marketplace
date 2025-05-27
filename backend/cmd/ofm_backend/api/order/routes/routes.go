package routes

import (
	file_repo "ofm_backend/cmd/ofm_backend/api/file/repository"
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/cmd/ofm_backend/api/order/controller"
	order_repo "ofm_backend/cmd/ofm_backend/api/order/repository"
	order_service "ofm_backend/cmd/ofm_backend/api/order/service"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterOrderRoutes(
	apiGroup fiber.Router,
	db *sqlx.DB,
	s3Client *s3.Client,
) {
	middlewareRepository := middleware_repo.NewMiddlewareRepository(db)
	middleware := middleware.NewMiddleware(middlewareRepository)

	fileReposotory := file_repo.NewFileRepository(db)
	fileService := file_service.NewFileService(fileReposotory, s3Client)

	orderRepository := order_repo.NewOrderRepository(db)
	orderService := order_service.NewOrderService(orderRepository, fileService)
	orderController := controller.NewOrderController(orderService)

	ordersGroup := apiGroup.Group("/orders")

	ordersGroup.Get("/:id/freelance-questions", middleware.ProcessRegularJWT(), orderController.GetOrderById)
	ordersGroup.Post("/:id/requirements", middleware.ProcessRegularJWT(), orderController.SubmitOrderRequirements)
}
