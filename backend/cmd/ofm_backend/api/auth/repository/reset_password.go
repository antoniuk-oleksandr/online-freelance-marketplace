package repository

import "github.com/jmoiron/sqlx"

func ChangeUserPasswordByEmail(
	encodedPassword string, email string, db *sqlx.DB,
) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`

	_, err := db.Exec(query, encodedPassword, email)
	return err
}
