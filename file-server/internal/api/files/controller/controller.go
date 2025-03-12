package controller

import "github.com/gofiber/fiber/v2"

type FileController interface {
	DeleteFile(ctx *fiber.Ctx) error
	UploadFiles(ctx *fiber.Ctx) error
}
