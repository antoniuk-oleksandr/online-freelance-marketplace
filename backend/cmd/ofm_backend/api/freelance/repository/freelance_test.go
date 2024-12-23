package repository

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetFreelanceServiceById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	rows := sqlmock.NewRows([]string{"id", "created_at", "description", "title", "reviews_count", "rating", "images", "category", "packages", "freelancer"}).
		AddRow(freelanceID, createdAt, "Test description", "Test title", 1, 5.0, `["image1", "image2"]`, `{"id": 1, "name": "Test Category"}`, `[{"id": 1, "delivery_days": 1, "description": "Test description", "price": 1, "title": "Test title"}]`, `{"id": 1, "username": "Test", "first_name": "Test", "surname": "Test", "avatar": "avatar1", "rating": 5, "level": 1, "reviews_count": 1}`)

	mock.ExpectQuery(regexp.QuoteMeta(".")).WillReturnRows(rows)

	actualFreelance, err := GetFreelanceServiceById(freelanceID, sqlxDB)
	assert.NoError(t, err, "Error was not expected")
	
	assert.Equal(t, expectedFreelance, *actualFreelance, "Fetched freelance service does not match expected result")
	
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Error was not expected")
	
}
