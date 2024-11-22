package routes

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterFreelanceRoutes(apiGroup fiber.Router) {
	freelancesGroup := apiGroup.Group("/freelances")
	
	freelancesGroup.Get("/services/:id", controller.GetFreelanceById)
}
