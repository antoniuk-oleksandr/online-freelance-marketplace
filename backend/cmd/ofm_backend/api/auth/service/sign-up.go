package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SignUp(user *body.SignUpBody) (string, error) {
	db := database.GetDB()
	redisDB := database.GetRedisDB()

	encryptedPassword, err := middleware.Encrypt(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = encryptedPassword

	available, err := repository.CheckIfUsernameIsAvailable(user.Username, db)
	if err != nil {
		return "", err
	}

	if !available {
		return "", fiber.ErrConflict
	}

	token, err := middleware.GenerateAccessToken(user.Username)
	if err != nil {
		return "", err
	}

	err = repository.AddTempUserData(user, redisDB)
	if err != nil {
		return "", err
	}
	
	return token, nil
}