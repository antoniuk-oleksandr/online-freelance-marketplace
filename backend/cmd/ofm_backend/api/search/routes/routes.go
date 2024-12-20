package routes

import (
	"ofm_backend/cmd/ofm_backend/api/search/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterSearchRoutes(apiGroup fiber.Router){
	searchGroup := apiGroup.Group("/search")
	
	searchGroup.Get("", controller.SearchFreelances)
}