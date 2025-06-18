package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/dto"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"ofm_backend/cmd/ofm_backend/api/auth/queries"
	"ofm_backend/cmd/ofm_backend/utils"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	posgresqlDb *sqlx.DB
	redisDb     *redis.Client
}

func NewAuthRepository(posgresqlDb *sqlx.DB, redisDb *redis.Client) AuthRepository {
	return &authRepository{
		posgresqlDb: posgresqlDb,
		redisDb:     redisDb,
	}
}

func (ar *authRepository) AddUser(user *body.SignUpBody) error {
	_, err := ar.posgresqlDb.NamedExec(queries.AddUserQuery, user)
	return err
}

func (ar *authRepository) GetUserTempData(uuid string) (*body.SignUpBody, error) {
	data, err := ar.redisDb.HGetAll(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, utils.ErrTempTokenExpired
	}

	return &body.SignUpBody{
		Email:          data["email"],
		FirstName:      data["firstName"],
		Password:       data["password"],
		Surname:        data["surname"],
		Username:       data["username"],
		PrivateKey:     []byte(data["privateKey"]),
		PrivateKeyIV:   []byte(data["privateKeyIV"]),
		PrivateKeySalt: []byte(data["privateKeySalt"]),
		PublicKey:      []byte(data["publicKey"]),
	}, nil
}

func (ar *authRepository) ClearTempData(uuid string) error {
	_, err := ar.redisDb.Del(context.Background(), uuid).Result()
	return err
}

func (ar *authRepository) GetEmailByUsernameIfExists(username string) (string, bool, error) {
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

func (ar *authRepository) CheckIfUsernameIsAvailable(username string) (bool, error) {
	var available bool
	err := ar.posgresqlDb.Get(&available, queries.CheckIfUsernameIsAvailableQuery, username)
	if err != nil {
		return false, err
	}

	return !available, nil
}

func (ar *authRepository) CheckIfEmailIsAvailable(email string) (bool, error) {
	var available bool
	err := ar.posgresqlDb.Get(&available, queries.CheckIfEmailIsAvailableQuery, email)
	if err != nil {
		return false, err
	}

	return !available, nil
}

func (ar *authRepository) AddTempUserData(user *body.SignUpBody, userUUID string) error {
	pipe := ar.redisDb.Pipeline()

	pipe.HSet(context.Background(), userUUID, map[string]any{
		"email":          user.Email,
		"firstName":      user.FirstName,
		"password":       user.Password,
		"surname":        user.Surname,
		"username":       user.Username,
		"privateKey":     user.PrivateKey,
		"publicKey":      user.PublicKey,
		"privateKeyIV":   user.PrivateKeyIV,
		"privateKeySalt": user.PrivateKeySalt,
	})

	pipe.Expire(context.Background(), user.Username, 15*time.Minute)

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) GetUserPassword(usernameOrEmail string) (*model.UsernamePassword, error) {
	var usernamePassword model.UsernamePassword

	var err error
	err = ar.posgresqlDb.Get(&usernamePassword, queries.GetUserPasswordQuery, usernameOrEmail)
	if err != nil {
		return nil, err
	}

	return &usernamePassword, nil
}

func (ar *authRepository) ChangeUserPasswordPrivateKeyByEmail(
	encryptedPassword string, encryptedPrivateKey string, email string,
) error {
	_, err := ar.posgresqlDb.Exec(
		queries.ChangeUserPasswordPrivateKeyByEmailQuery, encryptedPassword, encryptedPrivateKey, email,
	)

	return err
}

func (ar *authRepository) GetUsernameByEmailIfExists(email string) (string, bool, error) {
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

func (ar *authRepository) AddUserWithGoogleAuth(
	googleUserInfo *model.GoogleUserInfo, avatarID int, signUpBody *body.GoogleSignUpBody,
) (int64, error) {
	var userId int64
	params := map[string]any{
		"email":            googleUserInfo.Email,
		"username":         googleUserInfo.Email,
		"password":         "",
		"first_name":       googleUserInfo.GivenName,
		"surname":          googleUserInfo.FamilyName,
		"avatar_id":        avatarID,
		"private_key":      signUpBody.PrivateKey,
		"public_key":       signUpBody.PublicKey,
		"private_key_iv":   signUpBody.PrivateKeyIV,
		"private_key_salt": signUpBody.PrivateKeySalt,
		"master_key":       signUpBody.MasterKey,
	}

	rows, err := ar.posgresqlDb.NamedQuery(queries.AddUserWithGoogleAuthQuery, params)
	if err != nil {
		return userId, err
	}

	if rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			return userId, err
		}
	}

	return userId, err
}

