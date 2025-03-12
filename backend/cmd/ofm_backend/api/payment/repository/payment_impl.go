package repository

import (
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/model"
	"ofm_backend/cmd/ofm_backend/api/payment/utils"
	"ofm_backend/cmd/ofm_backend/enum"

	"github.com/jmoiron/sqlx"
)

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (pRepo *paymentRepository) AddPayment(
	data body.DecryptedPaymentData,
) (int64, error) {
	var lastId int64

	err := pRepo.db.
		QueryRow(utils.PaymentQuery, data.CardCredentials.CardNumber, enum.Pending).
		Scan(&lastId)

	return lastId, err
}

func (pRepo *paymentRepository) UpdatePaymentStatus(
	status enum.Status,
	paymentId int64,
) (bool, error) {
	res, err := pRepo.db.Exec(utils.UpdatePaymentStatusQuery, status, paymentId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return false, err
	}

	return true, nil
}

func (pRepo *paymentRepository) AddOrder(
	data body.DecryptedPaymentData,
) (int64, error) {
	var orderId int64

	err := pRepo.db.QueryRow(
		utils.AddOrderQuery, data.Username, data.ServiceId,
		data.PackageId, enum.Incomplete, data.PaymentId,
	).Scan(&orderId)

	return orderId, err
}

func (pRepo *paymentRepository) GetPaymentReceipt(
	orderId int64,
	serviceFee string,
) (*model.PaymentReceipt, error) {
	var paymentReceipt model.PaymentReceipt

	err := pRepo.db.Get(&paymentReceipt, utils.GetPaymentReceiptQuery, serviceFee, orderId)
	if err != nil {
		return nil, err
	}

	return &paymentReceipt, nil
}

func (pRepo *paymentRepository) CreateTransaction() (*sqlx.Tx, error) {
	return pRepo.db.Beginx()
}

func (pRepo *paymentRepository) RollBackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (pRepo *paymentRepository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}
