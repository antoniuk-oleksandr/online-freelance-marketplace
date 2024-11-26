package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/api/auth/utils"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"
	fileRepo "ofm_backend/cmd/ofm_backend/api/file/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GoogleAuth(ctx *fiber.Ctx, code string) (string, string, error) {
	db := database.GetDB()

	tokenResponse, err := utils.ExchangeAuthCodeForToken(code)
	if err != nil {
		return "", "", err
	}

	claims, err := utils.ParseJWT(tokenResponse.IdToken)
	if err != nil {
		return "", "", err
	}

	username, err := handleUserAccount(claims, db)
	if err != nil {
		return "", "", err
	}

	accessToken, fefreshToken, err := middleware.GenerateTokens(username)
	if err != nil {
		return "", "", err
	}

	return accessToken, fefreshToken, nil
}

func handleUserAccount(claims *body.GoogleJwtClaims, db *sqlx.DB) (string, error) {
	username, userExists, err := repository.GetUsernameByEmailIfExists(claims.Email, db)
	if err != nil {
		return "", err
	}

	if userExists {
		return username, nil
	}
	
	avatarID, err := fileRepo.AddFile(claims.PicURL, db)
	if err != nil {
		return "", err
	}
	
	err = repository.AddUserWithGoogleAuth(claims, avatarID, db)
	if err != nil {
		return "", err
	}

	return claims.Email, nil
}
