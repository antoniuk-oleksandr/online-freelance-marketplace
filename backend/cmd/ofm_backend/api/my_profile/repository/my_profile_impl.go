package repository

import (
	"database/sql"
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/api/my_profile/queries"
	"ofm_backend/cmd/ofm_backend/enum"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/jmoiron/sqlx"
)

type myProfileRepository struct {
	db *sqlx.DB
}

func NewMyProfileRepository(db *sqlx.DB) MyProfileRepository {
	return &myProfileRepository{db: db}
}

func (mpr *myProfileRepository) GetMyProfileOrders(
	params *dto.MyProfileParams,
) (*model.OrdersData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileOrdersQuery, params.UserId, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.OrderTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.OrdersData{
		OrderTableData: *data,
		TotalPages:     totalPages,
	}, nil
}

func (mpr *myProfileRepository) GetMyProfileServices(
	params *dto.MyProfileParams,
) (*model.ServicesData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileServicesQuery, params.UserId, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.ServiceTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.ServicesData{
		ServiceTableData: *data,
		TotalPages:       totalPages,
	}, nil
}

func (mpr *myProfileRepository) GetMyProfileRequests(
	params *dto.MyProfileParams,
) (*model.RequestsData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileRequestsQuery, params.UserId, params.Status, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.RequestTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.RequestsData{
		RequestTableData: *data,
		TotalPages:       totalPages,
	}, nil
}

func (mpr *myProfileRepository) GetMyProfileChatByOrderId(orderId int, userId int) (*model.OrderChat, error) {
	rows, err := mpr.db.Queryx(queries.GetMyProfileChatByOrderIdQuery, orderId, userId)
	if err != nil {
		return nil, err
	}

	orderChat, err := helpers.ParseMyProfileChatByOrderIdFromRows(rows)
	if err != nil {
		return nil, err
	}

	return orderChat, nil
}

func (mpr *myProfileRepository) GetMyProfileOverviewByOrderId(
	orderId int, serviceFee float32, userId int,
) (*model.OrderOverview, error) {
	var orderOverview model.OrderOverview
	var serviceJson []byte
	var freelancerJson []byte
	var deliveryDate sql.NullTime

	if err := mpr.db.QueryRowx(queries.GetMyProfileOverviewByOrderIdQuery, orderId, serviceFee, userId).
		Scan(&orderOverview.Id, &deliveryDate, &orderOverview.CreatedAt,
			&orderOverview.Subtotal, &orderOverview.ServiceFee, &orderOverview.TotalPrice,
			&orderOverview.Status, &serviceJson, &freelancerJson,
		); err != nil {
		return nil, err
	}

	if deliveryDate.Valid {
		orderOverview.DeliveryDate = &deliveryDate.Time
	}

	if err := json.Unmarshal(serviceJson, &orderOverview.OrderOverviewService); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(freelancerJson, &orderOverview.OrderOverviewFreelancer); err != nil {
		return nil, err
	}

	return &orderOverview, nil
}

func (mpr *myProfileRepository) GetMyProfileRequirementsByOrderId(
	orderId int, userId int,
) ([]model.OrderQuestionAnswer, error) {
	var orderQuestionArray = make([]model.OrderQuestionAnswer, 0)
	if err := mpr.db.Select(
		&orderQuestionArray, queries.GetMyProfileRequirementsByOrderIDQuery, orderId, userId,
	); err != nil {
		return nil, err
	}

	return orderQuestionArray, nil
}

