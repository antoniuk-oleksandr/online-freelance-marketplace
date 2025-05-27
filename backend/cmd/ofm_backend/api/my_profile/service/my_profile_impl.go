package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/mapper"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
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
	params *dto.MyProfileParams,
) (*dto.OrdersResponse, error) {
	var err error
	params.Limit, err = helpers.GetOrdersPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	ordersDataModel, err := mps.myProfileRepository.GetMyProfileOrders(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapOrdersDataModelToOrdersResponseDto(ordersDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileServices(
	params *dto.MyProfileParams,
) (*dto.ServicesResponse, error) {
	var err error
	params.Limit, err = helpers.GetServicesPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	servicesDataModel, err := mps.myProfileRepository.GetMyProfileServices(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapServicesDataModelToServicesResponseDto(servicesDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileRequests(params *dto.MyProfileParams) (*dto.RequestsResponse, error) {
	var err error
	params.Limit, err = helpers.GetRequestsPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	requestsDataModel, err := mps.myProfileRepository.GetMyProfileRequests(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapRequestsDataModelToRequestsResponseDto(requestsDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileChatByOrderId(orderId int) (*model.OrderChat, error) {
	return mps.myProfileRepository.GetMyProfileChatByOrderId(orderId)
}
