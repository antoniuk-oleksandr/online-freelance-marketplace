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

func ProcessConfirmPasswordJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Unauthorized"},
			)
		}

		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if strings.ToLower(err.Error()) == jwt.ErrTokenExpired.Error() {
				return c.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{"error": err.Error()},
				)
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{"error": "Invalid token"},
				)
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid token"},
			)
		}

		if strconv.Itoa(int(claims["type"].(float64))) != strconv.Itoa(int(enums.ConfirmEmail)) {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid token"},
			)
		}

		c.Locals("uuid", claims["uuid"])

		return c.Next()
	}
}

func GenerateConfirmPasswordToken(userUUID string) (string, error) {
	claims := jwt.MapClaims{
		"uuid": userUUID,
		"type": enums.ConfirmEmail,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
