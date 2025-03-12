package service

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
	filter_params_dto "ofm_backend/cmd/ofm_backend/api/filter_params/dto"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"ofm_backend/cmd/ofm_backend/api/freelance/utils"
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

func (m *MockRepository) GetResrictedFreelanceById(id int) (*model.FreelanceByIdRestricted, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.FreelanceByIdRestricted), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepository) GetFreelanceServiceById(id int) (*model.FreelanceByID, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.FreelanceByID), nil
	}
	return nil, args.Error(1)
}

func (m *MockRepository) GetFreelanceServiceByIdReviews(id int, cursorData string, lastID int64, maxReviews int) (*[]model.Review, error) {
	args := m.Called(id, cursorData, lastID, maxReviews)
	if args.Get(0) != nil {
		return args.Get(0).(*[]model.Review), nil
	}
	return nil, args.Error(1)
}

func TestGetFreelanceById_Success(t *testing.T) {
	maxReviews := 2
	os.Setenv("FILE_SERVER_HOST", "localost")
	os.Setenv("FILE_SERVER_PORT", "8083")
	os.Setenv("MAX_FREELANCE_BY_ID_REVIEWS", strconv.Itoa(maxReviews))

	defer func() {
		os.Unsetenv("FILE_SERVER_HOST")
		os.Unsetenv("FILE_SERVER_PORT")
		os.Unsetenv("MAX_FREELANCE_BY_ID_REVIEWS")
	}()

	timeNow := time.Now()

	mockRepo := new(MockRepository)
	mockRepoFreelanceByIdResponse := &model.FreelanceByID{
		ID:           1,
		CreatedAt:    timeNow,
		Description:  "test",
		ReviewsCount: 1,
		Rating:       5,
		Title:        "test",
		Images:       &[]string{"test1.jpg", "test2.jpg"},
		Category:     &model.Category{ID: 1, Name: "test"},
		Packages:     &[]model.Package{{ID: 1, DeliveryDays: 1, Description: "test", Price: 1, Title: "test"}},
		Freelancer:   &model.FreelanceServiceFreelancer{ID: 1, Username: "test", FirstName: "test", Surname: "test", Avatar: "test", Rating: 5, Level: 1, ReviewsCount: 1},
	}
	mockRepoReviewsResponse := &[]model.Review{
		{ID: 1, Content: "test", Rating: 5, CreatedAt: timeNow, EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
		{ID: 2, Content: "test", Rating: 5, CreatedAt: timeNow, EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
	}

	extectedData := &dto.FreelanceByIDResponse{
		Service: &dto.Freelance{
			ID:           1,
			CreatedAt:    timeNow,
			Description:  "test",
			ReviewsCount: 1,
			Rating:       5,
			Title:        "test",
			Images:       &[]string{"http://localost:8083/files/test1.jpg", "http://localost:8083/files/test2.jpg"},
			Category:     &filter_params_dto.FilterItem{ID: 1, Name: "test"},
			Packages:     &[]dto.Package{{ID: 1, DeliveryDays: 1, Description: "test", Price: 1, Title: "test"}},
			Freelancer:   &dto.FreelanceServiceFreelancer{ID: 1, Username: "test", FirstName: "test", Surname: "test", Avatar: "test", Rating: 5, Level: 1, ReviewsCount: 1},
			Reviews:      mockRepoReviewsResponse,
		},
		HasMoreReviews: false,
		ReviewsCursor:  nil,
	}

	mockRepo.On("GetFreelanceServiceById", 1).Return(mockRepoFreelanceByIdResponse)
	mockRepo.On("GetFreelanceServiceByIdReviews", 1, "", int64(-1), maxReviews+1).Return(mockRepoReviewsResponse)

	fs := NewFreelanceService(mockRepo)
	actualData, err := fs.GetFreelanceById(1)

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, extectedData, actualData, "Data should be equal")
}

func TestGetReviewsByFreelanceID_WithoutCursor(t *testing.T) {
	mockRepo := new(MockRepository)

	maxReviews := utils.GetMaxReviewsValue()

	timeNow := time.Now()

	reviews := &[]model.Review{{ID: 1, Content: "test", Rating: 5, CreatedAt: timeNow, EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}}}

	mockRepoResponse := &dto.FreelanceReviewsResponse{
		Reviews:        reviews,
		HasMoreReviews: false,
		ReviewsCursor:  nil,
	}

	mockRepo.On("GetFreelanceServiceByIdReviews", 1, "", int64(-1), maxReviews+1).Return(reviews, nil)

	fs := NewFreelanceService(mockRepo)
	actualData, err := fs.GetReviewsByFreelanceID(1, "")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, mockRepoResponse, actualData, "Data should be equal")

	mockRepo.AssertExpectations(t)
}

