package config

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis"
)

func ConfigCors() func(*fiber.Ctx) error {
	host := os.Getenv("FRONTEND_HOST")
	port := os.Getenv("FRONTEND_PORT")

	return cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("http://%s:%s", host, port),
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	})
}

func ConfigRateLimiter() func(ctx *fiber.Ctx) error {
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
    maxConnections, _ := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
    
	host := os.Getenv("REDIS_HOST")
    password := os.Getenv("REDIS_PASSWORD")

	return limiter.New(limiter.Config{
		Max: maxConnections,
		Expiration: 1 * time.Minute,
		Storage: redis.New(redis.Config{
			Host: host,
			Password: password,
			Port: port,
		}),
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": utils.ErrTooManyRequests.Error(),
			})
		},
	})
}
