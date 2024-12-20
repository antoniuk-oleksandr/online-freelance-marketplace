package service

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/search/body"
	"ofm_backend/cmd/ofm_backend/api/search/dto"
	"ofm_backend/cmd/ofm_backend/api/search/enum"
	"ofm_backend/cmd/ofm_backend/api/search/mapper"
	"ofm_backend/cmd/ofm_backend/api/search/model"
	"ofm_backend/cmd/ofm_backend/api/search/repository"
	"ofm_backend/cmd/ofm_backend/api/search/utils"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/database"
	"ofm_backend/internal/middleware"
	"os"
	"strconv"
	"strings"
)

func SearchFreelances(searchBody body.Search) (*dto.SearchFreelances, error) {
	db := database.GetDB()
	
	maxResults, err := strconv.Atoi(os.Getenv("MAX_RESULTS"))
	if err != nil {
		return nil, err
	}

	cursorData, err := getDataFromCursor(searchBody)
	if err != nil {
		return nil, err
	}

	query := utils.BuildSearchServicesQuery(searchBody, cursorData, maxResults)

	freelanceModels, err := repository.SearchFreelances(db, query)
	if err != nil {
		return nil, err
	}

	freelanceDTOs := mapper.MapSearchFreelancesModelToDTO(freelanceModels)


	cursor, hasMore := buildCursor(searchBody, freelanceModels, maxResults)
	return &dto.SearchFreelances{
		Services: freelanceDTOs,
		HasMore:  hasMore,
		Cursor:   cursor,
	}, nil
}

func buildCursor(
	searchBody body.Search,
	freelanceModels []model.SearchService,
	maxResults int,
) (*string, bool) {
	if len(freelanceModels) <= maxResults {
		return nil, false
	}

	lastModel := freelanceModels[maxResults-1]
	value := getCursorValue(searchBody.Sort, lastModel)
	cursor := fmt.Sprintf("lastID:%d;value:%s", lastModel.ID, value)
	encodedCursor := middleware.EncodeString(cursor)

	return &encodedCursor, true
}

func getCursorValue(sort *int, lastModel model.SearchService) string {
	if sort == nil {
		return strconv.FormatInt(lastModel.LastMonthCompletedOrdersCount, 10)
	}

	switch *sort {
	case enum.Name:
		return lastModel.Title
	case enum.Popularity:
		return strconv.FormatInt(lastModel.LastMonthCompletedOrdersCount, 10)
	case enum.Rating:
		return strconv.FormatFloat(lastModel.Rating, 'f', -1, 64)
	case enum.Price:
		return strconv.FormatFloat(lastModel.MinPrice, 'f', -1, 64)
	case enum.Level:
		return strconv.FormatFloat(lastModel.Level, 'f', -1, 64)
	default:
		return ""
	}
}

func getDataFromCursor(searchBody body.Search) (*[]string, error) {
	if searchBody.Cursor == nil {
		return nil, nil
	}

	decodedCursor, err := middleware.DecodeString(*searchBody.Cursor)
	if err != nil {
		return nil, err
	}

	keyArr := strings.Split(decodedCursor, ";")

	if len(keyArr) != 2 {
		return nil, main_utils.ErrInvalidCursor
	}
	
	valueArr := make([]string, 2)
	valueArr[0] = strings.Split(keyArr[0], ":")[1]
	valueArr[1] = strings.Split(keyArr[1], ":")[1]
	
	return &valueArr, nil
}
