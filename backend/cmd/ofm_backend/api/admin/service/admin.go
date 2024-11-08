package service

import "github.com/gofiber/fiber/v2"

func DoAdminStuff(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "You are an admin",
		"username": c.Locals("username"),
	})
}
