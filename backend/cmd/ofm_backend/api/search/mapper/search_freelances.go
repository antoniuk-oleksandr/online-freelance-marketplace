package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/search/dto"
	"ofm_backend/cmd/ofm_backend/api/search/model"
)

func MapSearchFreelancesModelToDTO(searchServiceModels []model.SearchService) []*dto.SearchFreelance {
	var freelances []*dto.SearchFreelance

	for _, item := range searchServiceModels {
		freelances = append(freelances, &dto.SearchFreelance{
			ID:           item.ID,
			CreatedAt:    item.CreatedAt,
			Description:  item.Description,
			Title:        item.Title,
			CategoryId:   item.CategoryId,
			FreelancerId: item.FreelancerId,
			Image:        item.Image,
			ReviewsCount: item.ReviewsCount,
			Rating:       item.Rating,
			MinPrice:     item.MinPrice,
		})
	}

	return freelances
}
