package service

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
)

type FreelanceService interface {
	GetFreelanceById(id int) (*dto.FreelanceByIDResponse, error)
	GetReviewsByFreelanceID(id int, reviewsCursor string) (*dto.FreelanceReviewsResponse, error)
	GetResrictedFreelanceById(id int) (*dto.FreelanceByIdRestricted, error)
}
