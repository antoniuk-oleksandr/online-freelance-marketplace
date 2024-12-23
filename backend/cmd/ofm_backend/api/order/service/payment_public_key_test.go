package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentPublicKeyValidKey(t *testing.T){
	expectedKey := "key-test"
	os.Setenv("PAYMENT_PUBLIC_KEY", expectedKey)
	defer os.Unsetenv("PAYMENT_PUBLIC_KEY")
	
	actualKey, err := GetPaymentPublicKey()
	
	assert.Equal(t, expectedKey, actualKey, "Actual and expected keys should be equal")
	assert.NoError(t, err, "Error should be nil")
}

func TestGetPaymentPublicKeyMissingKey(t *testing.T){
	actualKey, err := GetPaymentPublicKey()
	
	assert.Error(t, err, "Error should not be nil when key is missing")
	assert.Empty(t, actualKey, "Actual key should be empty when key is missing")
}