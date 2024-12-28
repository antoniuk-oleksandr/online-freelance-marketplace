package utils

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/internal/middleware"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxReviewsValueValidCase(t *testing.T) {
	expectedVal := 10

	os.Setenv("MAX_FREELANCE_BY_ID_REVIEWS", fmt.Sprintf("%d", expectedVal))
	defer os.Unsetenv("MAX_FREELANCE_BY_ID_REVIEWS")

	actualVal := GetMaxReviewsValue()

	assert.Equal(t, expectedVal, actualVal, "Expected value should be equal to actual value")
}

func TestGetMaxReviewsValueDefaultValue(t *testing.T) {
	expectedVal := 5
	actualVal := GetMaxReviewsValue()
	assert.Equal(t, expectedVal, actualVal, "Expected value should be equal to actual value")
}

func TestBuildReviewsCursor(t *testing.T) {
	layout := "2006-01-02 15:04:05.999999999"
	timeStr := "2023-12-26 00:42:33.925400897"

	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	var lastID int64 = 5

	cursor := fmt.Sprintf("reviewsCursor:%s;lastID:%d", timeStr, lastID)
	expectedEncodedCursor := middleware.EncodeString(cursor)

	lastReview := model.Review{
		ID:      lastID,
		EndedAt: parsedTime,
	}

	actualCursor := BuildReviewsCursor(lastReview.EndedAt, lastReview.ID)

	assert.Equal(t, expectedEncodedCursor, *actualCursor, "Expected cursor should be equal to actual cursor")
}

func TestGetDataFromReviewsCursor_Success(t *testing.T) {
	expectedData := "2023-12-26 00:42:33.925400897"
	var expectedLastID int64 = 5
	cursor := fmt.Sprintf("reviewsCursor:%s;lastID:%d", expectedData, expectedLastID)
	encodedCursor := middleware.EncodeString(cursor)

	actualData, actualLastID, err := GetDataFromReviewsCursor(encodedCursor)

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, expectedData, actualData, "Expected data should be equal to actual data")
	assert.Equal(t, expectedLastID, actualLastID, "Expected lastID should be equal to actual lastID")
}
