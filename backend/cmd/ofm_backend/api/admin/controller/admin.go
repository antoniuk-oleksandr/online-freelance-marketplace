package controller

import (
	"ofm_backend/cmd/ofm_backend/api/admin/service"

	"github.com/gofiber/fiber/v2"
)

func DoAdminStuff(c *fiber.Ctx) error {
	return service.DoAdminStuff(c)
}
