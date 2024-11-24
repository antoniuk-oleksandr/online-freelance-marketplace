package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func ConfirmEmail(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid").(string)

	redisDB := database.GetRedisDB()
	db := database.GetDB()

	user, err := repository.GetUserTempData(uuid, redisDB)
	if err != nil {
		return err
	}

	if err = repository.AddUser(user, db); err != nil {
		return err
	}

	return nil
}
