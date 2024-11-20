package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
