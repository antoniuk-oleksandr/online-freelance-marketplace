package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/home_data/model"
	"ofm_backend/cmd/ofm_backend/api/home_data/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetHomeData(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "Error while creating mock database")
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer func() {
		db.Close()
		sqlxDB.Close()
	}()

	testCases := []struct {
		name         string
		expectedData *model.HomeData
		mockResponse []byte
		mockError    error
	}{
		{
			name:         "Valid data",
			expectedData: ptr(createHomeDataModel()),
			mockResponse: convertToJSON(createHomeDataModel()),
			mockError:    nil,
		},
		{
			name:         "Error",
			expectedData: nil,
			mockResponse: nil,
			mockError:    main_utils.ErrUnexpectedError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"data"}).AddRow(tc.mockResponse)
			mock.ExpectQuery(regexp.QuoteMeta(utils.GetHomeDataQuery)).
				WillReturnRows(rows).
				WillReturnError(tc.mockError)

			homeRepository := NewHomeRepository(sqlxDB)
			actualData, err := homeRepository.GetHomeData()

			if tc.mockError == nil && actualData != nil {
				assert.NoError(t, err, "Error while getting home data")
				assert.Equal(t, tc.expectedData, actualData, "Data is not equal")
			} else {
				assert.Equal(t, tc.mockError, err, "Error is not equal")
			}
		})
	}
}

func convertToJSON[T any](data T) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return jsonData
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
