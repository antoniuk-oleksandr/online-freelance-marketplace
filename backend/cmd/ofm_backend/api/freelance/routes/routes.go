package routes

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/controller"
	"ofm_backend/cmd/ofm_backend/api/freelance/repository"
	"ofm_backend/cmd/ofm_backend/api/freelance/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterFreelanceRoutes(apiGroup fiber.Router, db *sqlx.DB) {
	repo := repository.NewFreelanceRepository(db)
	freelanceService := service.NewFreelanceService(repo)
	freelanceController := controller.NewFreelanceController(freelanceService)

	freelancesGroup := apiGroup.Group("/freelances")
	freelancesGroup.Get("/:id", freelanceController.GetFreelanceById)
	freelancesGroup.Get("/:id/reviews", freelanceController.GetReviewsByFreelanceId)
	freelancesGroup.Get("/:id/restricted", freelanceController.GetResrictedFreelanceById)
}
