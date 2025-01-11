package repository

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
)

type FreelanceRepository interface {
	GetFreelanceServiceById(id int) (*model.FreelanceByID, error)
	GetFreelanceServiceByIdReviews(id int, cursorData string, lastID int64, maxReviews int) (*[]model.Review, error)
	GetResrictedFreelanceById(id int) (*model.FreelanceByIdRestricted, error)
}
