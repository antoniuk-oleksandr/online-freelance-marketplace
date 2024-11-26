package repository

import (
	"database/sql"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/jmoiron/sqlx"
)

func GetEmailByUsernameIfExists(
	username string, db *sqlx.DB,
) (string, bool, error) {
	var email string

	query := `SELECT email FROM users WHERE username = $1`

	err := db.Get(&email, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, utils.ErrUsernameDoesNotExist
		}
		return "", false, err
	}

	return email, true, nil
}
