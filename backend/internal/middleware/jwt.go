package middleware

import (
	enums "ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ProcessRegularJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if strings.ToLower(err.Error()) == jwt.ErrTokenExpired.Error() {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if strconv.Itoa(int(claims["type"].(float64))) != strconv.Itoa(int(enums.Access)) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("username", claims["username"])

		return c.Next()
	}
}

func GenerateOneToken(username string, tokenType int, experation int64) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"type":     tokenType,
		"exp":      experation,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateAccessToken(username string) (string, error) {
	return GenerateOneToken(username, enums.Access, time.Now().Add(time.Hour*24*365).Unix())
}

func GenerateRefreshToken(username string) (string, error) {
	return GenerateOneToken(username, enums.Refresh, time.Now().Add(time.Hour*24*30).Unix())
}

func GenerateTokens(username string) (string, string, error) {
	accessToken, err := GenerateAccessToken(username)
	if err != nil {
		return "", "", err
	}
	
	refreshToken, err := GenerateRefreshToken(username)
	if err != nil {
		return "", "", err
	}
	
	return accessToken, refreshToken, nil
}
