package service

import (
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/api/user/mapper"
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"ofm_backend/cmd/ofm_backend/api/user/repository"
	"ofm_backend/cmd/ofm_backend/api/user/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) GetReviewsByUserId(id int, cursor string) (*dto.ReviewsResponse, error) {
	maxReviews := utils.GetMaxReviews()
	endedAt, lastId, err := utils.ParseReviewsCursor(cursor)
	if err != nil {
		return nil, err
	}

	userByIdReviewsModels, err := us.userRepository.GetReviewsByUserId(id, endedAt, lastId, maxReviews+1)

	if err != nil {
		return nil, err
	}

	reviewsUsers := mapper.MapUserByIdReviewModelsToReviewUserDTOs(userByIdReviewsModels)
	reviewsServices := mapper.MapUserByIdReviewModelToReviewServiceDTOs(userByIdReviewsModels)
	reviews := mapper.MapReviewUsersServicesToUserByIdReviewDTOs(
		userByIdReviewsModels,
		reviewsUsers,
		reviewsServices,
	)

	hasMoreReviews, reviewsCursor := utils.GetMoreReviewsCursorData(reviews, maxReviews)

	return &dto.ReviewsResponse{
		Reviews:        main_utils.AddServerURLToFiles(reviews),
		HasMoreReviews: hasMoreReviews,
		ReviewsCursor:  reviewsCursor,
	}, nil
}

func (us *userService) GetServicesByUserId(id int, cursor string) (*dto.ServicesResponse, error) {
	maxServices := utils.GetMaxServices()

	reviewsCount, lastId, err := utils.ParseServicesCursor(cursor)
	if err != nil {
		return nil, err
	}

	serviceModels, err := us.userRepository.GetServicesByUserId(id, reviewsCount, lastId, maxServices+1)
	if err != nil {
		return nil, main_utils.ErrUserNotFound
	}

	serviceDTOs := mapper.MapUserByIdServiceModelsToDTO(serviceModels)

	hasMoreServices, servicesCursor := utils.GetMoreServicesCursorData(serviceDTOs, maxServices)

	return &dto.ServicesResponse{
		Services:        main_utils.AddServerURLToFiles(serviceDTOs),
		HasMoreServices: hasMoreServices,
		ServicesCursor:  servicesCursor,
	}, nil
}

func (us *userService) GetUserById(id int) (*dto.UserByIDResponse, error) {
	maxReviews := utils.GetMaxReviews()
	maxServices := utils.GetMaxServices()

	userModel, err := us.userRepository.GetUserById(id)
	if err != nil {
		return nil, main_utils.ErrUserNotFound
	}

	userByIdReviewsModels, _ := us.userRepository.GetReviewsByUserId(id, nil, -1, maxReviews+1)
	if userByIdReviewsModels == nil {
		emptyArr := make([]model.UserByIdReview, 0)
		userByIdReviewsModels = &emptyArr
	}

	reviewsUsers := mapper.MapUserByIdReviewModelsToReviewUserDTOs(userByIdReviewsModels)
	reviewsServices := mapper.MapUserByIdReviewModelToReviewServiceDTOs(userByIdReviewsModels)
	reviews := mapper.MapReviewUsersServicesToUserByIdReviewDTOs(
		userByIdReviewsModels,
		reviewsUsers,
		reviewsServices,
	)

	serviceModels, _ := us.userRepository.GetServicesByUserId(id, -1, -1, maxServices+1)
	if serviceModels == nil {
		emptyArr := make([]model.UserByIdFreelanceService, 0)
		serviceModels = &emptyArr
	}

	serviceDTOs := mapper.MapUserByIdServiceModelsToDTO(serviceModels)

	hasMoreReviews, reviewsCursor := utils.GetMoreReviewsCursorData(reviews, maxReviews)
	hasMoreServices, servicesCursor := utils.GetMoreServicesCursorData(serviceDTOs, maxServices)

	userDto := mapper.MapUserByIdModelToDTO(
		userModel,
		reviews,
		serviceDTOs,
	)

	return &dto.UserByIDResponse{
		User:            main_utils.AddServerURLToFiles(userDto),
		HasMoreReviews:  hasMoreReviews,
		ReviewsCursor:   reviewsCursor,
		HasMoreServices: hasMoreServices,
		ServicesCursor:  servicesCursor,
	}, nil
}
