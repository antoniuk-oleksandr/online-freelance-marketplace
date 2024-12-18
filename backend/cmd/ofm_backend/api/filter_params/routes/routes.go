package routes

import (
	"github.com/gofiber/fiber/v2"
	"ofm_backend/cmd/ofm_backend/api/filter_params/controller"
)

func RegisterFilterParamsRoutes(apiGroup fiber.Router){
	filterParamsGroup := apiGroup.Group("/filter-params")
	
	filterParamsGroup.Get("/get-all", controller.FilterParamsGetAll)
}