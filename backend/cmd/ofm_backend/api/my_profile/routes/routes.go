package routes

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/controller"
	my_profile_repo "ofm_backend/cmd/ofm_backend/api/my_profile/repository"
	"ofm_backend/cmd/ofm_backend/api/my_profile/service"
	"ofm_backend/internal/middleware"
	"ofm_backend/internal/middleware/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterMyProfileRoutes(apiGroup fiber.Router, db *sqlx.DB) {
	myProfileGroup := apiGroup.Group("/my-profile")

	myProfileRepository := my_profile_repo.NewMyProfileRepository(db)
	myProfileService := service.NewMyProfileService(myProfileRepository)
	myProfileController := controller.NewMyProfileController(myProfileService)

	middlewareRepository := repository.NewMiddlewareRepository(db)
	middleware := middleware.NewMiddleware(middlewareRepository)

	myProfileGroup.Get("/orders", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileOrders)
	myProfileGroup.Get("/services", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileServices)
	myProfileGroup.Get("/requests", middleware.ProcessRegularJWT(), myProfileController.GetMyProfileRequests)
}
