package controller

import (
	"log"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/service"
	"ofm_backend/cmd/ofm_backend/utils"
	"strconv"

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

func (mpc *myProfileController) GetMyProfileServices(ctx *fiber.Ctx) error {
	paginationParams, err := helpers.ParseMyProfileParams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	servicesResponse, err := mpc.myProfileService.GetMyProfileServices(paginationParams)
	if err != nil {
		log.Println("err:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(servicesResponse)
}

func (mpc *myProfileController) GetMyProfileOrders(ctx *fiber.Ctx) error {
	paginationParams, err := helpers.ParseMyProfileParams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ordersResponse, err := mpc.myProfileService.GetMyProfileOrders(paginationParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ordersResponse)
}

func (mpc *myProfileController) GetMyProfileRequests(ctx *fiber.Ctx) error {
	params, err := helpers.ParseMyProfileRequestsParams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	requestsResponse, err := mpc.myProfileService.GetMyProfileRequests(params)
	if err != nil {
		log.Println("err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(requestsResponse)
}

func (mpc *myProfileController) GetMyProfileOrderChat(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("orderId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidPathParam.Error(),
		})
	}
	
	orderChat, err := mpc.myProfileService.GetMyProfileChatByOrderId(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(orderChat)
}