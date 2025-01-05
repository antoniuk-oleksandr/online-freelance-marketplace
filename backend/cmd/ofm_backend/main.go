package main

import (
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	order_routes "ofm_backend/cmd/ofm_backend/api/order/routes"
	search_routes "ofm_backend/cmd/ofm_backend/api/search/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	home_data_routes "ofm_backend/cmd/ofm_backend/api/home_data/routes"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvValues()

	db := database.ConnectToPostgresDB()
	database.ConnectToRedisDB()

	app := fiber.New()
	
	app.Use(config.ConfigCors())
	app.Use(config.ConfigRateLimiter())

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup)
	freelance_routes.RegisterFreelanceRoutes(apiGroup, db)
	user_routes.RegisterUserRoutes(apiGroup, db)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	search_routes.RegisterSearchRoutes(apiGroup)
	order_routes.RegisterOrderRoutes(apiGroup)
	home_data_routes.RegisterHomeDataRoutes(apiGroup, db)
	app.Listen(":8080")
}