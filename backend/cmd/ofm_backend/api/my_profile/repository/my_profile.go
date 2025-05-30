package repository

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
)

type MyProfileRepository interface {
	GetMyProfileOrders(params *dto.MyProfileParams) (*model.OrdersData, error)
	GetMyProfileServices(params *dto.MyProfileParams) (*model.ServicesData, error)
	GetMyProfileRequests(params *dto.MyProfileParams) (*model.RequestsData, error)
	GetMyProfileChatByOrderId(orderId int) (*model.OrderChat, error)
}
