package routes

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/controller"
	"ofm_backend/cmd/ofm_backend/api/home_data/repository"
	"ofm_backend/cmd/ofm_backend/api/home_data/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterHomeDataRoutes(
	apiGroup fiber.Router, db *sqlx.DB,
){
	homeRepository := repository.NewHomeRepository(db)
	homeService := service.NewHomeService(homeRepository)
	homeController := controller.NewHomeController(homeService)
	
	homeGroup := apiGroup.Group("/home-data")
	homeGroup.Get("", homeController.GetHomeData)
}