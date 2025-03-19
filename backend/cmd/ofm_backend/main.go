package main

import (
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	file_repo "ofm_backend/cmd/ofm_backend/api/file/repository"
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	home_data_routes "ofm_backend/cmd/ofm_backend/api/home_data/routes"
	my_profile_routes "ofm_backend/cmd/ofm_backend/api/my_profile/routes"
	order_routes "ofm_backend/cmd/ofm_backend/api/order/routes"
	payment_routes "ofm_backend/cmd/ofm_backend/api/payment/routes"
	search_routes "ofm_backend/cmd/ofm_backend/api/search/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/rsa_encryption"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvValues()

	posgresqlDb := database.ConnectToPostgresDB()
	redisDb := database.ConnectToRedisDB()

	middlewareRepository := middleware_repo.NewMiddlewareRepository(posgresqlDb)
	middleware := middleware.NewMiddleware(middlewareRepository)
	rsa_encryption.TryToLoadRSAKeys()

	app := fiber.New()

	app.Use(config.ConfigCors())
	app.Use(config.ConfigRateLimiter())

	fileReposotory := file_repo.NewFileRepository(posgresqlDb)
	fileService := file_service.NewFileService(fileReposotory)

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup, posgresqlDb, redisDb, middleware)
	freelance_routes.RegisterFreelanceRoutes(apiGroup, posgresqlDb)
	user_routes.RegisterUserRoutes(apiGroup, posgresqlDb)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	search_routes.RegisterSearchRoutes(apiGroup)
	home_data_routes.RegisterHomeDataRoutes(apiGroup, posgresqlDb)
	payment_routes.RegisterPaymentRoutes(apiGroup, posgresqlDb, middleware)
	order_routes.RegisterOrderRoutes(apiGroup, posgresqlDb, fileService, middleware)
	my_profile_routes.RegisterMyProfileRoutes(apiGroup, posgresqlDb)

	app.Listen(":8080")
}
