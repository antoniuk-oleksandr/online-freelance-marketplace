package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/order/model"
	"ofm_backend/cmd/ofm_backend/api/order/queries"
	"ofm_backend/cmd/ofm_backend/enum"
	main_utils "ofm_backend/cmd/ofm_backend/utils"

	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) GetOrderById(
	id int,
) (*model.OrderByIdResponse, error) {
	var response model.OrderByIdResponse
	var orderJSON, freelanceJSON, freelanceQuestionsJSON []byte

	err := or.db.
		QueryRow(queries.GetOrderByIdQuery, id).
		Scan(&orderJSON, &freelanceJSON, &freelanceQuestionsJSON)
	if err != nil {
		return nil, err
	}

	if err := or.validateResponseJSON(
		orderJSON, freelanceJSON, freelanceQuestionsJSON,
	); err != nil {
		return nil, err
	}

	if err := or.unmarshalOrderByIdResponse(
		orderJSON, freelanceJSON, freelanceQuestionsJSON, &response,
	); err != nil {
		return nil, err
	}

	return &response, err
}

func (or *orderRepository) CreateTransaction() (*sqlx.Tx, error) {
	return or.db.Beginx()
}

func (or *orderRepository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (or *orderRepository) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (or *orderRepository) UpdateOrderCustomerMessageAndStatus(
	content string,
	orderId int,
	status enum.Status,
) error {
	_, err := or.db.Exec(queries.AddOrderCustomerMessageAndStatusQuery, content, status, orderId)
	return err
}

func (or *orderRepository) AddOrderQuestionAnswers(
	orderAnswerModels []model.OrderAnswer,
) error {
	_, err := or.db.NamedExec(queries.AddOrderAnswersQuery, orderAnswerModels)
	return err
}

func (or *orderRepository) AddOrderFiles(
	orderFilesData []model.OrderFile,
) error {
	_, err := or.db.NamedExec(queries.AddOrderFilesQuery, orderFilesData)
	return err
}

func (or *orderRepository) CheckIfOrderRequirementsSubmitted(
	orderId int, status enum.Status,
) (bool, error) {
	var submitted bool
	
	err := or.db.
		QueryRow(queries.CheckIfOrderRequirementsSubmittedQuery, status, orderId).
		Scan(&submitted)
	if err != nil {
		return false, nil
	}

	return submitted, nil
}

func (or *orderRepository) validateResponseJSON(
	orderJSON, freelanceJSON, freelanceQuestionsJSON []byte,
) error {
	if len(orderJSON) == 0 {
		return main_utils.ErrOrderNotFound
	}
	if len(freelanceJSON) == 0 {
		return main_utils.ErrFreelanceNotFound
	}
	if len(freelanceQuestionsJSON) == 0 {
		return main_utils.ErrFreelanceQuestionsNotFound
	}

	return nil
}

func (or *orderRepository) unmarshalOrderByIdResponse(
	orderJSON, freelanceJSON, freelanceQuestionsJSON []byte,
	response *model.OrderByIdResponse,
) error {
	if err := json.Unmarshal(orderJSON, &response.Order); err != nil {
		return err
	}
	if err := json.Unmarshal(freelanceJSON, &response.Service); err != nil {
		return err
	}
	if err := json.Unmarshal(freelanceQuestionsJSON, &response.ServiceQuestions); err != nil {
		return err
	}

	return nil
}
