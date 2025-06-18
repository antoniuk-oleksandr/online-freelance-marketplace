package service

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
)

type MyProfileService interface {
	GetMyProfileOrders(params *dto.MyProfileParams) (*dto.OrdersResponse, error)
	GetMyProfileServices(params *dto.MyProfileParams) (*dto.ServicesResponse, error)
	GetMyProfileRequests(params *dto.MyProfileParams) (*dto.RequestsResponse, error)
	GetMyProfileChatByOrderId(orderId, userId int) (*dto.OrderChat, error)
	GetMyProfileOverviewByOrderId(orderId, userId int) (*dto.OrderOverview, error)
	GetMyProfileRequirementsByOrderId(orderId, userId int) ([]model.OrderQuestionAnswer, error)
	GetMyProfileDeliveryByOrderId(orderId, userId int) (*model.OrderDelivery, error)
	CompleteOrderWithDelivery(orderId, userId int, orderDeliveryBody model.OrderDeliveryBody) (*model.OrderDelivery, error)
	AddOrderReview(reviewRequestBody model.ReviewRequestBody) (*model.OrderReview, error)
	GetOrderReviewByOrderId(orderId int) (*model.OrderReview, error)
	UpdateOrderStatusByOrderId(orderId, userId int, statusId int) error
}
