package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Middleware interface {
	ProcessConfirmPasswordJWT() fiber.Handler
	ProcessResetPasswordJWT() fiber.Handler
	GenerateConfirmPasswordToken(userUUID string) (string, error)
	GenerateResetPasswordJWT(email string) (string, error)
	ProcessRegularJWT() fiber.Handler
	ProcessWebSocketJWT() fiber.Handler
	GenerateRefreshToken(username string, refreshTime int) (string, error)
	GenerateSignInAccessToken(username string, accessTime int) (string, error)
	ParseToken(tokenString string, tokenType int) (jwt.MapClaims, error)
	ProcessRefreshToken(ctx *fiber.Ctx) (string, int, error)
	GenerateSignInAccessTokenWithData(username, avatar string, userId int64, accessTime int) (string, error)
}
