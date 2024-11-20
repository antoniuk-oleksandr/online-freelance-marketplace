package repository

import (
	"context"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func CheckIfUsernameIsAvailable(username string, db *sqlx.DB) (bool, error) {
	var available bool
	
	query := `SELECT EXISTS (SELECT  * FROM users WHERE username = $1)`
	
	err := db.Get(&available, query, username)
	if err != nil {
		return false, err
	}
	
	return !available, nil
} 

func AddTempUserData(user *body.SignUpBody, redisDB *redis.Client) error {
	pipe := redisDB.Pipeline()
	
	pipe.HSet(context.Background(), user.Username, map[string]interface{}{
		"email":     user.Email,
		"firstName": user.FirstName,
		"password":  user.Password,
		"surname":   user.Surname,
	})
	
	pipe.Expire(context.Background(), user.Username, 15*time.Minute)
	
	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}
	
	return nil
}