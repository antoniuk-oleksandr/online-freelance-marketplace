package repository

import "ofm_backend/internal/middleware/model"

type MiddlewareRepository interface {
	CheckIfTokenBacklisted(token string) (bool, error)
	GetUserSignInDataByUsername(username string) (*model.UserSignInData, error)
}
