package service

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
)

type AuthService interface {
	ConfirmEmail(uuid string) error
	ForgotPassword(usernameOrEmail string) error
	GeneratePasswordResetEmailHTML(username string, link string) (string, error)
	ResolveUserCredentials(usernameOrEmail string) (string, string, error)
	ValidateUsernameAndFetchEmail(username string) (string, string, error)
	ValidateEmailAndFetchUsername(email string) (string, string, error)
	HandleUserAccount(googleUserInfo *model.GoogleUserInfo, signUpBody *body.GoogleSignUpBody) (int64, error)
	ResetPassword(resetPasswordBody body.ResetPassword, email string, token string) error
	SignIn(signInBody body.SignInBody) (*model.SignResponse, *model.SignInData, error)
	SignUp(user *body.SignUpBody) error
	SignOut(signOutBody body.SignOut) error
	GetUserSessionData(userId int64) (*model.UserSessionData, error)
	CheckIfEmailIsAvailable(email string) (bool, error)
	SignInWithGoogle(signInBody *body.GoogleSignInBody) (*model.SignResponse, *model.SignInData, error)
	SignUpWithGoogle(signUpBody *body.GoogleSignUpBody) (*model.SignResponse, *model.SignInData, error)
	CreateSignInTokens(usernameOrEmail string) (*model.SignResponse, *model.SignInData, error)
	CreateSignInTokensWithData(googleUserInfo *model.GoogleUserInfo, userId int64, signUpBody *body.GoogleSignUpBody) (*model.SignResponse, *model.SignInData, error) 
}
