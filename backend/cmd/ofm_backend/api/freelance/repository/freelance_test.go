package repository

import (
	"encoding/json"
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/api/freelance/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetGetFreelanceServiceByIdReviews_WithCursor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	endedAt := time.Now().Add(time.Hour * 24)
	cursorData := main_utils.GetCurrentTime()
	var lastID int64 = 5
	id := 1
	maxReviews := 2

	expectedData := []model.Review{
		{
			ID:        1,
			Content:   "test1",
			Rating:    5,
			CreatedAt: time.Now().Add(time.Hour * -24 * 3),
			EndedAt:   endedAt,
			Customer:  &model.Customer{ID: 1, Username: "test1", Avatar: "test1.jpg"},
			Freelance: &model.ReviewFreelance{Price: 1},
		},
		{
			ID:        2,
			Content:   "test2",
			Rating:    5,
			CreatedAt: time.Now().Add(time.Hour * -24 * 2),
			EndedAt:   endedAt,
			Customer:  &model.Customer{ID: 2, Username: "test2", Avatar: "test2.jpg"},
			Freelance: &model.ReviewFreelance{Price: 1},
		},
		{
			ID:        lastID,
			Content:   "test3",
			Rating:    5,
			CreatedAt: time.Now().Add(time.Hour * -24),
			EndedAt:   endedAt,
			Customer:  &model.Customer{ID: 3, Username: "test3", Avatar: "test3.jpg"},
			Freelance: &model.ReviewFreelance{Price: 1},
		},
	}

	rows := sqlmock.
		NewRows([]string{"id", "content", "rating", "created_at", "ended_at", "customer", "service"})
	for _, review := range expectedData {
		rows.AddRow(
			review.ID,
			review.Content,
			review.Rating,
			review.CreatedAt,
			review.EndedAt,
			fmt.Sprintf(`{"id": %d, "username": "%s", "avatar": "%s"}`, review.Customer.ID, review.Customer.Username, review.Customer.Avatar),
			fmt.Sprintf(`{"price": %f}`, review.Freelance.Price),
		)
	}

	mock.ExpectQuery(regexp.QuoteMeta(".")).WillReturnRows(rows)

	freelanceRepo := NewFreelanceRepository(sqlxDB)
	actualData, err := freelanceRepo.GetFreelanceServiceByIdReviews(id, cursorData, lastID, maxReviews+1)

	assert.NoError(t, err, "Error was not expected")
	assert.Equal(t, expectedData, *actualData, "Fetched reviews do not match expected result")
}

func TestGetFreelanceServiceById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "Error was not expected")

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-12-21 00:00:00")
	if err != nil {
		t.Fatalf("Error parsing date: %v", err)
	}

	freelanceID := 1

	expectedFreelance := model.FreelanceByID{
		ID:           int64(freelanceID),
		CreatedAt:    createdAt,
		Description:  "Test description",
		Title:        "Test title",
		ReviewsCount: 1,
		Rating:       5.0,
		Images:       &[]string{"image1", "image2"},
		Category:     &model.Category{ID: 1, Name: "Test Category"},
		Packages:     &[]model.Package{{ID: 1, DeliveryDays: 1, Description: "Test description", Price: 1, Title: "Test title"}},
		Freelancer:   &model.FreelanceServiceFreelancer{ID: 1, Username: "Test", FirstName: "Test", Surname: "Test", Avatar: "avatar1", Rating: 5.0, Level: 1, ReviewsCount: 1},
	}

	rows := sqlmock.
		NewRows([]string{"id", "created_at", "description", "title", "reviews_count", "rating", "images", "category", "packages", "freelancer"}).
		AddRow(freelanceID, createdAt, "Test description", "Test title", 1, 5.0, `["image1", "image2"]`, `{"id": 1, "name": "Test Category"}`, `[{"id": 1, "delivery_days": 1, "description": "Test description", "price": 1, "title": "Test title"}]`, `{"id": 1, "username": "Test", "first_name": "Test", "surname": "Test", "avatar": "avatar1", "rating": 5, "level": 1, "reviews_count": 1}`)

	mock.ExpectQuery(regexp.QuoteMeta(".")).WillReturnRows(rows)

	freelanceRepo := NewFreelanceRepository(sqlxDB)
	actualFreelance, err := freelanceRepo.GetFreelanceServiceById(freelanceID)
	assert.NoError(t, err, "Error was not expected")

	assert.Equal(t, expectedFreelance, *actualFreelance, "Fetched freelance service does not match expected result")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Error was not expected")
}

func TestGetResrictedFreelanceById(t *testing.T) {
	testCases := []struct {
		name      string
		mockId    int
		mockError error
		mockData  *model.FreelanceByIdRestricted
	}{
		{
			name:      "Success",
			mockId:    1,
			mockError: nil,
			mockData:  ptr(createRestrictedFreelanceByIdModel(1)),
		},
		{
			name:      "Error Not Found",
			mockId:    1,
			mockError: main_utils.ErrNoFound,
			mockData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err, "Error was not expected")
			sqlxDb := sqlx.NewDb(db, "sqlmock")

			defer func() {
				db.Close()
				sqlxDb.Close()
			}()

			rows := mock.NewRows([]string{"id", "title", "reviews_count", "rating", "image", "packages"})
			addRestrictedFreelanceByIdRow(rows, tc.mockId)

			mock.ExpectQuery(regexp.QuoteMeta(utils.RestrictedFreelanceQuery)).
				WillReturnRows(rows).
				WillReturnError(tc.mockError)

			freelanceRepo := NewFreelanceRepository(sqlxDb)
			actualData, err := freelanceRepo.GetResrictedFreelanceById(tc.mockId)

			assert.Equal(t, tc.mockError, err, "Error should be nil")
			assert.Equal(t, tc.mockData, actualData, "Fetched data does not match expected result")
		})
	}
}

func addRestrictedFreelanceByIdRow(rows *sqlmock.Rows, id int) {
	freelanceModel := createRestrictedFreelanceByIdModel(id)
	packagesJSON, _ := json.Marshal(freelanceModel.Packages)
	rows.AddRow(
		freelanceModel.Id, freelanceModel.Title, freelanceModel.ReviewsCount,
		freelanceModel.ReviewsCount, freelanceModel.Image, packagesJSON,
	)
}

func createRestrictedFreelanceByIdModel(id int) model.FreelanceByIdRestricted {
	return model.FreelanceByIdRestricted{
		Id:           int64(id),
		ReviewsCount: 0,
		Rating:       0,
		Title:        "test",
		Image:        ptr("test.jpg"),
		Packages:     &[]model.Package{{ID: 1, DeliveryDays: 0, Description: "test", Price: 0, Title: "test"}},
	}
}

func ptr[T any](value T) *T {
	return &value
}
