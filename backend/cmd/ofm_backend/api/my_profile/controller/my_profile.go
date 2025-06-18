package controller

import "github.com/gofiber/fiber/v2"

type MyProfileController interface {
	GetMyProfileOrders(ctx *fiber.Ctx) error
	GetMyProfileServices(ctx *fiber.Ctx) error
	GetMyProfileRequests(ctx *fiber.Ctx) error
	GetMyProfileOrderChat(ctx *fiber.Ctx) error
	GetMyProfileOrderOverview(ctx *fiber.Ctx) error
	GetMyProfileOrderRequirements(ctx *fiber.Ctx) error
	GetMyProfileOrderDelivery(ctx *fiber.Ctx) error
	AddMyProfileOrderDelivery(ctx *fiber.Ctx) error
	AddMyProfileOrderReview(ctx *fiber.Ctx) error
	GetMyProfileOrderReview(ctx *fiber.Ctx) error
	UpdateMyProfileOrderStatus(ctx *fiber.Ctx) error
}
