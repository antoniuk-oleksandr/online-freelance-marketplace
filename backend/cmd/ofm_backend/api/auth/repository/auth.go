package repository

import (
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"

	"github.com/jmoiron/sqlx"
)

func ValidateUser(credentials *model.Credentials, db *sqlx.DB) (bool, error) {
	var valid bool

	query := `SELECT EXISTS(
    	SELECT * FROM users WHERE username = $1 AND password = $2
    )`

	var err error
	err = db.Get(&valid, query, credentials.Username, credentials.Password)
	if err != nil {
		return false, err
	}

	return valid, nil
}

func AddUser(user *body.SignUpBody, db *sqlx.DB) error {
	query := "INSERT INTO users" +
		" (first_name, surname, email, username, password) " +
		"VALUES (:first_name, :surname, :email, :username, :password)"

	_, err := db.NamedExec(query, user)
	return err
}
