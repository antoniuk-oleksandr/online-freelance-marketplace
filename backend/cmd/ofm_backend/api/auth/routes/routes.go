package routes

import (
	"ofm_backend/cmd/ofm_backend/api/auth/controller"
	authRepo "ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/api/auth/service"
	file_repo "ofm_backend/cmd/ofm_backend/api/file/repository"
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(
	apiGroup fiber.Router,
	posgresqlDb *sqlx.DB, redisDb *redis.Client,
	s3Client *s3.Client,
) {
	authGroup := apiGroup.Group("/auth")

	middlewareRepository := middleware_repo.NewMiddlewareRepository(posgresqlDb)
	middleware := middleware.NewMiddleware(middlewareRepository)

	fileReposotory := file_repo.NewFileRepository(posgresqlDb)
	fileService := file_service.NewFileService(fileReposotory, s3Client)

	authRepository := authRepo.NewAuthRepository(posgresqlDb, redisDb)
	authService := service.NewAuthService(authRepository, middleware, fileService)
	authController := controller.NewAuthController(authService)

	authGroup.Post("/sign-in", authController.SignIn)
	authGroup.Post("/sign-up", authController.SignUp)
	authGroup.Post("/forgot-password", authController.ForgotPassword)
	authGroup.Post("/sign-out", authController.SignOut)
	authGroup.Post("/confirm-email", middleware.ProcessConfirmPasswordJWT(), authController.ConfirmEmail)
	authGroup.Post("/reset-password", middleware.ProcessResetPasswordJWT(), authController.ResetPassword)
	authGroup.Get("/session", middleware.ProcessRegularJWT(), authController.Session)
	authGroup.Get("/email-availability", authController.CheckIfEmailIsAvailable)
	authGroup.Post("/google", authController.SignInWithGoogle)
	authGroup.Put("/google", authController.SignUpWithGoogle)
}
