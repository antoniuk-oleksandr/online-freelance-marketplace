package routes

import (
	"ofm_backend/cmd/ofm_backend/api/auth/controller"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(apiGroup fiber.Router) {
	authGroup := apiGroup.Group("/auth")
	
	authGroup.Post("/sign-in", controller.SignIn)
	authGroup.Post("/sign-up", controller.SignUp)
	authGroup.Post("/refresh-token", controller.RefreshToken)
	authGroup.Post("/confirm-email", middleware.ProcessConfirmPasswordJWT(), controller.ConfirmEmail)
}