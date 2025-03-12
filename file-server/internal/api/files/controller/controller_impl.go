package controller

import (
	"file-server/internal/api/files/service"
	"file-server/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type fileController struct {
	fileService service.FileService
}

func NewFileController(fileService service.FileService) FileController {
	return &fileController{
		fileService: fileService,
	}
}

func (dc fileController) DeleteFile(ctx *fiber.Ctx) error {
	fileName := ctx.Params("filename", "")
	if fileName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrNoFilenameSpecified.Error(),
		})
	}

	if status, err := dc.fileService.DeleteFile(fileName); err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}

func (fc fileController) UploadFiles(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrNoFileUploaded.Error(),
		})
	}

	files := form.File["files"]
	filesNum := len(files)
	if filesNum == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrNoFileUploaded.Error(),
		})
	}

	if status, err := fc.fileService.UploadFiles(files); err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("%d file%s uploaded successfully", filesNum, utils.Plural(filesNum)),
	})
}
