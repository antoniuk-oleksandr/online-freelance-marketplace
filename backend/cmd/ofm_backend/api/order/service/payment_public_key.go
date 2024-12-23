package service

import (
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
)

func GetPaymentPublicKey() (string, error) {
	key := os.Getenv("PAYMENT_PUBLIC_KEY")
	if key == "" {
		return "", utils.ErrInvalidPaymentPublicKey
	}

	return key, nil
}
