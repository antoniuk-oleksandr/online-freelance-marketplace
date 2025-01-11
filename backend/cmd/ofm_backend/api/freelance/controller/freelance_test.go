package controller

import (
	"fmt"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	test_utils "ofm_backend/cmd/ofm_backend/test_utils"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetResrictedFreelanceById(id int) (*dto.FreelanceByIdRestricted, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.FreelanceByIdRestricted), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockService) GetReviewsByFreelanceID(id int, reviewsCursor string) (*dto.FreelanceReviewsResponse, error) {
	args := m.Called(id, reviewsCursor)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.FreelanceReviewsResponse), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockService) GetFreelanceById(id int) (*dto.FreelanceByIDResponse, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.FreelanceByIDResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetFreelanceById_Success(t *testing.T) {
	mockService := new(MockService)
	mockResponse := &dto.FreelanceByIDResponse{
		Service: &dto.Freelance{
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
		},
		HasMoreReviews: false,
		ReviewsCursor:  nil,
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
	mockResponse := &dto.FreelanceReviewsResponse{
		Reviews:        &[]model.Review{{ID: 1, Content: "test", Rating: 5, CreatedAt: time.Now(), EndedAt: time.Now(), Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}}},
		HasMoreReviews: false,
	}
	mockService.On("GetReviewsByFreelanceID", 1, "").Return(mockResponse, nil)

	fc := NewFreelanceController(mockService)

	app := fiber.New()
	app.Get("/freelances/:id/reviews", fc.GetReviewsByFreelanceId)

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
	app.Get("/freelances/:id/reviews", fc.GetReviewsByFreelanceId)

	req, err := http.NewRequest("GET", "/freelances/1/reviews", nil)
	assert.NoError(t, err, "Error creating request")

	resp, err := app.Test(req, -1)

	assert.NoError(t, err, "Error testing request")
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, "Status code should be 404")

	mockService.AssertExpectations(t)
}

func TestGetResrictedFreelanceById(t *testing.T) {
	testCases := []struct {
		name               string
		mockResponse       *dto.FreelanceByIdRestricted
		mockId             int
		requestid          int
		mockError          error
		expectedBody       interface{}
		expectedStatusCode int
	}{
		{
			name:               "Success",
			mockResponse:       ptr(createResrictedFreelanceByIdDto(1)),
			mockId:             1,
			requestid:          1,
			mockError:          nil,
			expectedBody:       createResrictedFreelanceByIdDto(1),
			expectedStatusCode: fiber.StatusOK,
		},
		{
			name:         "Error not found",
			mockResponse: nil,
			mockId:       1,
			requestid:    1,
			mockError:    utils.ErrNotFound,
			expectedBody: utils.ErrorResponse{
				Error: utils.ErrNotFound.Error(),
			},
			expectedStatusCode: fiber.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockService)

			if tc.mockError != nil {
				mockService.On("GetResrictedFreelanceById", tc.mockId).Return(nil, tc.mockError)
			} else {
				mockService.On("GetResrictedFreelanceById", 1).Return(tc.mockResponse, tc.mockError)
			}

			fc := NewFreelanceController(mockService)

			app := test_utils.SetupFiberApp("/freelances/:id/restricted", fc.GetResrictedFreelanceById)
			resp := test_utils.PerformRequest(t, app, "GET", fmt.Sprintf("/freelances/%d/restricted", tc.requestid))

			var actualBody interface{}
			if tc.mockError != nil {
				actualBody = test_utils.ParseResponseBody[utils.ErrorResponse](t, resp)
			} else {
				actualBody = test_utils.ParseResponseBody[dto.FreelanceByIdRestricted](t, resp)
			}

			assert.Equal(t, tc.expectedBody, actualBody, "Response body should be equal")
			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode, "Status code should be equal")

			mockService.AssertExpectations(t)
		})
	}
}

func createResrictedFreelanceByIdDto(id int) dto.FreelanceByIdRestricted {
	return dto.FreelanceByIdRestricted{
		Id:           int64(id),
		ReviewsCount: 0,
		Rating:       0,
		Title:        "test",
		Image:        ptr("test.jpg"),
		Packages:     &[]dto.Package{{ID: 1, DeliveryDays: 0, Description: "test", Price: 0, Title: "test"}},
	}
}

func ptr[T any](data T) *T {
	return &data
}
