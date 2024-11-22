package repository

import (
	"ofm_backend/cmd/ofm_backend/api/auth/model"

	"github.com/jmoiron/sqlx"
)

func GetUserPassword(usernameOrEmail string, db *sqlx.DB) (*model.UsernamePassword, error) {
	var usernamePassword model.UsernamePassword

	query := `SELECT username, password FROM users WHERE username = $1 OR email = $1`

	var err error
	err = db.Get(&usernamePassword, query, usernameOrEmail)
	if err != nil {
		return nil, err
	}

	return &usernamePassword, nil
}