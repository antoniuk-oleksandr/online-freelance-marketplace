package routes

import (
	"ofm_backend/cmd/ofm_backend/api/user/controller"
	"ofm_backend/cmd/ofm_backend/api/user/repository"
	"ofm_backend/cmd/ofm_backend/api/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(apiGroup fiber.Router, db *sqlx.DB) {
	usersGroup := apiGroup.Group("/users")

	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)
	contr := controller.NewUserController(serv)

	usersGroup.Get("/:id", contr.GetUserById)
	usersGroup.Get("/:id/reviews", contr.GetReviewsByUserId)
	usersGroup.Get("/:id/services", contr.GetServicesByUserId)
}
