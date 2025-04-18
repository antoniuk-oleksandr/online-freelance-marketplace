package controller

import "github.com/gofiber/fiber/v2"

type ChatController interface {
	ConnectToWS(ctx *fiber.Ctx) error
}
