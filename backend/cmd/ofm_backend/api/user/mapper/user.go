package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/api/user/model"
)

func MapUserByIdReviewModelsToReviewUserDTOs(reviews *[]model.UserByIdReview) *[]dto.ReviewUserDTO {
	var reviewUsers []dto.ReviewUserDTO

	for _, review := range *reviews {
		user := dto.ReviewUserDTO{
			ID:        review.UserID,
			FirstName: review.FirstName,
			Surname:   review.Surname,
			Avatar:    review.Avatar,
		}

		reviewUsers = append(reviewUsers, user)
	}

	return &reviewUsers
}

func MapUserByIdReviewModelToReviewServiceDTOs(reviews *[]model.UserByIdReview) *[]dto.ReviewService {
	var reviewServices []dto.ReviewService

	for _, review := range *reviews {
		service := dto.ReviewService{
			ID:    review.ServiceID,
			Price: review.Price,
			Image: review.ServiceImage,
			Title: review.Title,
		}

		reviewServices = append(reviewServices, service)
	}

	return &reviewServices
}

func MapReviewUsersServicesToUserByIdReviewDTOs(
	reviews *[]model.UserByIdReview,
	users *[]dto.ReviewUserDTO,
	services *[]dto.ReviewService,
) *[]dto.UserByIdReviewDto {
	var reviewDTOs []dto.UserByIdReviewDto

	for i, review := range *reviews {
		reviewDTO := dto.UserByIdReviewDto{
			ID:            review.ID,
			Rating:        review.Rating,
			Content:       review.Content,
			CreatedAt:     review.CreatedAt,
			EndedAt:       review.EndedAt,
			Customer:      (*users)[i],
			ReviewService: (*services)[i],
		}

		reviewDTOs = append(reviewDTOs, reviewDTO)
	}

	return &reviewDTOs
}

func MapUserByIdModelToDTO(
	model *model.User,
	reviews *[]dto.UserByIdReviewDto,
	services *[]model.UserByIdFreelanceService,
) *dto.UserByIdDTO {
	return &dto.UserByIdDTO{
		ID:        model.ID,
		About:     model.About,
		CreatedAt: model.CreatedAt,
		FirstName: model.FirstName,
		Level:     model.Level,
		Surname:   model.Surname,
		Username:  model.Username,
		Avatar:    model.Avatar,
		Languages: model.Languages,
		Skills:    model.Skills,

		Rating:       &model.Rating,
		ReviewsCount: &model.Count,
		Reviews:      reviews,
		Services:     services,
	}
}
