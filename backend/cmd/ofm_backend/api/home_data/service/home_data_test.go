package service

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/dto"
	"ofm_backend/cmd/ofm_backend/api/home_data/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetHomeData() (*model.HomeData, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	} else if args.Get(0) != nil {
		return args.Get(0).(*model.HomeData), nil
	}

	return nil, nil
}

func TestGetData(t *testing.T) {
	testCases := []struct {
		name         string
		expectedData *dto.HomeData
		mockData     *model.HomeData
		mockError    error
	}{
		{
			name:         "Valid data",
			expectedData: ptr(createHomeDataDto()),
			mockData:     ptr(createHomeDataModel()),
			mockError:    nil,
		},
		{
			name:         "Repository error",
			expectedData: ptr(createHomeDataDto()),
			mockData:     nil,
			mockError:    utils.ErrUnexpectedError,
		},
		{
			name:         "Nil data",
			expectedData: nil,
			mockData:     nil,
			mockError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(mockRepository)
			mockRepository.On("GetHomeData").Return(tc.mockData, tc.mockError)

			homeService := NewHomeService(mockRepository)

			actualData, err := homeService.GetHomeData()
			if tc.mockError == nil && actualData != nil {
				assert.NoError(t, err, "Error should be nil")
				assert.Equal(t, tc.expectedData, actualData, "Expected and actual data are not equal")
			} else if tc.mockError != nil {
				assert.Equal(t, tc.mockError, err, "Expected and actual error are not equal")
			} else {
				assert.Equal(t, utils.ErrUnexpectedError, err, "Expected and actual error are not equal")
			}

			mockRepository.AssertExpectations(t)
		})
	}
}

func createHomeDataDto() dto.HomeData {
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

func createHomeDataModel() model.HomeData {
	return model.HomeData{
		BestFreelancers: []model.BestFreelancer{
			{
				Id:                0,
				FirstName:         "test",
				Surname:           "test",
				Rating:            0,
				CompletedProjects: 0,
				Avatar:            ptr("test"),
			},
		},
		BestFreelances: []model.BestFreelance{
			{
				Id:          0,
				Title:       "test",
				Description: "test",
				Image:       ptr("test"),
			},
		},
		KeyMetrics: model.KeyMetrics{
			FreelancesAvailable: 0,
			ProjectsCompleted:   0,
			AvgRating:           0,
		},
		BestReviews: []model.BestReview{
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
