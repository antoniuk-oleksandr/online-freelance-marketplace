package repository

import (
	"ofm_backend/internal/middleware/model"
	"ofm_backend/internal/middleware/queries"

	"github.com/jmoiron/sqlx"
)

type middlewareRepository struct {
	db *sqlx.DB
}

func NewMiddlewareRepository(db *sqlx.DB) MiddlewareRepository {
	return &middlewareRepository{
		db: db,
	}
}

func (mr *middlewareRepository) CheckIfTokenBacklisted(token string) (bool, error) {
	var exists bool

	err := mr.db.Get(&exists, queries.CheckIfTokenBlacklistedQuery, token)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (mr *middlewareRepository) GetUserSignInDataByUsername(
	username string,
) (*model.UserSignInData, error) {
	var userSignInData model.UserSignInData
	err := mr.db.Get(&userSignInData, queries.GetUserSignInDataByUsernameQuery, username)
	if err != nil {
		return nil, err
	}

	return &userSignInData, nil
}
