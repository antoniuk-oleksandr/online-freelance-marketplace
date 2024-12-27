package controller

import (
	"github.com/gofiber/fiber/v2"
)

type FreelanceController interface {
	GetFreelanceById(ctx *fiber.Ctx) error
	GetReviewsByFreelanceID(ctx *fiber.Ctx) error
}
