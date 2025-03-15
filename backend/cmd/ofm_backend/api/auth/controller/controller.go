package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	ConfirmEmail(ctx *fiber.Ctx) error
	ForgotPassword(ctx *fiber.Ctx) error
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	GoogleAuth(ctx *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
}
