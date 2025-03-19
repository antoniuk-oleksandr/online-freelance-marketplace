package controller

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/service"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type myProfileController struct {
	myProfileService service.MyProfileService
}

func NewMyProfileController(myProfileService service.MyProfileService) MyProfileController {
	return &myProfileController{
		myProfileService: myProfileService,
	}
}

func (mpc *myProfileController) GetMyProfileOrders(ctx *fiber.Ctx) error {
	ordersPaginationParams, err := helpers.ParseMyProfileOrdersParams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ordersResponse, err := mpc.myProfileService.GetMyProfileOrders(ordersPaginationParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ordersResponse)
}
