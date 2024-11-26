package service

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/database"
	filereader "ofm_backend/internal/file_reader"
	"ofm_backend/internal/mailer"
	"ofm_backend/internal/middleware"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

func ForgotPassword(usernameOrEmail string) error {
	db := database.GetDB()

	username, email, err := resolveUserCredentials(usernameOrEmail, db)
	if err != nil {
		return err
	}

	token, err := middleware.GenerateResetPasswordJWT(email)
	if err != nil {
		return err
	}

	link := generateResetLink(token)
	
	html, err := generatePasswordResetEmailHTML(username, link)
	if err != nil {
		return err
	}
	
	return mailer.SendEmail(email, "Password reset", html, "text/html")
}

func generateResetLink(token string) string {
	host := os.Getenv("FRONTEND_HOST")
	port := os.Getenv("FRONTEND_PORT")
	return fmt.Sprintf("http://%s:%s/reset-password?token=%s", host, port, token)
}

func generatePasswordResetEmailHTML(username string, link string) (string, error) {
	html, err := filereader.GetHTMLTempalate("forgot_password.html")
	if err != nil {
		return "", err
	}
	html = strings.Replace(html, "{url}", link, -1)
	html = strings.Replace(html, "{username}", username, -1)

	return html, nil
}

func resolveUserCredentials(usernameOrEmail string, db *sqlx.DB) (string, string, error) {
	if strings.Contains(usernameOrEmail, "@") {
		return validateEmailAndFetchUsername(usernameOrEmail, db)
	} else {
		return validateUsernameAndFetchEmail(usernameOrEmail, db)
	}
}

func validateUsernameAndFetchEmail(username string, db *sqlx.DB) (string, string, error) {
	email, exists, err := repository.GetEmailByUsernameIfExists(username, db)
	if err != nil {
		return "", "", err
	}

	if !exists {
		return "", "", utils.ErrUsernameDoesNotExist
	}

	return username, email, nil
}

func validateEmailAndFetchUsername(email string, db *sqlx.DB) (string, string, error) {
	usernameFromDB, exists, err := repository.GetUsernameByEmailIfExists(email, db)
	if err != nil {
		return "", "", err
	}

	if !exists {
		return "", "", utils.ErrEmailDoesNotExist
	}

	return usernameFromDB, email, nil
}
