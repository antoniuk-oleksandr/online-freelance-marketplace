package controller

import (
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetReviewsByFreelanceID(id int, reviewsCursor string) (*dto.FreelanceReviews, error) {
	args := m.Called(id, reviewsCursor)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.FreelanceReviews), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockService) GetFreelanceById(id int) (*dto.Freelance, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.Freelance), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetFreelanceById_Success(t *testing.T) {
	mockService := new(MockService)
	mockResponse := &dto.Freelance{
		ID:           1,
		CreatedAt:    time.Now(),
		Description:  "test",
		ReviewsCount: 1,
		Rating:       5,
		Title:        "test",
		Images:       &[]string{"test1.jpg", "test2.jpg"},
		Category:     &model.Category{ID: 1, Name: "test"},
		Packages:     &[]dto.Package{{ID: 1, DeliveryDays: 1, Description: "test", Price: 1, Title: "test"}},
		Freelancer:   &dto.FreelanceServiceFreelancer{ID: 1, Username: "test", FirstName: "test", Surname: "test", Avatar: "test", Rating: 5, Level: 1, ReviewsCount: 1},
		Reviews:      &[]model.Review{{ID: 1, Content: "test", Rating: 5, CreatedAt: time.Now(), EndedAt: time.Now(), Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}}},
	}

	mockService.On("GetFreelanceById", 1).Return(mockResponse, nil)

	fc := NewFreelanceController(mockService)

	app := fiber.New()

	app.Get("/freelance/:id", fc.GetFreelanceById)

	req, _ := http.NewRequest("GET", "/freelance/1", nil)

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestGetFreelanceById_ServiceError(t *testing.T) {
	mockService := new(MockService)
	mockService.On("GetFreelanceById", 1).Return(nil, utils.ErrNotFound)

	fc := NewFreelanceController(mockService)

	app := fiber.New()

	app.Get("/freelances/:id", fc.GetFreelanceById)

	req, err := http.NewRequest("GET", "/freelances/1", nil)

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, "Status code should be 404")

	mockService.AssertExpectations(t)
}

func TestGetReviewsByFreelanceID_Success(t *testing.T) {
	mockService := new(MockService)
	mockResponse := &dto.FreelanceReviews{
		Reviews:        &[]model.Review{{ID: 1, Content: "test", Rating: 5, CreatedAt: time.Now(), EndedAt: time.Now(), Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}}},
		HasMoreReviews: false,
	}
	mockService.On("GetReviewsByFreelanceID", 1, "").Return(mockResponse, nil)

	fc := NewFreelanceController(mockService)

	app := fiber.New()
	app.Get("/freelances/:id/reviews", fc.GetReviewsByFreelanceID)

	req, err := http.NewRequest("GET", "/freelances/1/reviews", nil)
	assert.NoError(t, err, "Error creating request")

	resp, err := app.Test(req, -1)

	assert.NoError(t, err, "Error testing request")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")

	mockService.AssertExpectations(t)
}

func TestGetReviewsByFreelanceID_Error(t *testing.T) {
	mockService := new(MockService)
	mockService.On("GetReviewsByFreelanceID", 1, "").Return(nil, utils.ErrNotFound)

	fc := NewFreelanceController(mockService)

	app := fiber.New()
	app.Get("/freelances/:id/reviews", fc.GetReviewsByFreelanceID)

	req, err := http.NewRequest("GET", "/freelances/1/reviews", nil)
	assert.NoError(t, err, "Error creating request")

	resp, err := app.Test(req, -1)

	assert.NoError(t, err, "Error testing request")
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, "Status code should be 404")

	mockService.AssertExpectations(t)
}
