package middleware

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	ProcessConfirmPasswordJWT() fiber.Handler
	ProcessResetPasswordJWT() fiber.Handler
	GenerateConfirmPasswordToken(userUUID string) (string, error)
	GenerateResetPasswordJWT(email string) (string, error)
	ProcessRegularJWT() fiber.Handler
	ProcessWebSocketJWT() fiber.Handler
	GenerateRefreshToken(username string) (string, error)
	GenerateSignInAccessToken(username string) (string, error)
}
