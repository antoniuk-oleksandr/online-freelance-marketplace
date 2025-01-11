package controller

import (
	"github.com/gofiber/fiber/v2"
)

type FreelanceController interface {
	GetFreelanceById(ctx *fiber.Ctx) error
	GetReviewsByFreelanceId(ctx *fiber.Ctx) error
	GetResrictedFreelanceById(ctx *fiber.Ctx) error
}
