package utils

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/middleware"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxReviews(t *testing.T) {
	testCases := []struct {
		name string
		env  string
	}{
		{
			name: "Default Value",
			env:  "",
		},
		{
			name: "Env Value",
			env:  "3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.env == "" {
				maxReviews := GetMaxReviews()
				assert.Equal(t, 5, maxReviews, "Default value should be 5")
			} else {
				os.Setenv("MAX_USER_BY_ID_REVIEWS", tc.env)

				maxReviews := GetMaxReviews()
				expectedNum, err := strconv.Atoi(tc.env)
				assert.NoError(t, err, "Error converting string to int")
				assert.Equal(t, expectedNum, maxReviews, fmt.Sprintf("Default value should be %s", tc.env))
			}
		})
	}
}

func TestGetMaxServices(t *testing.T) {
	testCases := []struct {
		name string
		env  string
	}{
		{
			name: "Default Value",
			env:  "",
		},
		{
			name: "Env Value",
			env:  "3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.env == "" {
				maxReviews := GetMaxServices()
				assert.Equal(t, 5, maxReviews, "Default value should be 5")
			} else {
				os.Setenv("MAX_USER_BY_ID_SERVICES", tc.env)

				maxReviews := GetMaxServices()
				expectedNum, err := strconv.Atoi(tc.env)
				assert.NoError(t, err, "Error converting string to int")
				assert.Equal(t, expectedNum, maxReviews, fmt.Sprintf("Default value should be %s", tc.env))
			}
		})
	}
}

func TestBuildServicesCursor(t *testing.T) {
	testCases := []struct {
		name          string
		reviewsCount  int64
		lastId        int64
		expectedPlain string
	}{
		{
			name:          "Standard case",
			reviewsCount:  10,
			lastId:        5,
			expectedPlain: "reviewsCount:10;lastId:5",
		},
		{
			name:          "Zero values",
			reviewsCount:  0,
			lastId:        0,
			expectedPlain: "reviewsCount:0;lastId:0",
		},
		{
			name:          "Large values",
			reviewsCount:  1000000,
			lastId:        999999,
			expectedPlain: "reviewsCount:1000000;lastId:999999",
		},
		{
			name:          "Negative values",
			reviewsCount:  -1,
			lastId:        -50,
			expectedPlain: "reviewsCount:-1;lastId:-50",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BuildServicesCursor(tc.reviewsCount, tc.lastId)
			expected := middleware.EncodeString(tc.expectedPlain)
			
			assert.Equal(t, expected, actual, "Cursor encoding mismatch")
		})
	}
}

func TestBuildReviewsCursor(t *testing.T) {
	testCases := []struct {
		name          string
		endedAt       string
		lastId        int64
		expectedPlain string 
	}{
		{
			name:          "Standard case",
			endedAt:       "2024-12-31T23:59:59.999+00:00",
			lastId:        5,
			expectedPlain: "endedAt:2024-12-31T23:59:59.999+00:00;lastId:5",
		},
		{
			name:          "Empty endedAt",
			endedAt:       "",
			lastId:        42,
			expectedPlain: "endedAt:;lastId:42",
		},
		{
			name:          "Zero lastId",
			endedAt:       "2023-01-01T00:00:00.000+00:00",
			lastId:        0,
			expectedPlain: "endedAt:2023-01-01T00:00:00.000+00:00;lastId:0",
		},
		{
			name:          "Negative lastId",
			endedAt:       "2024-06-15T12:34:56.789+00:00",
			lastId:        -1,
			expectedPlain: "endedAt:2024-06-15T12:34:56.789+00:00;lastId:-1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BuildReviewsCursor(tc.endedAt, tc.lastId)
			expected := middleware.EncodeString(tc.expectedPlain)
			assert.Equal(t, expected, actual, "Cursor encoding mismatch")
		})
	}
}

func TestParseReviewsCursor(t *testing.T) {
	testCases := []struct {
		name          string
		cursor        string
		expectedEnded *string
		expectedLast  int64
		expectedError error
	}{
		{
			name:          "Valid cursor",
			cursor:        middleware.EncodeString("endedAt:2024-12-31T23:59:59.999+00:00;lastId:42"),
			expectedEnded: ptr("2024-12-31T23:59:59.999+00:00"),
			expectedLast:  42,
			expectedError: nil,
		},
		{
			name:          "Empty fields",
			cursor:        middleware.EncodeString("endedAt:;lastId:0"),
			expectedEnded: nil,
			expectedLast:  -1,
			expectedError: main_utils.ErrInvalidCursor,
		},
		{
			name:          "Malformed cursor",
			cursor:        middleware.EncodeString("invalidFormat"),
			expectedEnded: nil,
			expectedLast:  -1,
			expectedError: utils.ErrInvalidCursor,
		},
		{
			name:          "Empty cursor",
			cursor:        "",
			expectedEnded: nil,
			expectedLast:  -1,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			endedAt, lastId, err := ParseReviewsCursor(tc.cursor)

			assert.Equal(t, tc.expectedEnded, endedAt, "endedAt mismatch")
			assert.Equal(t, tc.expectedLast, lastId, "lastId mismatch")
			assert.Equal(t, tc.expectedError, err, "error mismatch")
		})
	}
}

