package body

import "github.com/gofiber/fiber/v2"

type OrderRequirementsForm struct {
	Answers         string            `json:"answers"`
	CustomerMessage string            `form:"customerMessage"`
	Files           []*fiber.FormFile `form:"files"`
}
