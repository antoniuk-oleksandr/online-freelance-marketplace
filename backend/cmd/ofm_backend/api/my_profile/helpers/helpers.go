package helpers

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CalcMyProfileOrdersOffset(
	ordersPaginationParams *dto.OrdersPaginationParams,
) int {
	return (ordersPaginationParams.Page - 1) * ordersPaginationParams.OrdersPerPage
}

func GetOrdersPerPage() (int, error) {
	return strconv.Atoi(os.Getenv("MAX_MY_PROFILE_ORDER_RESULTS"))
}

func ParseMyProfileOrdersParams(ctx *fiber.Ctx) (*dto.OrdersPaginationParams, error) {
	var ordersPaginationParams dto.OrdersPaginationParams

	var err error
	ordersPaginationParams.Page, err = parseMyProfileOrdersParam("page", ctx)
	if err != nil && err != utils.ErrInvalidPathParam {
		return nil, err
	}
	ordersPaginationParams.UserId = int(ctx.Locals("userId").(float64))

	if ordersPaginationParams.UserId <= 0 {
		return nil, utils.ErrInvalidPathParam
	}

	return &ordersPaginationParams, nil
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
