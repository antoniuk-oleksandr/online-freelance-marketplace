package config

import (
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
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://online-freelance-marketplace.xyz, https://www.online-freelance-marketplace.xyz",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	})
}

// func ConfigRateLimiter() func(ctx *fiber.Ctx) error {
// 	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
// 	maxConnections, _ := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))

// 	host := os.Getenv("REDIS_HOST")
// 	password := os.Getenv("REDIS_PASSWORD")

// 	return limiter.New(limiter.Config{
// 		Max:        maxConnections,
// 		Expiration: 1 * time.Minute,
// 		Storage: redis.New(redis.Config{
// 			Host:      host,
// 			Password:  password,
// 			Port:      port,
// 			TLSConfig: &tls.Config{},
// 		}),
// 		KeyGenerator: func(ctx *fiber.Ctx) string {
// 			return ctx.IP()
// 		},
// 		LimitReached: func(ctx *fiber.Ctx) error {
// 			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
// 				"error": utils.ErrTooManyRequests.Error(),
// 			})
// 		},
// 	})
// }

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