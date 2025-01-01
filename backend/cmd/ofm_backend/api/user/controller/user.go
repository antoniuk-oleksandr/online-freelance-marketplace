package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GetUserById(ctx *fiber.Ctx) error
	GetReviewsByUserId(ctx *fiber.Ctx) error
	GetServicesByUserId(ctx *fiber.Ctx) error
}
