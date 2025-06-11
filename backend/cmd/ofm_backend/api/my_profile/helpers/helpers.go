package helpers

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func CalcMyProfileLimit(
	myProfileParams *dto.MyProfileParams,
) int {
	return (myProfileParams.Page - 1) * myProfileParams.Limit
}

func GetOrdersPerPage() (int, error) {
	return strconv.Atoi(os.Getenv("MAX_MY_PROFILE_ORDER_RESULTS"))
}

func GetServicesPerPage() (int, error) {
	return strconv.Atoi(os.Getenv("MAX_MY_PROFILE_SERVICE_RESULTS"))
}

func GetRequestsPerPage() (int, error) {
	return strconv.Atoi(os.Getenv("MAX_MY_PROFILE_REQUEST_RESULTS"))
}

func ParseMyProfileParams(ctx *fiber.Ctx) (*dto.MyProfileParams, error) {
	var myProfileParams dto.MyProfileParams

	var err error
	myProfileParams.Page, err = parseMyProfileOrdersParam("page", ctx)
	if err != nil && err != utils.ErrInvalidPathParam {
		return nil, err
	}
	myProfileParams.UserId = int(ctx.Locals("userId").(float64))

	if myProfileParams.UserId <= 0 {
		return nil, utils.ErrInvalidPathParam
	}

	return &myProfileParams, nil
}

func parseMyProfileOrdersParam(paramName string, ctx *fiber.Ctx) (int, error) {
	valStr := ctx.Query(paramName)
	if valStr == "" {
		return 0, utils.ErrInvalidPathParam
	}

	value, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, utils.ErrInvalidPathParam
	}

	return value, err
}

func ParseMyProfileDataFromRows[T any](rows *sqlx.Rows) (*[]T, int, error) {
	var dataJSON []byte
	var totalPages int
	var orderTableData []T

	if rows.Next() {
		err := rows.Scan(&dataJSON, &totalPages)
		if err != nil {
			return nil, 0, err
		}
	}

	if err := json.Unmarshal(dataJSON, &orderTableData); err != nil {
		return nil, 0, err
	}

	return &orderTableData, totalPages, nil
}

func ParseMyProfileChatByOrderIdFromRows(rows *sqlx.Rows) (*model.OrderChat, error) {
	var chatPartnherJSON []byte
	var chatMessagesJSON []byte
	var chatPartner model.ChatPartner
	var rawChatMessages = make([]model.ChatMessageRaw, 0)

	if !rows.Next() {
		return nil, utils.ErrNoDataFound
	}

	err := rows.Scan(&chatPartnherJSON, &chatMessagesJSON)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(chatPartnherJSON, &chatPartner); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(chatMessagesJSON, &rawChatMessages); err != nil {
		return nil, err
	}

	chatMessages, err := ConvertRawChatMessages(rawChatMessages)
	if err != nil {
		return nil, err
	}

	return &model.OrderChat{
		ChatPartner: chatPartner,
		Messages:    chatMessages,
	}, nil
}

func ConvertRawChatMessages(rawMessages []model.ChatMessageRaw) ([]model.ChatMessage, error) {
	var chatMessages = make([]model.ChatMessage, len(rawMessages))

	for i, msg := range rawMessages {
		decodedContent, err := utils.ConvertBase64ToBytes(msg.Content)
		if err != nil {
			return nil, err
		}

		decodedContentIV, err := utils.ConvertBase64ToBytes(msg.ContentIV)
		if err != nil {
			return nil, err
		}

		chatMessages[i] = model.ChatMessage{
			Id:        msg.Id,
			SenderId:  msg.SenderId,
			Content:   decodedContent,
			ContentIV: decodedContentIV,
			SentAt:    msg.SentAt,
			Files:     msg.Files,
			Type:      msg.Type,
		}
	}

	return chatMessages, nil
}

func ParseMyProfileRequestsParams(ctx *fiber.Ctx) (*dto.MyProfileParams, error) {
	var myProfileParams dto.MyProfileParams

	var err error
	myProfileParams.Page, err = parseMyProfileOrdersParam("page", ctx)
	if err != nil && err != utils.ErrInvalidPathParam {
		return nil, err
	}

	myProfileParams.Status, err = parseMyProfileOrdersParam("status", ctx)
	if err != nil && err != utils.ErrInvalidPathParam {
		return nil, err
	}

	myProfileParams.UserId = int(ctx.Locals("userId").(float64))

	if myProfileParams.UserId <= 0 {
		return nil, utils.ErrInvalidPathParam
	}

	return &myProfileParams, nil
}
