package utils

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/user/dto"
	"ofm_backend/cmd/ofm_backend/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"os"
	"regexp"
	"strconv"
)

const (
	DefaultMaxReviewsVal  int = 5
	DefaultMaxServicesVal int = 5
)

func GetMaxReviews() int {
	valStr := os.Getenv("MAX_USER_BY_ID_REVIEWS")

	if valStr == "" {
		return DefaultMaxReviewsVal
	}

	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return DefaultMaxReviewsVal
	}

	return valInt
}

func GetMaxServices() int {
	valStr := os.Getenv("MAX_USER_BY_ID_SERVICES")

	if valStr == "" {
		return DefaultMaxServicesVal
	}

	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return DefaultMaxServicesVal
	}

	return valInt
}

func BuildServicesCursor(reviewsCount, lastId int64) string {
	cursor := fmt.Sprintf("reviewsCount:%d;lastId:%d", reviewsCount, lastId)
	return utils.EncodeString(cursor)
}

func BuildReviewsCursor(endedAt string, lastId int64) string {
	cursor := fmt.Sprintf("endedAt:%s;lastId:%d", endedAt, lastId)
	return utils.EncodeString(cursor)
}

func ParseReviewsCursor(cursor string) (*string, int64, error) {
	if cursor == "" {
		return nil, int64(-1), nil
	}

	decodedCursor, err := main_utils.DecodeString(cursor)
	if err != nil {
		return nil, int64(-1), utils.ErrInvalidCursor
	}

	re := regexp.MustCompile(`endedAt:([^;]+);lastId:(\d+)`)
	matches := re.FindStringSubmatch(decodedCursor)
	if len(matches) != 3 {
		return nil, -1, utils.ErrInvalidCursor
	}

	endedAt := matches[1]
	lastId, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return nil, -1, utils.ErrInvalidCursor
	}

	return &endedAt, lastId, nil
}

func ParseServicesCursor(cursor string) (int64, int64, error) {
	if cursor == "" {
		return int64(-1), int64(-1), nil
	}

	decodedCursor, err := main_utils.DecodeString(cursor)
	if err != nil {
		return int64(-1), int64(-1), err
	}

	var reviewsCount int64
	var lastId int64
	_, err = fmt.Sscanf(decodedCursor, "reviewsCount:%d;lastId:%d", &reviewsCount, &lastId)
	if err != nil {
		return int64(-1), int64(-1), err
	}

	return reviewsCount, lastId, nil
}

func GetMoreServicesCursorData(
	services *[]dto.ServiceByIdDto, maxServices int,
) (bool, *string) {
	var hasMoreServices bool
	var servicesCursor *string
	if services != nil && len(*services) > maxServices {
		hasMoreServices = true
		lastService := (*services)[len(*services)-2]
		cursor := BuildServicesCursor(lastService.ReviewsCount, lastService.ID)
		servicesCursor = &cursor
		*services = (*services)[:len(*services)-1]
	}

	return hasMoreServices, servicesCursor
}

func GetMoreReviewsCursorData(
	reviews *[]dto.UserByIdReviewDto, maxReviews int,
) (bool, *string) {
	var hasMoreReviews bool
	var reviewsCursor *string
	if reviews != nil && len(*reviews) > maxReviews {
		hasMoreReviews = true
		lastReview := (*reviews)[len(*reviews)-2]
		lastReviewEndedAt := main_utils.ConvertTimeToSting(lastReview.EndedAt)
		cursor := BuildReviewsCursor(lastReviewEndedAt, lastReview.ID)
		reviewsCursor = &cursor
		*reviews = (*reviews)[:len(*reviews)-1]
	}

	return hasMoreReviews, reviewsCursor
}