func (mpr *myProfileRepository) GetMyProfileDeliveryByOrderId(
	orderId int, userId int, serviceFee float32,
) (*model.OrderDelivery, error) {
	var (
		orderDelivery                 model.OrderDelivery
		orderDeliveryDataJson         []byte
		orderDeliveryFreelancerJson   []byte
		orderDeliveryPaymentJson      []byte
		orderDeliveryCancellationJson []byte
	)

	err := mpr.db.QueryRowx(queries.GetMyProfileDeliveryByOrderIdQuery, orderId, userId, serviceFee).
		Scan(
			&orderDelivery.Status,
			&orderDeliveryDataJson,
			&orderDeliveryFreelancerJson,
			&orderDeliveryPaymentJson,
			&orderDeliveryCancellationJson,
		)
	if err != nil {
		return nil, err
	}

	switch orderDelivery.Status {
	case enum.Cancelled:
		var orderDeliveryPayment *model.OrderDeliveryPayment
		if err := json.Unmarshal(orderDeliveryPaymentJson, &orderDeliveryPayment); err != nil {
			return nil, err
		}

		var orderDeliveryCancellation *model.OrderDeliveryCancellation
		if err := json.Unmarshal(orderDeliveryCancellationJson, &orderDeliveryCancellation); err != nil {
			return nil, err
		}
		orderDelivery.Payment = orderDeliveryPayment
		orderDelivery.Cancellation = orderDeliveryCancellation
	case enum.Completed:
		var orderDeliveryData *model.OrderDeliveryData
		if err := json.Unmarshal(orderDeliveryDataJson, &orderDeliveryData); err != nil {
			return nil, err
		}
		orderDelivery.DeliveryData = orderDeliveryData
	case enum.InProgress:
		var orderDeliveryFreelancer *model.OrderDeliveryFreelancer
		if err := json.Unmarshal(orderDeliveryFreelancerJson, &orderDeliveryFreelancer); err != nil {
			return nil, err
		}
		orderDelivery.Freelancer = orderDeliveryFreelancer

		var orderDeliveryData *model.OrderDeliveryData
		if err := json.Unmarshal(orderDeliveryDataJson, &orderDeliveryData); err != nil {
			return nil, err
		}
		orderDelivery.DeliveryData = orderDeliveryData
	}

	return &orderDelivery, nil
}

func (mpr *myProfileRepository) CreateTransaction() (*sqlx.Tx, error) {
	return mpr.db.Beginx()
}

func (mpr *myProfileRepository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (mpr *myProfileRepository) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (mpr *myProfileRepository) CompleteMyProfileOrderById(tx *sqlx.Tx, orderId int, userId int) error {
	result, err := tx.Exec(queries.CompleteOrderQuery, orderId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return utils.ErrCompletingOrder
	}

	return nil
}

func (mpr *myProfileRepository) AddOrderDelivery(tx *sqlx.Tx, orderId int, userId int, message string) (int, error) {
	var deliveryId int

	err := tx.QueryRowx(queries.AddOrderDeliveryQuery, orderId, message, userId).Scan(&deliveryId)
	if err != nil {
		return 0, err
	}

	return deliveryId, nil
}

func (mpr *myProfileRepository) AddDeliveryFiles(tx *sqlx.Tx, deliveryFiles []model.DeliveryFile) error {
	rows, err := tx.NamedExec(queries.AddDeliveryFilesQuery, deliveryFiles)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != int64(len(deliveryFiles)) {
		return utils.ErrAddingDeliveryFiles
	}

	return nil
}

func (mpr *myProfileRepository) AddOrderReview(tx *sqlx.Tx, reviewRquestBody model.ReviewRequestBody) (int, error) {
	var reviewId int
	err := tx.QueryRowx(
		queries.AddReviewQuery, reviewRquestBody.OrderId, reviewRquestBody.UserId,
		reviewRquestBody.ReviewMessage, reviewRquestBody.Rating,
	).Scan(&reviewId)

	return reviewId, err
}

func (mpr *myProfileRepository) UpdateOrderReview(tx *sqlx.Tx, reviewRequestBody model.ReviewRequestBody) error {
	result, err := tx.Exec(
		queries.UpdateOrderReviewQuery, reviewRequestBody.ReviewId,
		reviewRequestBody.UserId, reviewRequestBody.OrderId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsAffected != 1 {
		return utils.ErrUnexpectedError
	}

	return nil
}

func (mpr *myProfileRepository) GetOrderReviewByOrderId(orderId int) (*model.OrderReview, error) {
	var orderReview model.OrderReview
	var customerJSON []byte
	if err := mpr.db.QueryRowx(queries.GetOrderReviewQuery, orderId).
		Scan(&orderReview.Content, &orderReview.Rating, &customerJSON); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(customerJSON, &orderReview.Customer); err != nil {
		return nil, err
	}

	return &orderReview, nil
}

func (mpr *myProfileRepository) GetOrderStatusByOrderId(orderId int) (int, error) {
	var status int
	err := mpr.db.Get(&status, queries.GetOrderStatusByOrderIdQuery, orderId)
	return status, err
}

func (mpr *myProfileRepository) ChangeOrderStatusByOrderId(orderId int, userId int, statusId int) error {
	result, err := mpr.db.Exec(queries.UpdateOrderStatusByOrderIdQuery, statusId, orderId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return utils.ErrUnexpectedError
	}

	return err
}
