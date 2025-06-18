package controller

import (
	"database/sql"
	"errors"
	"log"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
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

func (mpc *myProfileController) GetMyProfileServices(ctx *fiber.Ctx) error {
	paginationParams, err := helpers.ParseMyProfileParams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	servicesResponse, err := mpc.myProfileService.GetMyProfileServices(paginationParams)
	if err != nil {
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(requestsResponse)
}

func (mpc *myProfileController) GetMyProfileOrderChat(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderChat, err := mpc.myProfileService.GetMyProfileChatByOrderId(orderId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(orderChat)
}

func (mpc *myProfileController) GetMyProfileOrderOverview(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderOverview, err := mpc.myProfileService.GetMyProfileOverviewByOrderId(orderId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrOrderNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(orderOverview)
}

func (mpc *myProfileController) GetMyProfileOrderRequirements(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderQuestionArray, err := mpc.myProfileService.GetMyProfileRequirementsByOrderId(orderId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrOrderNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"questionsAnswers": orderQuestionArray,
	})
}

func (mpc *myProfileController) GetMyProfileOrderDelivery(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	deliveryData, err := mpc.myProfileService.GetMyProfileDeliveryByOrderId(orderId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils.ErrOrderNotFound.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(deliveryData)
}

func (mpc *myProfileController) AddMyProfileOrderDelivery(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		log.Println("err1", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		log.Println("err2", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	orderDeliveryBody := model.OrderDeliveryBody{
		Files:   multipartForm.File["files"],
		Message: ctx.FormValue("message", ""),
	}

	deliveryData, err := mpc.myProfileService.CompleteOrderWithDelivery(orderId, userId, orderDeliveryBody)
	if err != nil {
		log.Println("err3", err)
		if errors.Is(err, utils.ErrCompletingOrder) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": utils.ErrCompletingOrder.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(deliveryData)

}

func (mpc *myProfileController) AddMyProfileOrderReview(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	reviewRequestBody := model.ReviewRequestBody{
		OrderId: orderId,
		UserId:  userId,
	}
	err = ctx.BodyParser(&reviewRequestBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	orderReview, err := mpc.myProfileService.AddOrderReview(reviewRequestBody)
	if err != nil {
		log.Println("err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(orderReview)
}

func (mpc *myProfileController) GetMyProfileOrderReview(ctx *fiber.Ctx) error {
	orderId, err := helpers.GetOrderId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderReview, err := mpc.myProfileService.GetOrderReviewByOrderId(orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  utils.ErrOrderReviewNotFound.Error(),
				"status": orderReview.OrderStatus,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(orderReview)
}

func (mpc *myProfileController) UpdateMyProfileOrderStatus(ctx *fiber.Ctx) error {
	orderId, userId, err := helpers.GetOrderAndUserIds(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	statusId, err := ctx.ParamsInt("statusId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrInvalidPathParam.Error(),
		})
	}

	err = mpc.myProfileService.UpdateOrderStatusByOrderId(orderId, userId, statusId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
