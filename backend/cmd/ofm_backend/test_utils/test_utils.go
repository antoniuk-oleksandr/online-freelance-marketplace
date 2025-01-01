package test_utils

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SetupFiberApp(route string, handler fiber.Handler) *fiber.App {
	app := fiber.New()
	app.Get(route, handler)

	return app
}

func PerformRequest(t *testing.T, app *fiber.App, method, url string) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	assert.NoError(t, err, "Error creating request")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err, "Error testing request")

	return resp
}

func ParseResponseBody[T any](t *testing.T, resp *http.Response) T {
	var result T
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err, "Error parsing response body")

	return result
}