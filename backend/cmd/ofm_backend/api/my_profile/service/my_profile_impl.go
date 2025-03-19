package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/mapper"
	"ofm_backend/cmd/ofm_backend/api/my_profile/repository"
)

type myProfileService struct {
	myProfileRepository repository.MyProfileRepository
}

func NewMyProfileService(repo repository.MyProfileRepository) MyProfileService {
	return &myProfileService{
		myProfileRepository: repo,
	}
}

func (mps *myProfileService) GetMyProfileOrders(
	params *dto.OrdersPaginationParams,
) (*dto.OrdersResponse, error) {
	var err error
	params.OrdersPerPage, err = helpers.GetOrdersPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileOrdersOffset(params)

	ordersDataModel, err := mps.myProfileRepository.GetMyProfileOrders(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapOrdersDataModelToOrdersResponseDto(ordersDataModel)

	return ordersReponse, nil
}
