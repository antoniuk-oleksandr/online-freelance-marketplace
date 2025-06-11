package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
)

type MyProfileService interface {
	GetMyProfileOrders(params *dto.MyProfileParams) (*dto.OrdersResponse, error)
	GetMyProfileServices(params *dto.MyProfileParams) (*dto.ServicesResponse, error)
	GetMyProfileRequests(params *dto.MyProfileParams) (*dto.RequestsResponse, error)
	GetMyProfileChatByOrderId(orderId int, userId int) (*dto.OrderChat, error)
}
