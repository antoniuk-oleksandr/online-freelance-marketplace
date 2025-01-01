package service

import (
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"ofm_backend/cmd/ofm_backend/api/user/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserById(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.User), nil
	}

	return nil, args.Error(1)
}
func (m *MockRepository) GetReviewsByUserId(
	id int, endedAt *string, lastId int64, maxReviews int,
) (*[]model.UserByIdReview, error) {
	args := m.Called(id, endedAt, lastId, maxReviews)
	if args.Get(0) != nil {
		return args.Get(0).(*[]model.UserByIdReview), nil
	}

	return nil, args.Error(1)
}
func (m *MockRepository) GetServicesByUserId(
	id int, reviewsCount, lastId int64, maxServices int,
) (*[]model.UserByIdFreelanceService, error) {
	args := m.Called(id, reviewsCount, lastId, maxServices)
	if args.Get(0) != nil {
		return args.Get(0).(*[]model.UserByIdFreelanceService), nil
	}

	return nil, args.Error(1)
}

func TestGetUserById(t *testing.T) {
	maxValue := 2
	os.Setenv("MAX_USER_BY_ID_REVIEWS", strconv.Itoa(maxValue))
	os.Setenv("MAX_USER_BY_ID_SERVICES", strconv.Itoa(maxValue))

	defer func() {
		os.Unsetenv("MAX_USER_BY_ID_REVIEWS")
		os.Unsetenv("MAX_USER_BY_ID_SERVICES")
	}()

	now := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name                 string
		userId               int
		reviewsCount, lastId int64
		maxValue             int
		mockUser             *model.User
		mockReviews          *[]model.UserByIdReview
		mockServices         *[]model.UserByIdFreelanceService
		expectedResult       *dto.UserByIDResponse
		expectedError        error
	}{
		{
			name:           "Successful Fetch Without Cursor",
			userId:         1,
			reviewsCount:   int64(-1),
			lastId:         int64(-1),
			maxValue:       maxValue + 1,
			mockUser:       createUserModel(now),
			mockReviews:    createReviewsModels(now, maxValue),
			mockServices:   createServicesModels(maxValue),
			expectedResult: createExpectedUserDTO(now, maxValue, false),
			expectedError:  nil,
		},
		{
			name:           "User Not Found",
			userId:         2,
			reviewsCount:   int64(-1),
			lastId:         int64(-1),
			maxValue:       maxValue + 1,
			mockUser:       nil,
			mockReviews:    nil,
			mockServices:   nil,
			expectedResult: nil,
			expectedError:  main_utils.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)

			if tc.mockUser != nil {
				mockRepo.On("GetUserById", tc.userId).Return(tc.mockUser, nil)
				var val *string
				mockRepo.On("GetReviewsByUserId", tc.userId, val, tc.lastId, tc.maxValue).Return(tc.mockReviews, nil)
				mockRepo.On("GetServicesByUserId", tc.userId, tc.reviewsCount, tc.lastId, tc.maxValue).Return(tc.mockServices, nil)
			} else {
				mockRepo.On("GetUserById", tc.userId).Return(nil, main_utils.ErrUserNotFound)
			}

			us := NewUserService(mockRepo)
			result, err := us.GetUserById(tc.userId)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetReviewsByUserId(t *testing.T) {
	maxValue := 2
	os.Setenv("MAX_USER_BY_ID_REVIEWS", strconv.Itoa(maxValue))
	defer os.Unsetenv("MAX_USER_BY_ID_REVIEWS")

	mockRepository := new(MockRepository)
	timeNow := time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)
	timeNowStr := main_utils.ConvertTimeToSting(timeNow)
	reviewsModels := createReviewsModels(timeNow, maxValue)
	reviewsDtos := createReviewsDTOs(timeNow, maxValue)

	tests := []struct {
		name           string
		cursor         *string
		offset         int64
		expectedCursor *string
	}{
		{
			name:           "WithoutCursor",
			cursor:         nil,
			offset:         -1,
			expectedCursor: nil,
		},
		{
			name:           "WithCursor",
			cursor:         &timeNowStr,
			offset:         1,
			expectedCursor: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository.ExpectedCalls = nil
			mockRepository.On("GetReviewsByUserId", 1, tt.cursor, tt.offset, utils.GetMaxReviews()+1).Return(reviewsModels)

			expectedData := &dto.ReviewsResponse{
				Reviews:        reviewsDtos,
				HasMoreReviews: false,
				ReviewsCursor:  tt.expectedCursor,
			}

			us := NewUserService(mockRepository)
			cursor := ""
			if tt.cursor != nil {
				cursor = utils.BuildReviewsCursor(*tt.cursor, 1)
			}
			actualData, err := us.GetReviewsByUserId(1, cursor)
			assert.NoError(t, err, "Error should be nil")
			assert.Equal(t, expectedData, actualData, "Data should be equal")

			mockRepository.AssertExpectations(t)
		})
	}
}

