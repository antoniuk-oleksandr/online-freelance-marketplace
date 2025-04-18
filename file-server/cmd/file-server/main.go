

import (
	"log"

	"file-server/config"
	"file-server/internal/api/files/controller"
	"file-server/internal/api/files/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(config.GetAppConfig())
	app.Use(cors.New())
	
	bucketName := "online-freelance-marketplace"
	fileService := service.NewFileService(bucketName)
	fileController := controller.NewFileController(fileService)

	app.Post("/upload", fileController.UploadFiles)
	app.Delete("/delete/:filename", fileController.DeleteFile)

	log.Fatal(app.Listen(":8030"))
}
