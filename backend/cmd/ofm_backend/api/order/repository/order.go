package repository

import (
	"ofm_backend/cmd/ofm_backend/api/order/model"
	"ofm_backend/cmd/ofm_backend/enum"

	"github.com/jmoiron/sqlx"
)

type OrderRepository interface {
	GetOrderById(id int) (*model.OrderByIdResponse, error)
	CreateTransaction() (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
	AddOrderQuestionAnswers([]model.OrderAnswer) error
	UpdateOrderCustomerMessageAndStatus(content string, orderId int, status enum.Status) error
	AddOrderFiles(orderFilesData []model.OrderFile) error
	CheckIfOrderRequirementsSubmitted(orderId int, status enum.Status) (bool, error)
	validateResponseJSON(orderJSON, freelanceJSON, freelanceQuestionsJSON []byte) error
	unmarshalOrderByIdResponse(
		orderJSON, freelanceJSON, freelanceQuestionsJSON []byte,
		response *model.OrderByIdResponse,
	) error
}
