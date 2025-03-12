package service

import (
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/dto"
)

type PaymentService interface {
	ProcessPayment(data body.EncryptedPaymentData, username string) (dto.PaymentResponse, error)
	CreatePayment(decryptedData *body.DecryptedPaymentData) error
	CreateOrder(decryptedData body.DecryptedPaymentData, paymentResponse *dto.PaymentResponse) (int64, error)
	ProcessTransaction(paymentId int64) error
	SendOrderReceipt(username string, orderId int64, userTimezone string) error
}
