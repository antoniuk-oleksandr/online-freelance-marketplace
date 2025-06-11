package main

import (
	"log"

	"file-server/config"
	"file-server/internal/api/files/controller"
	"file-server/internal/api/files/service"
	"file-server/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const uploadDir = "./storage/uploads"

func main() {
	app := fiber.New(config.GetAppConfig())
	app.Use(cors.New())
	utils.CreateDirIfDoesNotExist(uploadDir)

	app.Static("/files", uploadDir)

	fileService := service.NewFileService(uploadDir)
	fileController := controller.NewFileController(fileService)

	app.Post("/upload", fileController.UploadFiles)
	app.Delete("/delete/:filename", fileController.DeleteFile)

	log.Fatal(app.Listen(":8030"))
}
