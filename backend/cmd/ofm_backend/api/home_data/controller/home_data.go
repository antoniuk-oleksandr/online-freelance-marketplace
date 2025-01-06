package controller

import "github.com/gofiber/fiber/v2"

type HomeController interface {
	GetHomeData(ctx *fiber.Ctx) error
}
