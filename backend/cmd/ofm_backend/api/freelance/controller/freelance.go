package controller

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/service"

	"github.com/gofiber/fiber/v2"
)

func GetFreelanceById(c *fiber.Ctx) error {
	return service.GetFreelanceById(c)
}
