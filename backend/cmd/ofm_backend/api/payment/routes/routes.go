package routes

import (
	"ofm_backend/cmd/ofm_backend/api/payment/controller"
	"ofm_backend/cmd/ofm_backend/api/payment/repository"
	"ofm_backend/cmd/ofm_backend/api/payment/service"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"
	
	
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterPaymentRoutes(apiGroup fiber.Router, db *sqlx.DB) {
	middlewareRepository := middleware_repo.NewMiddlewareRepository(db)
	middleware := middleware.NewMiddleware(middlewareRepository)
	
	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentController := controller.NewPaymentController(paymentService)

	paymentGroup := apiGroup.Group("/payment")

	paymentGroup.Get("/public-key", paymentController.GetPublicKey)
	paymentGroup.Post("/process-payment", middleware.ProcessRegularJWT(), paymentController.ProcessPayment)
}
