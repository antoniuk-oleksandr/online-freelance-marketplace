package middleware

import (
	enums "ofm_backend/cmd/ofm_backend/api/auth/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if strings.ToLower(err.Error()) == jwt.ErrTokenExpired.Error() {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		if strconv.Itoa(int(claims["type"].(float64))) != strconv.Itoa(int(enums.Access)) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		c.Locals("username", claims["username"])

		return c.Next()
	}
}

func GenerateToken(username string, tokenType int, experation int64) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"type":     tokenType,
		"exp":      experation,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateAccessToken(username string) (string, error) {
	return GenerateToken(username, enums.Access, time.Now().Add(time.Minute*15).Unix())
}

func GenerateRefreshToken(username string) (string, error) {
	return GenerateToken(username, enums.Refresh, time.Now().Add(time.Hour*24*30).Unix())
}
