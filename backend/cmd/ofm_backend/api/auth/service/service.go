package service

import "ofm_backend/cmd/ofm_backend/api/auth/body"

type AuthService interface {
	ConfirmEmail(uuid string) error
	ForgotPassword(usernameOrEmail string) error
	GeneratePasswordResetEmailHTML(username string, link string) (string, error)
	ResolveUserCredentials(usernameOrEmail string) (string, string, error)
	ValidateUsernameAndFetchEmail(username string) (string, string, error)
	ValidateEmailAndFetchUsername(email string) (string, string, error)
	GoogleAuth(code string) (string, string, error)
	HandleUserAccount(claims *body.GoogleJwtClaims) (string, error)
	ResetPassword(resetPasswordBody body.ResetPassword, email string, token string) error
	RefreshToken(tokenString string) (string, error)
	SignIn(signInBody body.SignInBody) (string, string, error)
	SignUp(user *body.SignUpBody) error
	SignOut(signOutBody body.SignOut) error
}
