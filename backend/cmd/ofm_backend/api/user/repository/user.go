package repository

import "ofm_backend/cmd/ofm_backend/api/user/model"

type UserRepository interface {
	GetUserById(id int) (*model.User, error)
	GetReviewsByUserId(id int, endedAt *string, lastId int64, maxReviews int) (*[]model.UserByIdReview, error)
	GetServicesByUserId(id int, reviewsCount, lastId int64, maxServices int) (*[]model.UserByIdFreelanceService, error)
	
}
