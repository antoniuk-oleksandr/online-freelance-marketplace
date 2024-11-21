package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func ConfirmEmail(ctx *fiber.Ctx) error{
	username :=	ctx.Locals("username").(string)

	redisDB := database.GetRedisDB()
	db := database.GetDB()
	
	user, err := repository.GetUserTempData(username, redisDB)
	if err != nil {
		return err
	}
	
	if err = repository.AddUser(user, db); err != nil {
		return err
	}

	return nil
}
