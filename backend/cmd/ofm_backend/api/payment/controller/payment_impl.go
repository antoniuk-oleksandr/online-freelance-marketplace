package controller

import (
	"github.com/gofiber/fiber/v2"
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/service"
	main_utils "ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/rsa_encryption"
)

type paymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &paymentController{
		paymentService: paymentService,
	}
}

func (p *paymentController) ProcessPayment(ctx *fiber.Ctx) error {
	var encryptedPaymentData body.EncryptedPaymentData
	if err := ctx.BodyParser(&encryptedPaymentData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": main_utils.ErrInvalidRequestBody.Error(),
		})
	}

	var username string = ctx.Locals("username").(string)

	paymentResponse, err := p.paymentService.ProcessPayment(encryptedPaymentData, username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(paymentResponse)
}

func (p *paymentController) GetPublicKey(ctx *fiber.Ctx) error {
	publicKey, err := rsa_encryption.GetPublicKey()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).SendString(string(publicKey))
}
