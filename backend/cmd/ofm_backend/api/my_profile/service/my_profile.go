package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
)

type MyProfileService interface {
	GetMyProfileOrders(params *dto.OrdersPaginationParams) (*dto.OrdersResponse, error)
}
