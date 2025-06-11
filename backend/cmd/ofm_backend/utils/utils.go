package utils

import (
	"strconv"
	"time"

	"ofm_backend/cmd/ofm_backend/api/filter_params/dto"
	"ofm_backend/cmd/ofm_backend/api/filter_params/model"

	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnvValues() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func Contains(slice []string, element string) bool {
	return slices.Contains(slice, element)
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05.999999999")
}

func ConvertTimeToSting(timeData time.Time) string {
	return timeData.Format("2006-01-02 15:04:05.999999999")
}

func MapFilterParamModelsToDTOs[T model.Identifiable](items *[]T) *[]dto.FilterItem {
	if items == nil {
		return nil
	}

	dtos := make([]dto.FilterItem, len(*items))

	for i, item := range *items {
		dtos[i] = dto.FilterItem{
			ID:   item.GetID(),
			Name: item.GetName(),
		}
	}

	return &dtos
}

func MapFilterParamModelToDTO[T model.Identifiable](item *T) *dto.FilterItem {
	if item == nil {
		return nil
	}

	return &dto.FilterItem{
		ID:   (*item).GetID(),
		Name: (*item).GetName(),
	}
}

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func GetUserIdFromLocals(paramName string, ctx *fiber.Ctx) (int, error) {
	var userId int
	userIdInterface := ctx.Locals("userId")
	if userIdInterface == nil {
		return userId, ctx.SendStatus(fiber.StatusUnauthorized)
	}
	userId = int(userIdInterface.(float64))
	return userId, nil
}
