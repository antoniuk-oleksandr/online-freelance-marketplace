package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/dto"
	"ofm_backend/cmd/ofm_backend/api/home_data/model"
)

func MapKeyMetricsModelToDto(
	KeyMetricsModel *model.KeyMetrics,
) dto.KeyMetrics {
	return dto.KeyMetrics{
		FreelancesAvailable: KeyMetricsModel.FreelancesAvailable,
		ProjectsCompleted:   KeyMetricsModel.ProjectsCompleted,
		AvgRating:           KeyMetricsModel.AvgRating,
	}
}

func MapBestFreelanceModelsToDto(
	bestFeelanceModels []model.BestFreelance,
) []dto.BestFreelance {
	var bestFreelanceDtos []dto.BestFreelance = make([]dto.BestFreelance, len(bestFeelanceModels))

	for index, item := range bestFeelanceModels {
		bestFreelanceDtos[index] = dto.BestFreelance{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			Image:       item.Image,
		}
	}

	return bestFreelanceDtos
}

func MapBestFreelancerModelsToDtos(
	bestFreelancerModels []model.BestFreelancer,
) []dto.BestFreelancer {
	var bestFreelancerDtos []dto.BestFreelancer = make([]dto.BestFreelancer, len(bestFreelancerModels))

	for index, item := range bestFreelancerModels {
		bestFreelancerDtos[index] = dto.BestFreelancer{
			Id:                item.Id,
			FirstName:         item.FirstName,
			Surname:           item.Surname,
			Rating:            item.Rating,
			CompletedProjects: item.CompletedProjects,
			Avatar:            item.Avatar,
		}
	}

	return bestFreelancerDtos
}

func MapBestReviewModelsToDtos(
	bestReviewsModels []model.BestReview,
) []dto.BestReview {
	var bestReviewDtos []dto.BestReview = make([]dto.BestReview, len(bestReviewsModels))

	for index, item := range bestReviewsModels {
		bestReviewDtos[index] = dto.BestReview{
			FirstName: item.FirstName,
			Surname:   item.Surname,
			Content:   item.Content,
			Rating:    item.Rating,
		}
	}

	return bestReviewDtos
}

func MapHomeDataModelToDto(
	keyMetricsDto dto.KeyMetrics,
	bestFreelancerDtos []dto.BestFreelancer,
	bestFreelanceDtos []dto.BestFreelance,
	bestReviewDtos []dto.BestReview,
) dto.HomeData {
	return dto.HomeData{
		KeyMetrics:      keyMetricsDto,
		BestFreelancers: bestFreelancerDtos,
		BestReviews:     bestReviewDtos,
		BestFreelances:  bestFreelanceDtos,
	}
}
