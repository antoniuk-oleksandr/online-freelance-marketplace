package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) (string, string, error) {
	var credentials model.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return "", "", err
	}

	db := database.GetDB()

	valid, err := repository.ValidateUser(&credentials, db)

	if err != nil || !valid {
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

func SignUp(user *body.SignUpBody) (string, string, error) {
	db := database.GetDB()

	if err := repository.AddUser(user, db); err != nil {
		return "", "", err
	}

	accessToken, err := middleware.GenerateAccessToken(user.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := middleware.GenerateRefreshToken(user.Username)

	return accessToken, refreshToken, err
}
