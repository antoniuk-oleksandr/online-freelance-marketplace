package main

import (
	auth_routes "ofm_backend/cmd/ofm_backend/api/auth/routes"
	filter_params_routes "ofm_backend/cmd/ofm_backend/api/filter_params/routes"
	freelance_routes "ofm_backend/cmd/ofm_backend/api/freelance/routes"
	order_routes "ofm_backend/cmd/ofm_backend/api/order/routes"
	search_routes "ofm_backend/cmd/ofm_backend/api/search/routes"
	user_routes "ofm_backend/cmd/ofm_backend/api/user/routes"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	utils.LoadEnvValues()

	db := database.ConnectToPostgresDB()
	database.ConnectToRedisDB()
	
	app := Setup(db)
	app.Use(config.ConfigCors())
	app.Use(config.ConfigRateLimiter())
	app.Listen(":8080")
}

func Setup(db *sqlx.DB) *fiber.App{
	app := fiber.New()

	apiGroup := app.Group("/api/v1")
	auth_routes.RegisterAuthRoutes(apiGroup)
	freelance_routes.RegisterFreelanceRoutes(apiGroup, db)
	user_routes.RegisterUserRoutes(apiGroup)
	filter_params_routes.RegisterFilterParamsRoutes(apiGroup)
	search_routes.RegisterSearchRoutes(apiGroup)
	order_routes.RegisterOrderRoutes(apiGroup)
	
	return app
}
