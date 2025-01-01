package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"ofm_backend/cmd/ofm_backend/api/user/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "Error while creating mock database")
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	maxValue := 3
	os.Setenv("MAX_USER_BY_ID_REVIEWS", strconv.Itoa(maxValue))
	os.Setenv("MAX_USER_BY_ID_SERVICES", strconv.Itoa(maxValue))

	defer func() {
		os.Unsetenv("MAX_USER_BY_ID_REVIEWS")
		os.Unsetenv("MAX_USER_BY_ID_SERVICES")
		db.Close()
		sqlxDB.Close()
	}()
	
	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	
	userRepository := NewUserRepository(sqlxDB)

	testCases := []struct {
		name         string
		actualId     int64
		expectError  bool
		expectedUser *model.User
	}{
		{
			name:         "Success",
			actualId:     int64(1),
			expectError:  false,
			expectedUser: createUserModel(1, timeNow, 1),
		},
		{
			name:         "User Not Found",
			actualId:     int64(2),
			expectError:  true,
			expectedUser: createUserModel(1, timeNow, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectError {
				mock.ExpectQuery(regexp.QuoteMeta(utils.UserByIdQuery)).WillReturnError(main_utils.ErrUserNotFound)
			} else {
				rows := mock.NewRows([]string{"id", "username", "about", "created_at", "first_name", "level", "surname", "avatar", "languages", "reviews", "reviews_count", "rating"})
				addUserRows(rows, tc.expectedUser, tc.actualId)

				mock.ExpectQuery(regexp.QuoteMeta(utils.UserByIdQuery)).WillReturnRows(rows)
			}

			actualData, err := userRepository.GetUserById(int(tc.actualId))

			if tc.expectError {
				assert.Error(t, err, main_utils.ErrUserNotFound)
			} else {
				assert.NoError(t, err, "Error while getting user by id")
				assert.Equal(t, tc.expectedUser, actualData, "Data is not equal")
			}
		})
	}
}

func TestGetReviewsByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "Error while creating mock database")
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	maxValue := 3
	os.Setenv("MAX_USER_BY_ID_REVIEWS", strconv.Itoa(maxValue))
	os.Setenv("MAX_USER_BY_ID_SERVICES", strconv.Itoa(maxValue))

	defer func() {
		os.Unsetenv("MAX_USER_BY_ID_REVIEWS")
		os.Unsetenv("MAX_USER_BY_ID_SERVICES")
		db.Close()
		sqlxDB.Close()
	}()

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	timeNowStr := main_utils.ConvertTimeToSting(timeNow)

	userRepository := NewUserRepository(sqlxDB)

	testCases := []struct {
		name                 string
		endedAt              *string
		userId, lastReviewId int
		expectedData         *[]model.UserByIdReview
		mockRowsNum          int
	}{
		{
			name:         "Success",
			endedAt:      nil,
			userId:       1,
			lastReviewId: -1,
			expectedData: createReviewModels(timeNow, maxValue, 1),
			mockRowsNum:  maxValue,
		},
		{
			name:         "With cursor",
			endedAt:      &timeNowStr,
			userId:       1,
			lastReviewId: 3,
			expectedData: createReviewModels(timeNow, 3, 3),
			mockRowsNum:  5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := mock.NewRows([]string{"user_exists", "data"})
			createMockRows(rows, tc.expectedData)

			mock.ExpectQuery(regexp.QuoteMeta(utils.ReviewsByUserIdQuery)).
				WithArgs(tc.userId, tc.endedAt, int64(tc.lastReviewId), maxValue).
				WillReturnRows(rows)

			actualData, err := userRepository.GetReviewsByUserId(int(tc.userId), tc.endedAt, int64(tc.lastReviewId), maxValue)
			assert.NoError(t, err, "Error getting reviews by user id")

			assert.Equal(t, tc.expectedData, actualData, "Data should be equal")
		})
	}
}

func TestGetServicesByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "Error while creating mock database")
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	maxValue := 3
	os.Setenv("MAX_USER_BY_ID_REVIEWS", strconv.Itoa(maxValue))
	os.Setenv("MAX_USER_BY_ID_SERVICES", strconv.Itoa(maxValue))

	defer func() {
		os.Unsetenv("MAX_USER_BY_ID_REVIEWS")
		os.Unsetenv("MAX_USER_BY_ID_SERVICES")
		db.Close()
		sqlxDB.Close()
	}()

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)

	userRepository := NewUserRepository(sqlxDB)

	testCases := []struct {
		name                  string
		reviewsCount          int64
		userId, lastServiceId int
		expectedData          *[]model.UserByIdFreelanceService
		mockRowsNum           int
	}{
		{
			name:          "Success",
			reviewsCount:  int64(0),
			userId:        1,
			lastServiceId: -1,
			expectedData:  createServiceModels(timeNow, maxValue, 1),
			mockRowsNum:   maxValue,
		},
		{
			name:          "With cursor",
			reviewsCount:  int64(0),
			userId:        1,
			lastServiceId: 3,
			expectedData:  createServiceModels(timeNow, 3, 3),
			mockRowsNum:   5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := mock.NewRows([]string{"user_exists", "data"})
			createMockRows(rows, tc.expectedData)

			mock.ExpectQuery(regexp.QuoteMeta(utils.ServicesByUserIdQuery)).
				WithArgs(tc.userId, tc.reviewsCount, int64(tc.lastServiceId), maxValue).
				WillReturnRows(rows)

			actualData, err := userRepository.GetServicesByUserId(tc.userId, tc.reviewsCount, int64(tc.lastServiceId), maxValue)
			assert.NoError(t, err, "Error getting services by user id")

			assert.Equal(t, tc.expectedData, actualData, "Data should be equal")

		})
	}
}

func createMockRows[T any](rows *sqlmock.Rows, arr *[]T) {
	rows.AddRow(true, marshalToJSON[T](arr))
}

func createReviewModels(timeNow time.Time, num int, firstId int) *[]model.UserByIdReview {
	reviews := make([]model.UserByIdReview, num)

	for i := 0; i < len(reviews); i++ {
		reviews[i] = model.UserByIdReview{
			ID:           int64(i + firstId),
			Content:      "test",
			Rating:       0,
			CreatedAt:    timeNow,
			EndedAt:      timeNow,
			UserID:       int64(0),
			Username:     "test",
			Avatar:       nil,
			ServiceID:    int64(0),
			Price:        float64(0.0),
			ServiceImage: nil,
			Title:        "test",
		}
	}

	return &reviews
}

func createServiceModels(timeNow time.Time, num int, firstId int) *[]model.UserByIdFreelanceService {
	reviews := make([]model.UserByIdFreelanceService, num)

	for i := 0; i < len(reviews); i++ {
		reviews[i] = model.UserByIdFreelanceService{
			ID:           int64(i + firstId),
			CreatedAt:    timeNow,
			Description:  "test",
			Title:        "test",
			CategoryId:   int64(0),
			FreelancerId: int64(0),
			Image:        nil,
			ReviewsCount: int64(0),
			Rating:       float64(0.0),
			MinPrice:     float64(0.0),
		}
	}

	return &reviews
}

func addReviewsRows(rows *sqlmock.Rows, timeNow time.Time, num int) {
	reviews := make([]model.UserByIdReview, num)

	for i := 0; i < len(reviews); i++ {
		reviews[i] = model.UserByIdReview{
			ID:           int64(i + 1),
			Content:      "test",
			Rating:       0,
			CreatedAt:    timeNow,
			EndedAt:      timeNow,
			UserID:       int64(0),
			Username:     "test",
			Avatar:       nil,
			ServiceID:    int64(0),
			Price:        float64(0.0),
			ServiceImage: nil,
			Title:        "test",
		}
	}

	rows.AddRow(true, marshalToJSON(&reviews))
}

func addUserRows(rows *sqlmock.Rows, user *model.User, actualId int64) *sqlmock.Rows {
	rows.AddRow(
		actualId,
		user.Username,
		user.About,
		user.CreatedAt,
		user.FirstName,
		user.Level,
		user.Surname,
		user.Avatar,
		marshalToJSON(user.Languages),
		marshalToJSON(user.Skills),
		user.Count,
		user.Rating,
	)
	return rows
}

func marshalToJSON[T any](data *[]T) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func createUserModel(id int64, timeNow time.Time, num int) *model.User {
	testStr := "test"

	return &model.User{
		ID:        id,
		Username:  testStr,
		About:     &testStr,
		CreatedAt: timeNow,
		FirstName: testStr,
		Level:     float64(0.0),
		Surname:   testStr,
		Avatar:    &testStr,
		Languages: createLanguages(num),
		Skills:    createSkills(num),
		Count:     int64(0),
		Rating:    float64(0.0),
	}
}

func createLanguages(num int) *[]model.Language {
	languages := make([]model.Language, num)
	for i := 0; i < num; i++ {
		languages[i] = model.Language{
			ID:   i + 1,
			Name: "test",
		}
	}
	return &languages
}

func createSkills(num int) *[]model.Skill {
	skills := make([]model.Skill, num)
	for i := 0; i < num; i++ {
		skills[i] = model.Skill{
			ID:   int64(i + 1),
			Name: "test",
		}
	}
	return &skills
}
