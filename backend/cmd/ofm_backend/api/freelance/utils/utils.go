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

func GetDataFromReviewsCursor(reviewsCursor string) (string, int64, error) {
	if reviewsCursor == "" {
		return "", -1, nil
	}

	decodedCursor, err := middleware.DecodeString(reviewsCursor)
	if err != nil {
		return "", -1, err
	}

	arrOfKeyValues := strings.Split(decodedCursor, ";")

	colonIndex := strings.Index(arrOfKeyValues[0], ":")
	if colonIndex == -1 {
		return "", -1, main_utils.ErrParsingError
	}

	timestamp := arrOfKeyValues[0][colonIndex+1:]
	lastIDStr := strings.Split(arrOfKeyValues[1], ":")[1]
	lastID, err := strconv.ParseInt(lastIDStr, 10, 64)

	return timestamp, lastID, err
}

func BuildReviewsCursor(lastReviewEndedAt time.Time, lastID int64) *string {
	lastReviewEndedAtStr := main_utils.ConvertTimeToSting(lastReviewEndedAt)
	str := fmt.Sprintf("reviewsCursor:%s;lastID:%d", lastReviewEndedAtStr, lastID)
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
