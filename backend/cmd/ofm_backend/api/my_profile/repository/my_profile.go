package repository

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"

	"github.com/jmoiron/sqlx"
)

type MyProfileRepository interface {
	CreateTransaction() (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
	GetMyProfileOrders(params *dto.MyProfileParams) (*model.OrdersData, error)
	GetMyProfileServices(params *dto.MyProfileParams) (*model.ServicesData, error)
	GetMyProfileRequests(params *dto.MyProfileParams) (*model.RequestsData, error)
	GetMyProfileChatByOrderId(orderId, userId int) (*model.OrderChat, error)
	GetMyProfileOverviewByOrderId(orderId int, serviceFee float32, userId int) (*model.OrderOverview, error)
	GetMyProfileRequirementsByOrderId(orderId, userId int) ([]model.OrderQuestionAnswer, error)
	GetMyProfileDeliveryByOrderId(orderId, userId int, serviceFee float32) (*model.OrderDelivery, error)
	CompleteMyProfileOrderById(tx *sqlx.Tx, orderId, userId int) error
	AddOrderDelivery(tx *sqlx.Tx, orderId, userId int, message string) (int, error)
	AddDeliveryFiles(tx *sqlx.Tx, deliveryFiles []model.DeliveryFile) error
	AddOrderReview(tx *sqlx.Tx, reviewRequestBody model.ReviewRequestBody) (int, error)
	UpdateOrderReview(tx *sqlx.Tx, reviewRequestBody model.ReviewRequestBody) error
	GetOrderReviewByOrderId(orderId int) (*model.OrderReview, error)
	GetOrderStatusByOrderId(orderId int) (int, error)
	ChangeOrderStatusByOrderId(orderId, userId, statusId int) error
}
