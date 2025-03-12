package controller

import "github.com/gofiber/fiber/v2"

type PaymentController interface {
	GetPublicKey(ctx *fiber.Ctx) error
	ProcessPayment(ctx *fiber.Ctx) error
}