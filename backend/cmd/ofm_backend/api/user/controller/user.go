package controller

import (
	"ofm_backend/cmd/ofm_backend/api/user/service"

	"github.com/gofiber/fiber/v2"
)

func GetUserById(c *fiber.Ctx) error {
	return service.GetUserById(c)
}
