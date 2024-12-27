package service

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
)

type FreelanceService interface {
	GetFreelanceById(id int) (*dto.Freelance, error)
	GetReviewsByFreelanceID(id int, reviewsCursor string) (*dto.FreelanceReviews, error)
}
