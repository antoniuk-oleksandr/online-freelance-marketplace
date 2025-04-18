package middleware

import (
	"ofm_backend/cmd/ofm_backend/utils"
	enums "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/middleware/repository"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type middleware struct {
	repo repository.MiddlewareRepository
}

func NewMiddleware(middlewareRepository repository.MiddlewareRepository) Middleware {
	return &middleware{
		repo: middlewareRepository,
	}
}

func (middleware *middleware) ProcessResetPasswordJWT() fiber.Handler {
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

		exists, err := middleware.repo.CheckIfTokenBacklisted(tokenString[7:])
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

func (middleware *middleware) GenerateResetPasswordJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"type":  enums.ResetPassword,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (middleware *middleware) ProcessConfirmPasswordJWT() fiber.Handler {
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

func (middleware *middleware) GenerateConfirmPasswordToken(userUUID string) (string, error) {
	claims := jwt.MapClaims{
		"uuid": userUUID,
		"type": enums.ConfirmEmail,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (middleware *middleware) ProcessRegularJWT() fiber.Handler {
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


		c.Locals("username", claims["username"].(string))
		c.Locals("avatar", claims["avatar"].(string))
		c.Locals("userId", claims["userId"].(float64))

		return c.Next()
	}
}

func (middleware *middleware) ProcessWebSocketJWT() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Query("token")
		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if strings.ToLower(err.Error()) == jwt.ErrTokenExpired.Error() {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
			} else {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
			}
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if strconv.Itoa(int(claims["type"].(float64))) != strconv.Itoa(int(enums.Access)) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}


		ctx.Locals("username", claims["username"].(string))
		ctx.Locals("avatar", claims["avatar"].(string))
		ctx.Locals("userId", claims["userId"].(float64))

		return ctx.Next()
	}
}

func (middleware *middleware) GenerateRefreshToken(username string) (string, error) {
	refreshTime, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if err != nil {
		return "", err
	}

	expiration := time.Now().Add(time.Minute * time.Duration(refreshTime)).Unix()
	claims := jwt.MapClaims{
		"username": username,
		"type":     enums.Refresh,
		"exp":      expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (middleware *middleware) GenerateSignInAccessToken(username string) (string, error) {
	accessTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if err != nil {
		return "", err
	}

	expiration := time.Now().Add(time.Minute * time.Duration(accessTime)).Unix()

	userSignInData, err := middleware.repo.GetUserSignInDataByUsername(username)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"username": username,
		"type":     enums.Access,
		"exp":      expiration,
		"avatar":   *utils.AddServerURLToFiles[*string](&userSignInData.Avatar),
		"userId":   userSignInData.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