func TestGetReviewsByFreelanceID_WithCursor(t *testing.T) {
	maxReviews := 2
	timeNow := time.Now()

	os.Setenv("FILE_SERVER_HOST", "localost")
	os.Setenv("FILE_SERVER_PORT", "8083")
	os.Setenv("MAX_FREELANCE_BY_ID_REVIEWS", strconv.Itoa(maxReviews))

	defer func() {
		os.Unsetenv("FILE_SERVER_HOST")
		os.Unsetenv("FILE_SERVER_PORT")
		os.Unsetenv("MAX_FREELANCE_BY_ID_REVIEWS")
	}()

	cursor := utils.BuildReviewsCursor(timeNow.Add(time.Hour*-24*2), -1)
	cursorData, lastID, err := utils.GetDataFromReviewsCursor(*cursor)
	assert.NoError(t, err, "Error should be nil")

	mockRepo := new(MockRepository)

	reviews := &[]model.Review{
		{ID: 1, Content: "test", Rating: 5, CreatedAt: timeNow.Add(time.Hour * -24), EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
		{ID: 2, Content: "test", Rating: 5, CreatedAt: timeNow.Add(time.Hour * -24 * 2), EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
		{ID: 3, Content: "test", Rating: 5, CreatedAt: timeNow.Add(time.Hour * -24 * 3), EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
		{ID: 4, Content: "test", Rating: 5, CreatedAt: timeNow.Add(time.Hour * -24 * 4), EndedAt: timeNow, Customer: &model.Customer{ID: 1, Username: "test", Avatar: "test"}, Freelance: &model.ReviewFreelance{Price: 1}},
	}

	mockRepoResponse := &dto.FreelanceReviewsResponse{
		Reviews:        reviews,
		HasMoreReviews: false,
		ReviewsCursor:  nil,
	}

	mockRepo.On("GetFreelanceServiceByIdReviews", 1, cursorData, lastID, maxReviews+1).Return(reviews, nil)

	fs := NewFreelanceService(mockRepo)
	actualData, err := fs.GetReviewsByFreelanceID(1, *cursor)

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, mockRepoResponse, actualData, "Data should be equal")

	mockRepo.AssertExpectations(t)
}

func TestGetResrictedFreelanceById(t *testing.T) {
	os.Setenv("FILE_SERVER_HOST", "localost")
	os.Setenv("FILE_SERVER_PORT", "8030")

	defer func() {
		os.Unsetenv("FILE_SERVER_HOST")
		os.Unsetenv("FILE_SERVER_PORT")
	}()

	testcases := []struct {
		Name         string
		MockId       int
		MockError    error
		MockData     *model.FreelanceByIdRestricted
		ExpectedData *dto.FreelanceByIdRestricted
	}{
		{
			Name:         "Success",
			MockId:       1,
			MockError:    nil,
			MockData:     ptr(createRestrictedFreelanceByIdModel(1)),
			ExpectedData: ptr(createResrictedFreelanceByIdDto(1)),
		},
		{
			Name:         "Error Not Found",
			MockId:       1,
			MockError:    main_utils.ErrNotFound,
			MockData:     nil,
			ExpectedData: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			mockRepo := new(MockRepository)

			fs := NewFreelanceService(mockRepo)
			mockRepo.On("GetResrictedFreelanceById", tc.MockId).Return(tc.MockData, tc.MockError)
			actualData, err := fs.GetResrictedFreelanceById(tc.MockId)

			assert.Equal(t, tc.ExpectedData, actualData, "Data should be equal")
			assert.Equal(t, tc.MockError, err, "Error should be equal")

			mockRepo.AssertExpectations(t)
		})
	}
}

func createResrictedFreelanceByIdDto(id int) dto.FreelanceByIdRestricted {
	return dto.FreelanceByIdRestricted{
		Id:           int64(id),
		ReviewsCount: 0,
		Rating:       0,
		Title:        "test",
		Image:        ptr("http://localost:8030/files/test.jpg"),
		Packages:     &[]dto.Package{{ID: 1, DeliveryDays: 0, Description: "test", Price: 0, Title: "test"}},
	}
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

func ptr[T any](data T) *T {
	return &data
}
