package repository

import (
	"context"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/dto"
	"ofm_backend/cmd/ofm_backend/api/auth/model"

	"github.com/jmoiron/sqlx"
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
	AddUserWithGoogleAuth(googleUserInfo *model.GoogleUserInfo, avatarID int, signUpBody *body.GoogleSignUpBody) (int64, error)
	AddJWTToBlacklist(token string) error
	AddMultipleJWTToBlacklist(tokens []model.Token) error
	GetUserPasswordPrivateKeyByEmail(email string) (string, string, error)
	GetUserSignInData(usernameOrEmail string) (*model.SignInData, error)
	GetUserSessionDataFromPostgres(userId int64) (*model.UserSessionData, error)
	GetUserSessionDataCache(userId int64, ctx context.Context) (*dto.UserSessionData, error)
	AddUserSessionDataCache(userId int64, userSessionData *dto.UserSessionData, ctx context.Context) error
	CreateTransaction() (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}
