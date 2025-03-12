package repository

import (
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/model"
	"ofm_backend/cmd/ofm_backend/enum"

	"github.com/jmoiron/sqlx"
)

type PaymentRepository interface {
	AddPayment(data body.DecryptedPaymentData) (int64, error)
	UpdatePaymentStatus(status enum.Status, paymentId int64) (bool, error)
	AddOrder(data body.DecryptedPaymentData) (int64, error)
	GetPaymentReceipt(orderId int64, serviceFee string) (*model.PaymentReceipt, error)
	CreateTransaction() (*sqlx.Tx, error)
	RollBackTransaction(tx *sqlx.Tx) error
	CommitTransaction(tx *sqlx.Tx) error
}