func TestGetServicesByUserId(t *testing.T) {
	maxValue := 2
	os.Setenv("MAX_USER_BY_ID_SERVICES", strconv.Itoa(maxValue))
	defer os.Unsetenv("MAX_USER_BY_ID_SERVICES")

	mockRepository := new(MockRepository)
	serviceModels := createServicesModels(maxValue)
	serviceDtos := createServicesDtos(maxValue)

	tests := []struct {
		name           string
		cursor         string
		expectedOffset int64
		expectedPage   int64
	}{
		{
			name:           "WithoutCursor",
			cursor:         "",
			expectedOffset: -1,
			expectedPage:   -1,
		},
		{
			name:           "WithCursor",
			cursor:         utils.BuildServicesCursor(int64(1), int64(1)),
			expectedOffset: 1,
			expectedPage:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository.ExpectedCalls = nil
			mockRepository.On("GetServicesByUserId", 1, tt.expectedOffset, tt.expectedPage, utils.GetMaxServices()+1).Return(serviceModels)

			expectedData := &dto.ServicesResponse{
				Services:        serviceDtos,
				HasMoreServices: false,
				ServicesCursor:  nil,
			}

			us := NewUserService(mockRepository)
			actualData, err := us.GetServicesByUserId(1, tt.cursor)
			assert.NoError(t, err, "Error should be nil")
			assert.Equal(t, expectedData, actualData, "Data should be equal")

			mockRepository.AssertExpectations(t)
		})
	}
}

func createExpectedUserDTO(timeNow time.Time, num int, addCurors bool) *dto.UserByIDResponse {
	var hasMoreReviews bool
	var reviewsCursor *string
	var hasMoreServices bool
	var servicesCursor *string

	if addCurors {
		hasMoreReviews = true
		hasMoreServices = true

		revCursor := utils.BuildReviewsCursor(main_utils.ConvertTimeToSting(timeNow), int64(num-1))
		reviewsCursor = &revCursor

		servCursor := utils.BuildServicesCursor(int64(0), int64(num-1))
		servicesCursor = &servCursor
	}

	return &dto.UserByIDResponse{
		User:            createUserDTO(timeNow, num),
		HasMoreReviews:  hasMoreReviews,
		HasMoreServices: hasMoreServices,
		ReviewsCursor:   reviewsCursor,
		ServicesCursor:  servicesCursor,
	}
}

func createUserModel(timeNow time.Time) *model.User {
	return &model.User{
		ID:        1,
		CreatedAt: timeNow,
		Email:     "test",
		FirstName: "test",
		Level:     float64(0.0),
		Surname:   "test",
		Username:  "test",
		Avatar:    nil,
		RoleID:    2,
		Languages: &[]model.Language{{ID: 0, Name: "test"}, {ID: 0, Name: "test"}},
		Skills:    &[]model.Skill{{ID: int64(0), Name: "test"}, {ID: int64(0), Name: "test"}},
	}
}

func createUserDTO(timeNow time.Time, num int) *dto.UserByIdTO {
	rating := float64(0.0)
	reviewsCount := int64(0)

	return &dto.UserByIdTO{
		ID:           int64(1),
		CreatedAt:    timeNow,
		FirstName:    "test",
		Level:        float64(0.0),
		Surname:      "test",
		Username:     "test",
		Avatar:       nil,
		Rating:       &rating,
		ReviewsCount: &reviewsCount,
		Languages:    &[]model.Language{{ID: int(0), Name: "test"}, {ID: int(0), Name: "test"}},
		Skills:       &[]model.Skill{{ID: int64(0), Name: "test"}, {ID: int64(0), Name: "test"}},
		Reviews:      *createReviewsDTOs(timeNow, num),
		Services:     *createServicesDtos(num),
	}
}

func createReviewsModels(timeNow time.Time, num int) *[]model.UserByIdReview {
	arr := make([]model.UserByIdReview, num)

	for i := 0; i < len(arr); i++ {
		arr[i] = model.UserByIdReview{ID: int64(i + 1), Content: "test", Rating: 0, CreatedAt: timeNow, EndedAt: timeNow, UserID: int64(0), Username: "test", Avatar: nil, ServiceID: int64(0), Price: 0.0, ServiceImage: nil, Title: "test"}
	}

	return &arr
}

func createReviewsDTOs(timeNow time.Time, num int) *[]dto.UserByIdReviewDto {
	arr := make([]dto.UserByIdReviewDto, num)

	for i := 0; i < len(arr); i++ {
		arr[i] = dto.UserByIdReviewDto{ID: int64(i + 1), Rating: 0, Content: "test", CreatedAt: timeNow, EndedAt: timeNow, Customer: dto.ReviewUserDTO{ID: int64(0), Username: "test"}, ReviewService: dto.ReviewService{ID: int64(0), Price: float64(0.0), Title: "test"}}
	}

	return &arr
}

func createServicesModels(num int) *[]model.UserByIdFreelanceService {
	arr := make([]model.UserByIdFreelanceService, num)

	for i := 0; i < len(arr); i++ {
		arr[i] = model.UserByIdFreelanceService{ID: int64(i + 1), Title: "test", ReviewsCount: int64(0), Rating: float64(0.0), MinPrice: 0.0}
	}

	return &arr
}

func createServicesDtos(num int) *[]dto.ServiceByIdDto {
	arr := make([]dto.ServiceByIdDto, num)

	for i := 0; i < len(arr); i++ {
		arr[i] = dto.ServiceByIdDto{ID: int64(i + 1), Title: "test", ReviewsCount: int64(0), Rating: float64(0.0), MinPrice: 0.0}
	}

	return &arr
}