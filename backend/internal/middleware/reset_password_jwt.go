package middleware

import (
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/utils"
	enums "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/database"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ProcessResetPasswordJWT() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")

		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Unauthorized"},
			)
		}

		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if strings.ToLower(err.Error()) == jwt.ErrTokenExpired.Error() {
				return ctx.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{"error": err.Error()},
				)
			} else {
				return ctx.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{"error": "Invalid token"},
				)
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid token"},
			)
		}

		if strconv.Itoa(int(claims["type"].(float64))) != strconv.Itoa(int(enums.ResetPassword)) ||
			claims["email"] == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"error": "Invalid token"},
			)
		}

		db := database.GetDB()
		exists, err := repository.CheckIfTokenBacklisted(tokenString[7:], db)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": utils.ErrUnexpectedError.Error(),
			})
		}
		
		if exists {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": utils.ErrBlacklistedToken.Error(),
			})
		}

		ctx.Locals("email", claims["email"].(string))
		ctx.Locals("token", tokenString[7:])

		return ctx.Next()
	}
}

func GenerateResetPasswordJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"type":  enums.ResetPassword,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
