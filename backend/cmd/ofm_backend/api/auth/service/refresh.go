package service

import (
	"ofm_backend/internal/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RefreshToken(c *fiber.Ctx, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fiber.ErrUnauthorized
	}

	if claims["type"] != "refresh" {
		return "", fiber.ErrUnauthorized
	}
	
	

	return middleware.GenerateAccessToken(claims["username"].(string))
}
