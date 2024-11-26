package repository

import (
	"context"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func AddUser(user *body.SignUpBody, db *sqlx.DB) error {
	query := "INSERT INTO users" +
		" (first_name, surname, email, username, password) " +
		"VALUES (:first_name, :surname, :email, :username, :password)"

	_, err := db.NamedExec(query, user)
	return err
}

func GetUserTempData(uuid string, redisDB *redis.Client) (*body.SignUpBody, error) {
	data, err := redisDB.HGetAll(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, utils.ErrTempTokenExpired
	}

	return &body.SignUpBody{
		Email:     data["email"],
		FirstName: data["firstName"],
		Password:  data["password"],
		Surname:   data["surname"],
		Username:  data["username"],
	}, nil
}

func ClearTempData(uuid string, redisDB *redis.Client) error {
	_, err := redisDB.Del(context.Background(), uuid).Result()
	return err
}