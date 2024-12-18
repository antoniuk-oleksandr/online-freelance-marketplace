package main

import (
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvValues()

	database.ConnectToPostgresDB()
	database.ConnectToRedisDB()

	app := fiber.New()
	app.Use(config.ConfigCors())

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup)
	freelance_routes.RegisterFreelanceRoutes(apiGroup)
	user_routes.RegisterUserRoutes(apiGroup)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	
	app.Listen(":8080")
}
