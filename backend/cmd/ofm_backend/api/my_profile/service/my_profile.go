package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
)

type MyProfileService interface {
	GetMyProfileOrders(params *dto.MyProfileParams) (*dto.OrdersResponse, error)
	GetMyProfileServices(params *dto.MyProfileParams) (*dto.ServicesResponse, error)
}
