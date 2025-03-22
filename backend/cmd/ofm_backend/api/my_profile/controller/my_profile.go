package controller

import "github.com/gofiber/fiber/v2"

type MyProfileController interface {
	GetMyProfileOrders(ctx *fiber.Ctx) error
	GetMyProfileServices(ctx *fiber.Ctx) error
	GetMyProfileRequests(ctx *fiber.Ctx) error
}
