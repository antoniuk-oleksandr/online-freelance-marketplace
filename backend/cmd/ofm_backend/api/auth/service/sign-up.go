package service

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/database"
	filereader "ofm_backend/internal/file_reader"
	"ofm_backend/internal/mailer"
	"ofm_backend/internal/middleware"
	"os"
	"strings"

	"github.com/google/uuid"
)

func SignUp(user *body.SignUpBody) error {
	db := database.GetDB()
	redisDB := database.GetRedisDB()

	encryptedPassword, err := middleware.Encrypt(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword

	isUsernameAvailable, err := repository.CheckIfUsernameIsAvailable(user.Username, db)
	if err != nil {
		return err
	}

	if !isUsernameAvailable {
		return utils.ErrUsernameIsTaken
	}
	
	isEmailAvailable, err := repository.CheckIfEmailIsAvailable(user.Email, db)
	if err != nil {
		return err
	}
	
	if !isEmailAvailable {
		return utils.ErrEmailIsTaken
	}

	userUUID := uuid.New().String()
	err = repository.AddTempUserData(user, userUUID, redisDB)
	if err != nil {
		return err
	}
	
	token, err := middleware.GenerateConfirmPasswordToken(userUUID)
	if err != nil {
		return err
	}
	
	host := os.Getenv("FRONTEND_HOST")
	port := os.Getenv("FRONTEND_PORT")
	link := fmt.Sprintf("http://%s:%s/confirm-email?token=%s", host, port, token)
	
	html, err := filereader.GetHTMLTempalate("confirm_email.html")
	if err != nil {
		return err
	}
	html = strings.Replace(html, "{url}", link, -1)
	html = strings.Replace(html, "{username}", user.Username, -1)
	
	err = mailer.SendEmail(user.Email, "Password confirmation", html, "text/html")
	if err != nil {
		return err
	} 
	
	return nil
}