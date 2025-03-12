package mapper

import (
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	modelPac "ofm_backend/cmd/ofm_backend/api/user/model"
)

func MapUserByIdReviewModelsToReviewUserDTOs(reviews *[]modelPac.UserByIdReview) *[]dto.ReviewUserDTO {
	var reviewUsers []dto.ReviewUserDTO

	for _, review := range *reviews {
		user := dto.ReviewUserDTO{
			ID:       review.UserID,
			Avatar:   review.Avatar,
			Username: review.Username,
		}

		reviewUsers = append(reviewUsers, user)
	}

	return &reviewUsers
}

func MapUserByIdReviewModelToReviewServiceDTOs(reviews *[]modelPac.UserByIdReview) *[]dto.ReviewService {
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
	reviews *[]modelPac.UserByIdReview,
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

func MapUserByIdServiceModelsToDTO(
	services *[]modelPac.UserByIdFreelanceService,
) *[]dto.ServiceByIdDto {
	servicesArr := make([]dto.ServiceByIdDto, len(*services))

	for i, elem := range *services {
		servicesArr[i] = dto.ServiceByIdDto{
			ID:           elem.ID,
			Title:        elem.Title,
			Image:        elem.Image,
			ReviewsCount: elem.ReviewsCount,
			Rating:       elem.Rating,
			MinPrice:     elem.MinPrice,
		}
	}

	return &servicesArr
}

func MapUserByIdModelToDTO(
	model *modelPac.User,
	reviews *[]dto.UserByIdReviewDto,
	services *[]dto.ServiceByIdDto,
) *dto.UserByIdTO {
	var reviewsArr []dto.UserByIdReviewDto
	if reviews == nil || (len(*reviews) == 0) {
		reviewsArr = make([]dto.UserByIdReviewDto, 0)
	} else {
		reviewsArr = *reviews
	}

	var servicesArr []dto.ServiceByIdDto
	if services == nil || (len(*services) == 0) {
		servicesArr = make([]dto.ServiceByIdDto, 0)
	} else {
		servicesArr = *services
	}

	return &dto.UserByIdTO{
		ID:           model.ID,
		About:        model.About,
		CreatedAt:    model.CreatedAt,
		FirstName:    model.FirstName,
		Level:        model.Level,
		Surname:      model.Surname,
		Username:     model.Username,
		Avatar:       model.Avatar,
		Languages:    utils.MapFilterParamModelsToDTOs(model.Languages),
		Skills:       utils.MapFilterParamModelsToDTOs(model.Skills),
		Rating:       &model.Rating,
		ReviewsCount: &model.Count,
		Reviews:      reviewsArr,
		Services:     servicesArr,
	}
}
