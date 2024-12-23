package utils_test

import (
	"os"
	"testing"
	"ofm_backend/cmd/ofm_backend/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadEnvValues_Success(t *testing.T) {
	originalEnv, err := os.ReadFile(".env")
	if err == nil {
		defer os.WriteFile(".env", originalEnv, 0644)
	} else {
		defer os.Remove(".env")
	}

	err = os.WriteFile(".env", []byte("DATABASE_URL=test_database_url"), 0644)
	require.NoError(t, err, "Failed to write test .env file")

	utils.LoadEnvValues()

	databaseURL := os.Getenv("DATABASE_URL")
	assert.Equal(t, "test_database_url", databaseURL, "Expected DATABASE_URL to be loaded from .env file")

	err = os.Remove(".env")
	require.NoError(t, err, "Failed to remove .env file after test")
}
