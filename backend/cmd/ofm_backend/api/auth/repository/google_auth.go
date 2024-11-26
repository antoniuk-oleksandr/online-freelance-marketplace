package repository

import (
	"database/sql"
	"ofm_backend/cmd/ofm_backend/api/auth/body"

	"github.com/jmoiron/sqlx"
)

func GetUsernameByEmailIfExists(
	email string, db *sqlx.DB,
) (string, bool, error) {
	var username string

	query := `SELECT username FROM users WHERE email = $1`

	err := db.Get(&username, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		}
		return "", false, err
	}
	return username, true, nil
}

func AddUserWithGoogleAuth(
	claims *body.GoogleJwtClaims, avatarID int, db *sqlx.DB,
) error {
	query := `
	INSERT INTO users
		(username, email, password, first_name, surname, avatar_id)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(
		query, claims.Email, claims.Email, "",
		claims.GivenName, claims.FamilyName, avatarID,
	)
	if err != nil {
		return err
	}

	return err
}
