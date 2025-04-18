package service

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/helpers"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/aes_encryption"
	"ofm_backend/cmd/ofm_backend/utils/bcrypt_encryption"
	filereader "ofm_backend/internal/file_reader"
	"ofm_backend/internal/mailer"
	"ofm_backend/internal/middleware"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type authService struct {
	authRepository repository.AuthRepository
	middleware     middleware.Middleware
}

func NewAuthService(
	authRepository repository.AuthRepository,
	middleware middleware.Middleware,
) AuthService {
	return &authService{
		authRepository: authRepository,
		middleware:     middleware,
	}
}

func (as *authService) ConfirmEmail(uuid string) error {
	user, err := as.authRepository.GetUserTempData(uuid)
	if err != nil {
		return err
	}

	if err = as.authRepository.AddUser(user); err != nil {
		return err
	}

	err = as.authRepository.ClearTempData(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (as *authService) ForgotPassword(usernameOrEmail string) error {
	username, email, err := as.ResolveUserCredentials(usernameOrEmail)
	if err != nil {
		return err
	}

	token, err := as.middleware.GenerateResetPasswordJWT(email)
	if err != nil {
		return err
	}

	link := helpers.GenerateResetLink(token)

	html, err := as.GeneratePasswordResetEmailHTML(username, link)
	if err != nil {
		return err
	}

	return mailer.SendEmail(email, "Password reset", html, "text/html")
}

func (as *authService) GeneratePasswordResetEmailHTML(username string, link string) (string, error) {
	html, err := filereader.GetHTMLTempalate("forgot_password.html")
	if err != nil {
		return "", err
	}
	html = strings.Replace(html, "{url}", link, -1)
	html = strings.Replace(html, "{username}", username, -1)

	return html, nil
}

func (as *authService) ResolveUserCredentials(usernameOrEmail string) (string, string, error) {
	if strings.Contains(usernameOrEmail, "@") {
		return as.ValidateEmailAndFetchUsername(usernameOrEmail)
	} else {
		return as.ValidateUsernameAndFetchEmail(usernameOrEmail)
	}
}

func (as *authService) ValidateUsernameAndFetchEmail(username string) (string, string, error) {
	email, exists, err := as.authRepository.GetEmailByUsernameIfExists(username)
	if err != nil {
		return "", "", err
	}

	if !exists {
		return "", "", utils.ErrUsernameDoesNotExist
	}

	return username, email, nil
}

func (as *authService) ValidateEmailAndFetchUsername(email string) (string, string, error) {
	usernameFromDB, exists, err := as.authRepository.GetUsernameByEmailIfExists(email)
	if err != nil {
		return "", "", err
	}

	if !exists {
		return "", "", utils.ErrEmailDoesNotExist
	}

	return usernameFromDB, email, nil
}

func (as *authService) GoogleAuth(code string) (string, string, error) {
	tokenResponse, err := helpers.ExchangeAuthCodeForToken(code)
	if err != nil {
		return "", "", err
	}

	claims, err := helpers.ParseJWT(tokenResponse.IdToken)
	if err != nil {
		return "", "", err
	}

	username, err := as.HandleUserAccount(claims)
	if err != nil {
		return "", "", err
	}

	accessToken, err := as.middleware.GenerateSignInAccessToken(username)
	if err != nil {
		return "", "", err
	}
	fefreshToken, err := as.middleware.GenerateRefreshToken(username)
	if err != nil {
		return "", "", err
	}

	return accessToken, fefreshToken, nil
}

func (as *authService) HandleUserAccount(claims *body.GoogleJwtClaims) (string, error) {
	username, userExists, err := as.authRepository.GetUsernameByEmailIfExists(claims.Email)
	if err != nil {
		return "", err
	}

	if userExists {
		return username, nil
	}

	panic("unimplemented")

	// avatarID, err := fileRepo.AddFile(claims.PicURL, db)
	// if err != nil {
	// 	return "", err
	// }

	// err = repository.AddUserWithGoogleAuth(claims, avatarID, db)
	// if err != nil {
	// 	return "", err
	// }

	// return claims.Email, nil
}

func (as *authService) ResetPassword(
	resetPasswordBody body.ResetPassword,
	email string,
	token string,
) error {
	oldPassword, encryptedPrivateKey, err := as.authRepository.GetUserPasswordPrivateKeyByEmail(email)
	if err != nil {
		return err
	}

	decryptedPrivateKey, err := aes_encryption.DecryptWithKey(encryptedPrivateKey, oldPassword)
	if err != nil {
		return err
	}

	newEncryptedPassword, err := bcrypt_encryption.HashPassword(resetPasswordBody.Password)
	if err != nil {
		return err
	}

	newEncryptedPrivateKey, err := aes_encryption.EncryptWithKey(decryptedPrivateKey, newEncryptedPassword)
	if err != nil {
		return err
	}

	err = as.authRepository.ChangeUserPasswordPrivateKeyByEmail(newEncryptedPassword, newEncryptedPrivateKey, email)
	if err != nil {
		return err
	}

	return as.authRepository.AddJWTToBlacklist(token)
}

func (as *authService) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fiber.ErrUnauthorized
	}

	if int(claims["type"].(float64)) != utils.Refresh {
		return "", fiber.ErrUnauthorized
	}

	return as.middleware.GenerateSignInAccessToken(claims["username"].(string))
}

func (as *authService) SignIn(signInBody body.SignInBody) (string, string, error) {
	hashPassword, err := as.authRepository.GetUserPassword(signInBody.UsernameOrEmail)
	if err != nil {
		return "", "", fiber.ErrUnauthorized
	}

	match := bcrypt_encryption.CheckPasswordHash(signInBody.Password, hashPassword.Password)
	if !match {
		return "", "", fiber.ErrUnauthorized
	}

	accessToken, err := as.middleware.GenerateSignInAccessToken(hashPassword.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := as.middleware.GenerateRefreshToken(hashPassword.Username)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (as *authService) SignUp(user *body.SignUpBody) error {
	isUsernameAvailable, err := as.authRepository.CheckIfUsernameIsAvailable(user.Username)
	if err != nil || !isUsernameAvailable {
		return utils.ErrUsernameIsTaken
	}

	isEmailAvailable, err := as.authRepository.CheckIfEmailIsAvailable(user.Email)
	if err != nil || !isEmailAvailable {
		return err
	}

	hashedPassword, err := bcrypt_encryption.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	encryptedPrivateKey, err := aes_encryption.EncryptWithKey(user.PrivateKey, user.Password)
	if err != nil {
		return err
	}
	user.PrivateKey = encryptedPrivateKey

	userUUID := uuid.New().String()
	err = as.authRepository.AddTempUserData(user, userUUID)
	if err != nil {
		return err
	}

	token, err := as.middleware.GenerateConfirmPasswordToken(userUUID)
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

func (as *authService) SignOut(signOutBody body.SignOut) error {
	var tokens = []model.Token{{Token: signOutBody.AccessToken}, {Token: signOutBody.RefreshToken}}
	return as.authRepository.AddMultipleJWTToBlacklist(tokens)
}
