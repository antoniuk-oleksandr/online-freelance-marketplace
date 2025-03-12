package service

import (
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/enum"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// AddOrder implements repository.PaymentRepository.
func (m *MockRepository) AddOrder(data body.DecryptedPaymentData) error {
	panic("unimplemented")
}

// AddPayment implements repository.PaymentRepository.
func (m *MockRepository) AddPayment(data body.DecryptedPaymentData) (int64, error) {
	panic("unimplemented")
}

// UpdatePaymentStatus implements repository.PaymentRepository.
func (m *MockRepository) UpdatePaymentStatus(status enum.Status, paymentId int64) (bool, error) {
	panic("unimplemented")
}
