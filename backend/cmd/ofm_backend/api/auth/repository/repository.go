package repository

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
)

type AuthRepository interface {
	AddUser(user *body.SignUpBody) error
	GetUserTempData(uuid string) (*body.SignUpBody, error)
	ClearTempData(uuid string) error
	GetEmailByUsernameIfExists(username string) (string, bool, error)
	CheckIfUsernameIsAvailable(username string) (bool, error)
	CheckIfEmailIsAvailable(email string) (bool, error)
	AddTempUserData(user *body.SignUpBody, userUUID string) error
	GetUserPassword(usernameOrEmail string) (*model.UsernamePassword, error)
	ChangeUserPasswordPrivateKeyByEmail(encryptedPassword string, encryptedPrivateKey string, email string) error
	GetUsernameByEmailIfExists(email string) (string, bool, error)
	AddUserWithGoogleAuth(claims *body.GoogleJwtClaims, avatarID int) error
	AddJWTToBlacklist(token string) error
	AddMultipleJWTToBlacklist(tokens []model.Token) error
	GetUserPasswordPrivateKeyByEmail(email string) (string, string, error)
}
