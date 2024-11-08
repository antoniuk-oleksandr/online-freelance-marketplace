package main

import (
	user_controller "ofm_backend/cmd/ofm_backend/api/user/controller"
	auth_controller "ofm_backend/cmd/ofm_backend/api/auth/controller"
	freelance_controller "ofm_backend/cmd/ofm_backend/api/freelance/controller"
	admin_controller "ofm_backend/cmd/ofm_backend/api/admin/controller"
	"ofm_backend/internal/config"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.Init()

	app := fiber.New()
	app.Use(config.ConfigCors())

	apiGroup := app.Group("/api/v1")
	protectedGroup := apiGroup.Group("/admins", middleware.JWTProtected())

	apiGroup.Get("/users/:id", user_controller.GetUserById)
	apiGroup.Get("/services/:id", freelance_controller.GetFreelanceById)
	apiGroup.Post("/sign-in", auth_controller.SignIn)
	apiGroup.Post("/sign-up", auth_controller.SignUp)
	apiGroup.Post("/refresh-token", auth_controller.RefreshToken)

	protectedGroup.Get("/", admin_controller.DoAdminStuff)

	app.Listen(":8080")
}
