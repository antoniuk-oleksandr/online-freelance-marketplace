package service

import "ofm_backend/cmd/ofm_backend/api/user/dto"

type UserService interface {
	GetUserById(id int) (*dto.UserByIDResponse, error)
	GetReviewsByUserId(id int, cursor string) (*dto.ReviewsResponse, error)
	GetServicesByUserId(id int, cursor string) (*dto.ServicesResponse, error)
}
