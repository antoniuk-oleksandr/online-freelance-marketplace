package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"
)

func ResetPassword(
	resetPasswordBody body.ResetPassword,
	email string,
	token string,
) error {
	db := database.GetDB()

	encodedPassword, err := middleware.Encrypt(resetPasswordBody.Password)
	if err != nil {
		return err
	}

	err = repository.ChangeUserPasswordByEmail(encodedPassword, email, db)
	if err != nil{
		return err
	}
	
	return repository.AddJWTToBlacklist(token, db)
}
