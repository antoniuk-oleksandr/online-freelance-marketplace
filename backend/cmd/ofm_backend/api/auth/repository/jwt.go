package repository

import "github.com/jmoiron/sqlx"

func AddJWTToBlacklist(token string, db *sqlx.DB) error {
	query := `INSERT INTO blacklisted_tokens (token) VALUES ($1)`

	_, err := db.Exec(query, token)
	return err
}

func CheckIfTokenBacklisted(token string, db *sqlx.DB) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (SELECT token FROM blacklisted_tokens WHERE token = $1)`

	err := db.Get(&exists, query, token)
	if err != nil {
		return false, err
	}
	return exists, nil
}
