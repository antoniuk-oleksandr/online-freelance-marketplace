package service

import (
	"ofm_backend/cmd/ofm_backend/api/order/body"
	"ofm_backend/cmd/ofm_backend/api/order/dto"
)

type OrderService interface {
	GetOrderById(id int) (*dto.OrderByIdResponse, error)
	SubmitOrderRequirements(orderRequirementsBody *body.OrderRequirementsBody) error
	AttachOrderFiles(orderRequirementsBody *body.OrderRequirementsBody) error
	CheckIfOrderRequirementsSubmitted(orderId int) (bool, error)
}
