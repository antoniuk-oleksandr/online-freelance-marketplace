package service

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
	"ofm_backend/cmd/ofm_backend/api/freelance/mapper"
	"ofm_backend/cmd/ofm_backend/api/freelance/repository"
	"ofm_backend/cmd/ofm_backend/api/freelance/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
)

type freelanceService struct {
	repository repository.FreelanceRepository
}

func NewFreelanceService(repo repository.FreelanceRepository) FreelanceService {
	return &freelanceService{
		repository: repo,
	}
}

func (fs *freelanceService) GetFreelanceById(id int) (*dto.Freelance, error) {
	maxReviews := utils.GetMaxReviewsValue()

	freelanceService, err := fs.repository.GetFreelanceServiceById(id)
	if err != nil {
		return nil, main_utils.ErrNotFound
	}

	freelanceServiceReviews, err := fs.repository.GetFreelanceServiceByIdReviews(id, "", maxReviews+1)
	if err != nil {
		return nil, err
	}

	freelanceDTO := mapper.MapFreelanceModelToDTO(freelanceService, freelanceServiceReviews)
	freelanceDTOWithFileLinks := main_utils.AddServerURLToFiles(&freelanceDTO)

	return freelanceDTOWithFileLinks, nil
}

func (fs *freelanceService) GetReviewsByFreelanceID(id int, reviewsCursor string) (*dto.FreelanceReviews, error) {
	cursorData, err := utils.GetDataFromReviewsCursor(reviewsCursor)
	if err != nil {
		return nil, err
	}
	maxReviews := utils.GetMaxReviewsValue()

	reviews, err := fs.repository.GetFreelanceServiceByIdReviews(id, cursorData, maxReviews+1)
	if err != nil {
		return nil, err
	}

	var hasMoreReviews bool
	var newReviewsCursor *string
	if len(*reviews) == maxReviews+1 {
		hasMoreReviews = true
		lastReview := (*reviews)[maxReviews]
		newReviewsCursor = utils.BuildReviewsCursor(lastReview.EndedAt)
		reviewsWithoutLast := (*reviews)[:maxReviews]
		reviews = &reviewsWithoutLast
	}

	return &dto.FreelanceReviews{
		Reviews:        reviews,
		HasMoreReviews: hasMoreReviews,
		ReviewsCursor:  newReviewsCursor,
	}, nil
}
