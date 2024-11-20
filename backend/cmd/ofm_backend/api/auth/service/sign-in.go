package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SignIn(ctx *fiber.Ctx) (string, string, error) {
	var credentials model.Credentials

	if err := ctx.BodyParser(&credentials); err != nil {
		return "", "", err
	}

	db := database.GetDB()

	passwordFromDB, err := repository.GetUserPassword(credentials.Username, db)
	if err != nil {
		return "", "", fiber.ErrUnauthorized
	}

	decryptedPassword, err := middleware.Decrypt(*passwordFromDB)
	if err != nil {
		return "", "", err
	}

	if decryptedPassword != credentials.Password {
		return "", "", fiber.ErrUnauthorized
	}

	accessToken, err := middleware.GenerateAccessToken(credentials.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := middleware.GenerateRefreshToken(credentials.Username)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}