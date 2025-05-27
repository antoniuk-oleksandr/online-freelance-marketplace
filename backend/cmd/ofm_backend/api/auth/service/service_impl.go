package service

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/helpers"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/repository"
	fileService "ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/aes_encryption"
	"ofm_backend/cmd/ofm_backend/utils/bcrypt_encryption"
	filereader "ofm_backend/internal/file_reader"
	"ofm_backend/internal/mailer"
	"ofm_backend/internal/middleware"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type authService struct {
	authRepository repository.AuthRepository
	middleware     middleware.Middleware
	fileService    fileService.FileService
}

func NewAuthService(
	authRepository repository.AuthRepository,
	middleware middleware.Middleware,
	fileService fileService.FileService,
) AuthService {
	return &authService{
		authRepository: authRepository,
		middleware:     middleware,
		fileService:    fileService,
	}
}

func (as *authService) ConfirmEmail(uuid string) error {
	user, err := as.authRepository.GetUserTempData(uuid)
	if err != nil {
		return err
	}

	user.MasterKey, err = aes_encryption.GenerateMasterKey()
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

func (as *authService) HandleUserAccount(
	googleUserInfo *model.GoogleUserInfo, signUpBody *body.GoogleSignUpBody,
) (int64, error) {
	tx, err := as.authRepository.CreateTransaction()
	if err != nil {
		return -1, err
	}

	// avatarId := 1
	avatarId, fileName, err := as.fileService.UploadFromURLWithoutTransaction(googleUserInfo.Picture)
	if err != nil {
		as.authRepository.RollbackTransaction(tx)
		return -1, err
	}
	googleUserInfo.Picture = *utils.AddServerURLToFiles(&fileName)
	// googleUserInfo.Picture = "image.png"

	signUpBody.MasterKey, err = aes_encryption.GenerateMasterKey()
	if err != nil {
		return -1, err
	}

	userId, err := as.authRepository.AddUserWithGoogleAuth(googleUserInfo, avatarId, signUpBody)
	if err != nil {
		as.authRepository.RollbackTransaction(tx)
		return -1, err
	}
	return userId, as.authRepository.CommitTransaction(tx)
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

	initialVector, err := aes_encryption.GenerateAESIV()
	if err != nil {
		return err
	}

	newEncryptedPrivateKey, err := aes_encryption.EncryptWithKey(decryptedPrivateKey, newEncryptedPassword, initialVector)
	if err != nil {
		return err
	}

	err = as.authRepository.ChangeUserPasswordPrivateKeyByEmail(newEncryptedPassword, newEncryptedPrivateKey, email)
	if err != nil {
		return err
	}

	return as.authRepository.AddJWTToBlacklist(token)
}

func (as *authService) SignIn(signInBody body.SignInBody) (*model.SignResponse, *model.SignInData, error) {
	hashPassword, err := as.authRepository.GetUserPassword(signInBody.UsernameOrEmail)
	if err != nil {
		return nil, nil, fiber.ErrUnauthorized
	}

	match := bcrypt_encryption.CheckPasswordHash(signInBody.Password, hashPassword.Password)
	if !match {
		return nil, nil, fiber.ErrUnauthorized
	}

	return as.CreateSignInTokens(signInBody.UsernameOrEmail)
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
	link := fmt.Sprintf("%s/confirm-email?token=%s", host, token)

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

func (as *authService) GetUserSessionData(userId int64) (*model.UserSessionData, error) {
	userSessionData, err := as.authRepository.GetUserSessionData(userId)
	if err != nil {
		return nil, err
	}

	userSessionData.Authenticated = true
	return userSessionData, nil
}

func (as *authService) CheckIfEmailIsAvailable(email string) (bool, error) {
	return as.authRepository.CheckIfEmailIsAvailable(email)
}

func (as *authService) SignInWithGoogle(
	signInBody *body.GoogleSignInBody,
) (*model.SignResponse, *model.SignInData, error) {
	googleUserInfo, err := helpers.GetGoogleUserInfo(signInBody.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	return as.CreateSignInTokens(googleUserInfo.Email)
}

func (as *authService) SignUpWithGoogle(
	signUpBody *body.GoogleSignUpBody,
) (*model.SignResponse, *model.SignInData, error) {
	googleUserInfo, err := helpers.GetGoogleUserInfo(signUpBody.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	userId, err := as.HandleUserAccount(googleUserInfo, signUpBody)
	if err != nil {
		return nil, nil, err
	}

	return as.CreateSignInTokensWithData(googleUserInfo, userId, signUpBody)
}

func (as *authService) CreateSignInTokens(usernameOrEmail string) (*model.SignResponse, *model.SignInData, error) {
	signInData, err := as.authRepository.GetUserSignInData(usernameOrEmail)
	if err != nil {
		return nil, nil, err
	}
	signInData.Avatar = *utils.AddServerURLToFiles(&signInData.Avatar)

	accessTokenExpiration, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if err != nil || accessTokenExpiration == 0 {
		return nil, nil, utils.ErrUnexpectedError
	}

	accessToken, err := as.middleware.GenerateSignInAccessToken(signInData.Username, accessTokenExpiration)
	if err != nil {
		return nil, nil, err
	}

	refreshTokenExpiration, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if err != nil || accessTokenExpiration == 0 {
		return nil, nil, utils.ErrUnexpectedError
	}

	refreshToken, err := as.middleware.GenerateRefreshToken(signInData.Username, refreshTokenExpiration)
	if err != nil {
		return nil, nil, err
	}

	signResponse := &model.SignResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessTokenExpiration,
		RefreshTokenExpiresAt: refreshTokenExpiration,
	}

	return signResponse, signInData, nil
}

func (as *authService) CreateSignInTokensWithData(
	googleUserInfo *model.GoogleUserInfo, userId int64, signUpBody *body.GoogleSignUpBody,
) (*model.SignResponse, *model.SignInData, error) {
	accessTokenExpiration, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if err != nil || accessTokenExpiration == 0 {
		return nil, nil, utils.ErrUnexpectedError
	}

	accessToken, err := as.middleware.GenerateSignInAccessTokenWithData(
		googleUserInfo.Email, googleUserInfo.Picture,
		userId, accessTokenExpiration,
	)
	if err != nil {
		return nil, nil, err
	}

	refreshTokenExpiration, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if err != nil || accessTokenExpiration == 0 {
		return nil, nil, utils.ErrUnexpectedError
	}

	refreshToken, err := as.middleware.GenerateRefreshToken(googleUserInfo.Email, refreshTokenExpiration)
	if err != nil {
		return nil, nil, err
	}

	signResponse := &model.SignResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessTokenExpiration,
		RefreshTokenExpiresAt: refreshTokenExpiration,
	}

	return signResponse, &model.SignInData{
		Id:             strconv.FormatInt(userId, 10),
		Username:       googleUserInfo.Email,
		Avatar:         googleUserInfo.Picture,
		PrivateKey:     signUpBody.PrivateKey,
		PrivateKeyIV:   signUpBody.PrivateKeyIV,
		PrivateKeySalt: signUpBody.PrivateKeySalt,
		MasterKey:      signUpBody.MasterKey,
	}, nil
}
