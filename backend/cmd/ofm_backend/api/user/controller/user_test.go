package controller

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"ofm_backend/cmd/ofm_backend/api/user/utils"
	"ofm_backend/cmd/ofm_backend/test_utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetReviewsByUserId(id int, cursor string) (*dto.ReviewsResponse, error) {
	args := m.Called(id, cursor)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.ReviewsResponse), nil
	}

	return nil, args.Error(1)
}

func (m *MockService) GetServicesByUserId(id int, cursor string) (*dto.ServicesResponse, error) {
	args := m.Called(id, cursor)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.ServicesResponse), nil
	}

	return nil, args.Error(1)
}

func (m *MockService) GetUserById(id int) (*dto.UserByIDResponse, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.UserByIDResponse), nil
	}

	return nil, args.Error(1)
}

func TestGetReviewsByUserId_WithoutCursor(t *testing.T) {
	mockService := new(MockService)

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	reviews := &[]dto.UserByIdReviewDto{
		{ID: int64(1), Rating: 5, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(1), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(1), Price: 1, Title: "test"}},
		{ID: int64(2), Rating: 3, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(2), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(2), Price: 1, Title: "test"}},
	}

	mockResponse := &dto.ReviewsResponse{
		HasMoreReviews: false,
		ReviewsCursor:  nil,
		Reviews:        reviews,
	}

	mockService.On("GetReviewsByUserId", 1, "").Return(mockResponse)

	uc := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/reviews", uc.GetReviewsByUserId)
	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1/reviews")
	respBody := test_utils.ParseResponseBody[*dto.ReviewsResponse](t, resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")
	assert.Equal(t, mockResponse, respBody, "Response body should be equal")

	mockService.AssertExpectations(t)
}

func TestGetReviewsByUserId_WithCursor(t *testing.T) {
	mockService := new(MockService)
	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	endedAt := timeNow.Add(time.Hour * 24)
	endedAtStr := main_utils.ConvertTimeToSting(endedAt)

	cursor := utils.BuildReviewsCursor(endedAtStr, int64(1))

	reviews := &[]dto.UserByIdReviewDto{
		{ID: int64(1), Rating: 5, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(1), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(1), Price: 1, Title: "test"}},
		{ID: int64(2), Rating: 3, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(2), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(2), Price: 1, Title: "test"}},
	}

	mockResponse := &dto.ReviewsResponse{
		HasMoreReviews: false,
		ReviewsCursor:  nil,
		Reviews:        reviews,
	}

	mockService.On("GetReviewsByUserId", 1, cursor).Return(mockResponse)

	uc := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/reviews", uc.GetReviewsByUserId)
	resp := test_utils.PerformRequest(t, app, "GET", fmt.Sprintf("/api/v1/users/1/reviews?cursor=%s", cursor))
	respBody := test_utils.ParseResponseBody[*dto.ReviewsResponse](t, resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")
	assert.Equal(t, mockResponse, respBody, "Response body should be equal")

	mockService.AssertExpectations(t)
}

