package repository

import "github.com/jmoiron/sqlx"

func GetUserPassword(username string, db *sqlx.DB) (*string, error) {
	var password string

	query := `SELECT password FROM users WHERE username = $1`

	var err error
	err = db.Get(&password, query, username)
	if err != nil {
		return nil, err
	}

	return &password, nil
}