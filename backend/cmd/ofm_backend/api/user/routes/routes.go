package routes

import (
	"ofm_backend/cmd/ofm_backend/api/user/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(apiGroup fiber.Router){
	usersGroup := apiGroup.Group("/users")
	
	usersGroup.Get("/:id", controller.GetUserById)
}