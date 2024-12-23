package controller

import (
	"encoding/json"
	"net/http/httptest"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPublicKeyEndpointValidKey(t *testing.T) {
	expectedKey := "test_key"
	os.Setenv("PAYMENT_PUBLIC_KEY", expectedKey)
	defer os.Unsetenv("PAYMENT_PUBLIC_KEY")
	
	endpoint := "/api/v1/orders/public-payment-key"
	
	app := fiber.New()
	app.Get(endpoint, GetPublicPaymentKey)

	req := httptest.NewRequest("GET", endpoint, nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err, "Error when sending request should be nil")

	assert.Equal(t, resp.StatusCode, fiber.StatusOK, "Response status code should be 200")

	var body map[string]string
	err = json.NewDecoder(resp.Body).Decode(&body)
	
	assert.NoError(t, err, "Error when decoding response body should be nil")
	assert.Equal(t, body["key"], expectedKey, "Public key should be equal to the one set in the environment variable")
}

func TestGetPublicKeyEndpointMissingKey(t *testing.T){
	endpoint := "/api/v1/orders/public-payment-key"
	
	app := fiber.New()
	app.Get(endpoint, GetPublicPaymentKey)

	req := httptest.NewRequest("GET", endpoint, nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err, "Error when sending request should be nil")

	assert.Equal(t, resp.StatusCode, fiber.StatusBadRequest, "Response status code should be 400")

	var body map[string]string
	err = json.NewDecoder(resp.Body).Decode(&body)
	
	assert.NoError(t, err, "Error when decoding response body should be nil")
	assert.Equal(t, body["error"], utils.ErrInvalidPaymentPublicKey.Error(), "Error message should be equal to the one returned by the function")
}
