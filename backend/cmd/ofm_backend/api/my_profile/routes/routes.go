package routes

import (
	file_repo "ofm_backend/cmd/ofm_backend/api/file/repository"
	"ofm_backend/cmd/ofm_backend/api/file/service"
	my_profile_controller "ofm_backend/cmd/ofm_backend/api/my_profile/controller"
	my_profile_repo "ofm_backend/cmd/ofm_backend/api/my_profile/repository"
	my_profile_service "ofm_backend/cmd/ofm_backend/api/my_profile/service"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterMyProfileRoutes(apiGroup fiber.Router, db *sqlx.DB, s3Client *s3.Client) {
	myProfileGroup := apiGroup.Group("/my-profile")

	fileRepository := file_repo.NewFileRepository(db)
	fileService := service.NewFileService(fileRepository, s3Client)

	myProfileRepository := my_profile_repo.NewMyProfileRepository(db)
	myProfileService := my_profile_service.NewMyProfileService(myProfileRepository, fileService)
	myProfileController := my_profile_controller.NewMyProfileController(myProfileService)

	middlewareRepository := middleware_repo.NewMiddlewareRepository(db)
	middleware := middleware.NewMiddleware(middlewareRepository)

	myProfileGroup.Get("/orders", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrders)
	myProfileGroup.Get("/services", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileServices)
	myProfileGroup.Get("/requests", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileRequests)
	myProfileGroup.Get("/orders/:orderId/chat", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrderChat)
	myProfileGroup.Get("/orders/:orderId/overview", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrderOverview)
	myProfileGroup.Get("/orders/:orderId/requirements", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrderRequirements)
	myProfileGroup.Get("/orders/:orderId/delivery", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrderDelivery)
	myProfileGroup.Post("/orders/:orderId/delivery", middleware.ProcessRegularJWT(), myProfileController.AddMyProfileOrderDelivery)
	myProfileGroup.Post("/orders/:orderId/review", middleware.ProcessRegularJWT(), myProfileController.AddMyProfileOrderReview)
	myProfileGroup.Get("/orders/:orderId/review", myProfileController.GetMyProfileOrderReview)
	myProfileGroup.Put("/orders/:orderId/status/:statusId", middleware.ProcessRegularJWT(), myProfileController.UpdateMyProfileOrderStatus)
}
