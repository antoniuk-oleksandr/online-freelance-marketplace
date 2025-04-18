package repository

import (
	"context"
	"database/sql"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/queries"
	"ofm_backend/cmd/ofm_backend/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type authRepisotory struct {
	posgresqlDb *sqlx.DB
	redisDb     *redis.Client
}

func NewAuthRepository(posgresqlDb *sqlx.DB, redisDb *redis.Client) AuthRepository {
	return &authRepisotory{
		posgresqlDb: posgresqlDb,
		redisDb:     redisDb,
	}
}

func (ar *authRepisotory) AddUser(user *body.SignUpBody) error {
	_, err := ar.posgresqlDb.NamedExec(queries.AddUserQuery, user)
	return err
}

func (ar *authRepisotory) GetUserTempData(uuid string) (*body.SignUpBody, error) {
	data, err := ar.redisDb.HGetAll(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, utils.ErrTempTokenExpired
	}

	return &body.SignUpBody{
		Email:      data["email"],
		FirstName:  data["firstName"],
		Password:   data["password"],
		Surname:    data["surname"],
		Username:   data["username"],
		PrivateKey: data["privateKey"],
		PublicKey:  data["publicKey"],
	}, nil
}

func (ar *authRepisotory) ClearTempData(uuid string) error {
	_, err := ar.redisDb.Del(context.Background(), uuid).Result()
	return err
}

func (ar *authRepisotory) GetEmailByUsernameIfExists(username string) (string, bool, error) {
	var email string
	err := ar.posgresqlDb.Get(&email, queries.GetEmailByUsernameIfExistsQuery, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, utils.ErrUsernameDoesNotExist
		}
		return "", false, err
	}

	return email, true, nil
}

func (ar *authRepisotory) CheckIfUsernameIsAvailable(username string) (bool, error) {
	var available bool
	err := ar.posgresqlDb.Get(&available, queries.CheckIfUsernameIsAvailableQuery, username)
	if err != nil {
		return false, err
	}

	return !available, nil
}

func (ar *authRepisotory) CheckIfEmailIsAvailable(email string) (bool, error) {
	var available bool
	err := ar.posgresqlDb.Get(&available, queries.CheckIfEmailIsAvailableQuery, email)
	if err != nil {
		return false, err
	}

	return !available, nil
}

func (ar *authRepisotory) AddTempUserData(user *body.SignUpBody, userUUID string) error {
	pipe := ar.redisDb.Pipeline()

	pipe.HSet(context.Background(), userUUID, map[string]interface{}{
		"email":      user.Email,
		"firstName":  user.FirstName,
		"password":   user.Password,
		"surname":    user.Surname,
		"username":   user.Username,
		"privateKey": user.PrivateKey,
		"publicKey":  user.PublicKey,
	})

	pipe.Expire(context.Background(), user.Username, 15*time.Minute)

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepisotory) GetUserPassword(usernameOrEmail string) (*model.UsernamePassword, error) {
	var usernamePassword model.UsernamePassword

	var err error
	err = ar.posgresqlDb.Get(&usernamePassword, queries.GetUserPasswordQuery, usernameOrEmail)
	if err != nil {
		return nil, err
	}

	return &usernamePassword, nil
}

func (ar *authRepisotory) ChangeUserPasswordPrivateKeyByEmail(
	encryptedPassword string, encryptedPrivateKey string, email string,
) error {
	_, err := ar.posgresqlDb.Exec(
		queries.ChangeUserPasswordPrivateKeyByEmailQuery, encryptedPassword, encryptedPrivateKey, email,
	)
	
	return err
}

func (ar *authRepisotory) GetUsernameByEmailIfExists(email string) (string, bool, error) {
	var username string

	err := ar.posgresqlDb.Get(&username, queries.GetUsernameByEmailIfExistsQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		}
		return "", false, err
	}
	return username, true, nil
}

func (ar *authRepisotory) AddUserWithGoogleAuth(claims *body.GoogleJwtClaims, avatarID int) error {
	_, err := ar.posgresqlDb.Exec(
		queries.AddUserWithGoogleAuthQuery, claims.Email, claims.Email, "",
		claims.GivenName, claims.FamilyName, avatarID,
	)
	if err != nil {
		return err
	}

	return err
}

func (ar *authRepisotory) AddJWTToBlacklist(token string) error {
	_, err := ar.posgresqlDb.Exec(queries.AddJWTToBlackListQuery, token)
	return err
}

func (ar *authRepisotory) AddMultipleJWTToBlacklist(tokens []model.Token) error {
	_, err := ar.posgresqlDb.NamedExec(queries.AddMultipleJWTToBlackListQuery, tokens)
	return err
}

func (ar *authRepisotory) GetUserPasswordPrivateKeyByEmail(email string) (string, string, error) {
	var password string
	var privateKey string

	row := ar.posgresqlDb.QueryRowx(queries.GetUserPasswordPrivateKeyByEmailQuery, email)
	if err := row.Scan(&password, &privateKey); err != nil {
		return "", "", err
	}

	return password, privateKey, nil
}
