package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SignIn(ctx *fiber.Ctx) (string, string, error) {
	var signInBody body.SignInBody
	
	if err := ctx.BodyParser(&signInBody); err != nil {
		return "", "", err
	}

	db := database.GetDB()

	usernamePassword, err := repository.GetUserPassword(signInBody.UsernameOrEmail, db)
	if err != nil {
		return "", "", fiber.ErrUnauthorized
	}

	decryptedPassword, err := middleware.Decrypt(usernamePassword.Password)
	if err != nil {
		return "", "", err
	}

	if decryptedPassword != signInBody.Password {
		return "", "", fiber.ErrUnauthorized
	}

	accessToken, err := middleware.GenerateAccessToken(usernamePassword.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := middleware.GenerateRefreshToken(usernamePassword.Username)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}