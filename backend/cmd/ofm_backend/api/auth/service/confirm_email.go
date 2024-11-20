package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ConfirmEmail(ctx *fiber.Ctx) (string, string, error){
	username :=	ctx.Locals("username").(string)

	redisDB := database.GetRedisDB()
	db := database.GetDB()
	
	user, err := repository.GetUserTempData(username, redisDB)
	if err != nil {
		return "", "", err
	}
	
	if err = repository.AddUser(user, db); err != nil {
		return "", "", err
	}

	accessToken, err := middleware.GenerateAccessToken(user.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := middleware.GenerateRefreshToken(user.Username)
	
	return accessToken, refreshToken, nil
}
