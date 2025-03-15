package routes

import (
	"ofm_backend/cmd/ofm_backend/api/auth/controller"
	authRepo "ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	"ofm_backend/internal/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(
	apiGroup fiber.Router,
	posgresqlDb *sqlx.DB, redisDb *redis.Client,
	middleware middleware.Middleware,
) {
	authGroup := apiGroup.Group("/auth")

	authRepository := authRepo.NewAuthRepository(posgresqlDb, redisDb)
	authService := service.NewAuthService(authRepository, middleware)
	authController := controller.NewAuthController(authService)

	authGroup.Post("/sign-in", authController.SignIn)
	authGroup.Post("/sign-up", authController.SignUp)
	authGroup.Post("/google", authController.GoogleAuth)
	authGroup.Post("/refresh-token", authController.RefreshToken)
	authGroup.Post("/forgot-password", authController.ForgotPassword)
	authGroup.Post("/sign-out", authController.SignOut)
	authGroup.Post("/confirm-email", middleware.ProcessConfirmPasswordJWT(), authController.ConfirmEmail)
	authGroup.Post("/reset-password", middleware.ProcessResetPasswordJWT(), authController.ResetPassword)
}
