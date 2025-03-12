package main

import (
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	file_repo "ofm_backend/cmd/ofm_backend/api/file/repository"
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	home_data_routes "ofm_backend/cmd/ofm_backend/api/home_data/routes"
	order_routes "ofm_backend/cmd/ofm_backend/api/order/routes"
	payment_routes "ofm_backend/cmd/ofm_backend/api/payment/routes"
	search_routes "ofm_backend/cmd/ofm_backend/api/search/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvValues()
	middleware.TryToLoadRSAKeys()

	db := database.ConnectToPostgresDB()
	database.ConnectToRedisDB()

	app := fiber.New()

	app.Use(config.ConfigCors())
	app.Use(config.ConfigRateLimiter())

	fileReposotory := file_repo.NewFileRepository(db)
	fileService := file_service.NewFileService(fileReposotory)

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup)
	freelance_routes.RegisterFreelanceRoutes(apiGroup, db)
	user_routes.RegisterUserRoutes(apiGroup, db)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	search_routes.RegisterSearchRoutes(apiGroup)
	home_data_routes.RegisterHomeDataRoutes(apiGroup, db)
	payment_routes.RegisterPaymentRoutes(apiGroup, db)
	order_routes.RegisterOrderRoutes(apiGroup, db, fileService)

	app.Listen(":8080")
}
