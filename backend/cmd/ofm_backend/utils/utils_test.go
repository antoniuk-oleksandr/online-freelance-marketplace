package utils

import (
	"os"
	"testing"

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

	LoadEnvValues()

	databaseURL := os.Getenv("DATABASE_URL")
	assert.Equal(t, "test_database_url", databaseURL, "Expected DATABASE_URL to be loaded from .env file")

	err = os.Remove(".env")
	require.NoError(t, err, "Failed to remove .env file after test")
}

func TestContainsPositiveCase(t *testing.T) {
	arr := []string{"test1", "test2", "test3"}
	val := Contains(arr, "test2")
	assert.True(t, val, "Expected true for test2 in the array")
}

func TestContainsNegativeCase(t *testing.T) {
	arr := []string{"test1", "test2", "test3"}
	val := Contains(arr, "test4")
	assert.False(t, val, "Expected test4 to not be in the array")
}

func TestContainsSensitivityCase(t *testing.T){
	arr := []string{"test1", "Test2", "test3"}
	val := Contains(arr, "test2")
	assert.False(t, val, "Expected false for case sensitivity")
}

func TestContainsEmptyArrCase(t* testing.T){
	arr := []string{}
	val := Contains(arr, "test")
	assert.False(t, val, "Expected false for empty array")
}