func TestParseServicesCursor(t *testing.T) {
	testCases := []struct {
		name            string
		cursor          string
		expectedReviews int64
		expectedLastId  int64
		expectedError   error
	}{
		{
			name:            "Valid cursor",
			cursor:          middleware.EncodeString("reviewsCount:10;lastId:42"),
			expectedReviews: 10,
			expectedLastId:  42,
			expectedError:   nil,
		},
		{
			name:            "Zero values",
			cursor:          middleware.EncodeString("reviewsCount:0;lastId:0"),
			expectedReviews: 0,
			expectedLastId:  0,
			expectedError:   nil,
		},
		{
			name:            "Malformed cursor",
			cursor:          middleware.EncodeString("invalidFormat"),
			expectedReviews: -1,
			expectedLastId:  -1,
			expectedError:   fmt.Errorf("input does not match format"),
		},
		{
			name:            "Empty cursor",
			cursor:          "",
			expectedReviews: -1,
			expectedLastId:  -1,
			expectedError:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reviewsCount, lastId, err := ParseServicesCursor(tc.cursor)

			assert.Equal(t, tc.expectedReviews, reviewsCount, "reviewsCount mismatch")
			assert.Equal(t, tc.expectedLastId, lastId, "lastId mismatch")

			if tc.expectedError == nil {
				assert.NoError(t, err, "Expected no error")
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Contains(t, err.Error(), tc.expectedError.Error(), "Error mismatch")
			}
		})
	}
}

func TestGetMoreServicesCursorData(t *testing.T) {
	mockServices := func(count int) *[]dto.ServiceByIdDto {
		services := make([]dto.ServiceByIdDto, count)
		for i := 0; i < count; i++ {
			services[i] = dto.ServiceByIdDto{
				ID:           int64(i + 1),
				ReviewsCount: int64(i * 10),
			}
		}
		return &services
	}

	testCases := []struct {
		name            string
		services        *[]dto.ServiceByIdDto
		maxServices     int
		expectedHasMore bool
		expectedCursor  *string
		expectedLength  int
	}{
		{
			name:            "Services nil",
			services:        nil,
			maxServices:     3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  0,
		},
		{
			name:            "Services less than max",
			services:        mockServices(2),
			maxServices:     3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  2,
		},
		{
			name:            "Services equal to max",
			services:        mockServices(3),
			maxServices:     3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  3,
		},
		{
			name:            "Services more than max",
			services:        mockServices(5),
			maxServices:     3,
			expectedHasMore: true,
			expectedCursor:  ptr(middleware.EncodeString("reviewsCount:30;lastId:4")),
			expectedLength:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hasMore, cursor := GetMoreServicesCursorData(tc.services, tc.maxServices)

			assert.Equal(t, tc.expectedHasMore, hasMore, "hasMore mismatch")

			if tc.expectedCursor != nil {
				assert.NotNil(t, cursor, "Cursor should not be nil")
				assert.Equal(t, *tc.expectedCursor, *cursor, "Cursor mismatch")
			} else {
				assert.Nil(t, cursor, "Cursor should be nil")
			}

			if tc.services != nil {
				assert.Equal(t, tc.expectedLength, len(*tc.services), "Services length mismatch")
			}
		})
	}
}

func TestGetMoreReviewsCursorData(t *testing.T) {
	mockReviews := func(count int) *[]dto.UserByIdReviewDto {
		reviews := make([]dto.UserByIdReviewDto, count)
		for i := 0; i < count; i++ {
			reviews[i] = dto.UserByIdReviewDto{
				ID:       int64(i + 1),
				EndedAt:  time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC),
				Content:  "Test review",
				Rating:   5,
			}
		}
		return &reviews
	}

	testCases := []struct {
		name             string
		reviews          *[]dto.UserByIdReviewDto
		maxReviews       int
		expectedHasMore  bool
		expectedCursor   *string
		expectedLength   int
	}{
		{
			name:            "Reviews nil",
			reviews:         nil,
			maxReviews:      3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  0,
		},
		{
			name:            "Reviews less than max",
			reviews:         mockReviews(2),
			maxReviews:      3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  2,
		},
		{
			name:            "Reviews equal to max",
			reviews:         mockReviews(3),
			maxReviews:      3,
			expectedHasMore: false,
			expectedCursor:  nil,
			expectedLength:  3,
		},
		{
			name:            "Reviews more than max",
			reviews:         mockReviews(5),
			maxReviews:      3,
			expectedHasMore: true,
			expectedCursor:  ptr(middleware.EncodeString("endedAt:2023-11-30 00:00:00;lastId:4")),
			expectedLength:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hasMore, cursor := GetMoreReviewsCursorData(tc.reviews, tc.maxReviews)

			assert.Equal(t, tc.expectedHasMore, hasMore, "hasMore mismatch")

			if tc.expectedCursor != nil {
				assert.NotNil(t, cursor, "Cursor should not be nil")
				assert.Equal(t, *tc.expectedCursor, *cursor, "Cursor mismatch")
			} else {
				assert.Nil(t, cursor, "Cursor should be nil")
			}

			if tc.reviews != nil {
				assert.Equal(t, tc.expectedLength, len(*tc.reviews), "Reviews length mismatch")
			}
		})
	}
}

func ptr(str string) *string {
	return &str
}
