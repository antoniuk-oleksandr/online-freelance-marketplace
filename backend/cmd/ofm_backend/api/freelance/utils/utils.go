package utils

import (
	"fmt"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/middleware"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetDataFromReviewsCursor(reviewsCursor string) (string, error) {
	if reviewsCursor == "" {
		return "", nil
	}

	decodedCursor, err := middleware.DecodeString(reviewsCursor)
	if err != nil {
		return "", err
	}

	colonIndex := strings.Index(decodedCursor, ":")
	if colonIndex == -1 {
		return "", main_utils.ErrParsingError
	}

	timestamp := decodedCursor[colonIndex+1:]

	return timestamp, nil
}

func BuildReviewsCursor(lastReviewEndedAt time.Time) *string {
	lastReviewEndedAtStr := main_utils.ConvertTimeToSting(lastReviewEndedAt)
	str := fmt.Sprintf("reviewsCursor:%s", lastReviewEndedAtStr)
	encodedStr := middleware.EncodeString(str)

	return &encodedStr
}

func GetMaxReviewsValue() int {
	defaultValue := 5

	maxReviewsStr := os.Getenv("MAX_FREELANCE_BY_ID_REVIEWS")
	if maxReviewsStr == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(maxReviewsStr)
	if err != nil {
		return defaultValue
	}

	return val
}