func (ar *authRepository) AddJWTToBlacklist(token string) error {
	_, err := ar.posgresqlDb.Exec(queries.AddJWTToBlackListQuery, token)
	return err
}

func (ar *authRepository) AddMultipleJWTToBlacklist(tokens []model.Token) error {
	_, err := ar.posgresqlDb.NamedExec(queries.AddMultipleJWTToBlackListQuery, tokens)
	return err
}

func (ar *authRepository) GetUserPasswordPrivateKeyByEmail(email string) (string, string, error) {
	var password string
	var privateKey string

	row := ar.posgresqlDb.QueryRowx(queries.GetUserPasswordPrivateKeyByEmailQuery, email)
	if err := row.Scan(&password, &privateKey); err != nil {
		return "", "", err
	}

	return password, privateKey, nil
}

func (ar *authRepository) GetUserSignInData(usernameOrEmail string) (*model.SignInData, error) {
	var signInData model.SignInData

	var userDataJson []byte
	var chatPartnersJson []byte
	if err := ar.posgresqlDb.QueryRowx(queries.GetUserSignInDataQuery, usernameOrEmail).
		Scan(&userDataJson, &chatPartnersJson); err != nil{
		return nil, err
	}

	if err := json.Unmarshal(userDataJson, &signInData.UserData); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(chatPartnersJson, &signInData.ChatPartners); err != nil {
		return nil, err
	}

	return &signInData, nil
}

func (ar *authRepository) GetUserSessionDataFromPostgres(userId int64) (*model.UserSessionData, error) {
	var userSessionData model.UserSessionData
	var jsonData []byte
	if err := ar.posgresqlDb.QueryRowx(queries.GetUserSessionDataQuery, userId).
		Scan(&jsonData, &userSessionData.MasterKey); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &userSessionData.ChatPartners); err != nil {
		return nil, err
	}

	return &userSessionData, nil
}

func (ar *authRepository) GetUserSessionDataCache(
    userId int64, ctx context.Context,
) (*dto.UserSessionData, error) {
    key := fmt.Sprintf("user:session:%d", userId)

    // Add debug logging
    log.Printf("Attempting to fetch session data from Redis for key: %s", key)

    result, err := ar.redisDb.Get(ctx, key).Result()
    if err != nil {
        if errors.Is(err, redis.Nil) {  // More robust nil check
            log.Printf("Cache miss - key not found: %s", key)
            return nil, utils.ErrSessionCacheNotFound
        }

        // Enhanced error logging
        log.Printf("Redis GET operation failed for key %s: %v (Type: %T)", key, err, err)

        // Check for specific error types
        if errors.Is(err, context.DeadlineExceeded) {
            log.Printf("Redis timeout occurred")
            return nil, fmt.Errorf("redis timeout: %w", err)
        }
        if strings.Contains(err.Error(), "connection refused") {
            log.Printf("Redis connection refused")
            return nil, fmt.Errorf("redis connection failed: %w", err)
        }

        return nil, fmt.Errorf("%w: %v", utils.ErrInvalidSessionCache, err)
    }

    // Debug successful fetch
    log.Printf("Successfully fetched %d bytes from Redis for key: %s", len(result), key)

    var userSessionData dto.UserSessionData
    if err := json.Unmarshal([]byte(result), &userSessionData); err != nil {
        log.Printf("Failed to unmarshal Redis data for key %s: %v", key, err)
        return nil, fmt.Errorf("json unmarshal failed: %w", err)
    }

    return &userSessionData, nil
}

func (ar *authRepository) AddUserSessionDataCache(
	userId int64, userSessionData *dto.UserSessionData, ctx context.Context,
) error {
	key := fmt.Sprintf("user:session:%d", userId)

	jsonData, err := json.Marshal(userSessionData)
	if err != nil {
		return err
	}

	return ar.redisDb.Set(ctx, key, jsonData, time.Hour*24).Err()
}

func (ar *authRepository) CreateTransaction() (*sqlx.Tx, error) {
	return ar.posgresqlDb.Beginx()
}

func (ar *authRepository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (ar *authRepository) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}
