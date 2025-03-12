package config

import "github.com/gofiber/fiber/v2"

const maxUploadSize = 50 * 1024 * 1024 // 50 MB

func GetAppConfig() fiber.Config {
	return fiber.Config{
		BodyLimit: maxUploadSize,
	}
}
