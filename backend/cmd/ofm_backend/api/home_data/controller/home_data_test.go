package controller

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/dto"
	"ofm_backend/cmd/ofm_backend/test_utils"
	"ofm_backend/cmd/ofm_backend/utils"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) GetHomeData() (*dto.HomeData, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	} 
	return args.Get(0).(*dto.HomeData), args.Error(1)
}

func TestGetHomeData(t *testing.T) {
	testCases := []struct {
		name         string
		expectedData *dto.HomeData
		mockError    error
		expectedCode int
	}{
		{
			name:         "Valid data",
			expectedData: ptr(createHomeData()),
			mockError:    nil,
			expectedCode: fiber.StatusOK,
		},
		{
			name:         "Service error",
			expectedData: nil,
			mockError:    utils.ErrUnexpectedError,
			expectedCode: fiber.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mockService)
			
			homeController := NewHomeController(mockService)
			mockApp := test_utils.SetupFiberApp("/api/v1/home-data", homeController.GetHomeData)
			mockService.On("GetHomeData").Return(tc.expectedData, tc.mockError)

			resp := test_utils.PerformRequest(t, mockApp, "GET", "/api/v1/home-data")

			assert.Equal(t, tc.expectedCode, resp.StatusCode)
			if resp.StatusCode == fiber.StatusOK {
				actualBody := test_utils.ParseResponseBody[dto.HomeData](t, resp)
				assert.Equal(t, tc.expectedData, &actualBody, "Expected and actual body are not equal")
			} else {
				actualBody := test_utils.ParseResponseBody[utils.ErrorResponse](t, resp)
				assert.Contains(t, actualBody.Error, utils.ErrUnexpectedError.Error(), "Error message is not as expected")
			}
			
			mockService.AssertExpectations(t)
		})
	}
}

func createHomeData() dto.HomeData {
	return dto.HomeData{
		BestFreelancers: []dto.BestFreelancer{
			{
				Id:                0,
				FirstName:         "test",
				Surname:           "test",
				Rating:            0,
				CompletedProjects: 0,
				Avatar:            ptr("test"),
			},
		},
		BestFreelances: []dto.BestFreelance{
			{
				Id:          0,
				Title:       "test",
				Description: "test",
				Image:       ptr("test"),
			},
		},
		KeyMetrics: dto.KeyMetrics{
			FreelancesAvailable: 0,
			ProjectsCompleted:   0,
			AvgRating:           0,
		},
		BestReviews: []dto.BestReview{
			{
				FirstName: "test",
				Surname:   "test",
				Content:   "test",
				Rating:    0,
			},
		},
	}
}

func ptr[T any](object T) *T {
	return &object
}