func TestGetReviewsByUserId_Error(t *testing.T) {
	mockService := new(MockService)
	mockService.On("GetReviewsByUserId", 1, "").Return(nil, main_utils.ErrUnexpectedError)

	us := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/reviews", us.GetReviewsByUserId)

	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1/reviews")
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetServicesByUserId_WithoutCursor(t *testing.T) {
	mockService := new(MockService)

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	services := &[]model.UserByIdFreelanceService{
		{ID: 1, CreatedAt: timeNow, Description: "Service description 1", Title: "Service title 1", CategoryId: 101, FreelancerId: 1001, ReviewsCount: 25, Rating: 4.8, MinPrice: 50.0},
		{ID: 2, CreatedAt: timeNow, Description: "Service description 2", Title: "Service title 2", CategoryId: 102, FreelancerId: 1002, ReviewsCount: 15, Rating: 4.5, MinPrice: 75.0},
	}

	mockResponse := &dto.ServicesResponse{
		Services:        services,
		HasMoreServices: false,
		ServicesCursor:  nil,
	}

	mockService.On("GetServicesByUserId", 1, "").Return(mockResponse)

	uc := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/services", uc.GetServicesByUserId)
	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1/services")
	respBody := test_utils.ParseResponseBody[*dto.ServicesResponse](t, resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")
	assert.Equal(t, mockResponse, respBody, "Response body should be equal")

	mockService.AssertExpectations(t)
}

func TestGetServicesByUserId_WithCursor(t *testing.T) {
	mockService := new(MockService)

	cursor := utils.BuildServicesCursor(0, 1)

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	services := &[]model.UserByIdFreelanceService{
		{ID: 1, CreatedAt: timeNow, Description: "Service description 1", Title: "Service title 1", CategoryId: 101, FreelancerId: 1001, ReviewsCount: 25, Rating: 4.8, MinPrice: 50.0},
		{ID: 2, CreatedAt: timeNow, Description: "Service description 2", Title: "Service title 2", CategoryId: 102, FreelancerId: 1002, ReviewsCount: 15, Rating: 4.5, MinPrice: 75.0},
	}
	mockResponse := &dto.ServicesResponse{
		Services:        services,
		HasMoreServices: false,
		ServicesCursor:  nil,
	}

	mockService.On("GetServicesByUserId", 1, cursor).Return(mockResponse)

	uc := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/services", uc.GetServicesByUserId)
	resp := test_utils.PerformRequest(t, app, "GET", fmt.Sprintf("/api/v1/users/1/services?cursor=%s", cursor))
	respBody := test_utils.ParseResponseBody[*dto.ServicesResponse](t, resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")
	assert.Equal(t, mockResponse, respBody, "Response body should be equal")

	mockService.AssertExpectations(t)
}

func TestGetServicesByUserId_Error(t *testing.T) {
	mockService := new(MockService)
	mockService.On("GetServicesByUserId", 1, "").Return(nil, main_utils.ErrUnexpectedError)

	us := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id/services", us.GetServicesByUserId)

	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1/services")
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	
	mockService.AssertExpectations(t)
}

func TestGetUserById_Success(t *testing.T) {
	mockService := new(MockService)

	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	services := &[]model.UserByIdFreelanceService{
		{ID: 1, CreatedAt: timeNow, Description: "Service description 1", Title: "Service title 1", CategoryId: 101, FreelancerId: 1001, ReviewsCount: 25, Rating: 4.8, MinPrice: 50.0},
		{ID: 2, CreatedAt: timeNow, Description: "Service description 2", Title: "Service title 2", CategoryId: 102, FreelancerId: 1002, ReviewsCount: 15, Rating: 4.5, MinPrice: 75.0},
	}
	reviews := &[]dto.UserByIdReviewDto{
		{ID: int64(1), Rating: 5, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(1), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(1), Price: 1, Title: "test"}},
		{ID: int64(2), Rating: 3, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(2), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(2), Price: 1, Title: "test"}},
	}
	mockResponse := &dto.UserByIDResponse{
		User: &dto.UserByIdTO{
			ID:           1,
			About:        nil,
			CreatedAt:    timeNow,
			FirstName:    "John",
			Level:        4.5,
			Surname:      "Doe",
			Username:     "johndoe",
			Avatar:       nil,
			Rating:       nil,
			ReviewsCount: nil,
			Skills:       &[]model.Skill{{ID: 1, Name: "test1"}, {ID: 2, Name: "test2"}},
			Languages:    &[]model.Language{{ID: 1, Name: "test1"}, {ID: 2, Name: "test2"}},
			Services:     *services,
			Reviews:      *reviews,
		},
	}

	mockService.On("GetUserById", 1).Return(mockResponse)

	uc := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id", uc.GetUserById)
	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1")
	respBody := test_utils.ParseResponseBody[*dto.UserByIDResponse](t, resp)
	
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code should be 200")
	assert.Equal(t, mockResponse, respBody, "Response body should be equal")
	
	mockService.AssertExpectations(t)
}

func TestGetUserById_Error(t *testing.T) {
	mockService := new(MockService)
	mockService.On("GetUserById", 1).Return(nil, main_utils.ErrUnexpectedError)

	us := NewUserController(mockService)
	app := test_utils.SetupFiberApp("/api/v1/users/:id", us.GetUserById)

	resp := test_utils.PerformRequest(t, app, "GET", "/api/v1/users/1")
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	
	mockService.AssertExpectations(t)
}
