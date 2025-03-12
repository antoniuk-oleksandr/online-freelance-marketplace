package controller

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/order/body"
	"ofm_backend/cmd/ofm_backend/api/order/service"
	"ofm_backend/cmd/ofm_backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (oc *orderController) SubmitOrderRequirements(ctx *fiber.Ctx) error {
	orderRequirementsRequestBody, err := oc.parseSubmitOrderRequirementsRequestBody(ctx)
	if err != nil {
		return err
	}

	if err := oc.orderService.SubmitOrderRequirements(orderRequirementsRequestBody); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order requirements submitted successfully",
	})
}

func (oc *orderController) GetOrderById(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id", "0"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrOrderNotFound.Error(),
		})
	}

	submitted, err := oc.orderService.CheckIfOrderRequirementsSubmitted(orderId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrUnexpectedError.Error(),
		})
	} else if submitted {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": utils.ErrAlreadySubmitted.Error(),
		})
	}

	response, err := oc.orderService.GetOrderById(orderId)
	if err != nil {
		var status int = fiber.StatusInternalServerError
		if err == utils.ErrOrderNotFound {
			status = fiber.StatusNotFound
		}

		return ctx.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (oc *orderController) getRequiredFormValue(
	name string, ctx *fiber.Ctx,
) (string, error) {
	value := ctx.FormValue(name, "")
	if value == "" {
		return "", ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	return value, nil
}

func (orderController) getRequiredFormIntValue(
	name string, ctx *fiber.Ctx,
) (int, error) {
	value := ctx.FormValue(name, "")
	if value == "" {
		return -1, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return -1, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	return intVal, nil
}

func (orderController *orderController) parseSubmitOrderRequirementsRequestBody(
	ctx *fiber.Ctx,
) (*body.OrderRequirementsBody, error) {
	var orderRequirementsBody body.OrderRequirementsBody
	var err error

	orderRequirementsBody.CustomerMessage, err = orderController.getRequiredFormValue("customerMessage", ctx)
	if err != nil {
		return nil, err
	}

	orderRequirementsBody.OrderId, err = orderController.getRequiredFormIntValue("orderId", ctx)
	if err != nil {
		return nil, err
	}

	answers, err := orderController.getRequiredFormValue("answers", ctx)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(answers), &orderRequirementsBody.Answers); err != nil {
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}

	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": utils.ErrInvalidRequestBody.Error(),
		})
	}
	orderRequirementsBody.Files = multipartForm.File["files"]

	return &orderRequirementsBody, nil
}
