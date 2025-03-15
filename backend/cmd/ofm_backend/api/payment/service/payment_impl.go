package service

import (
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/dto"
	"ofm_backend/cmd/ofm_backend/api/payment/helpers"
	"ofm_backend/cmd/ofm_backend/api/payment/repository"
	"ofm_backend/cmd/ofm_backend/enum"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/utils/aes_encryption"
	filereader "ofm_backend/internal/file_reader"
	"ofm_backend/internal/mailer"
	"os"
)

type paymentService struct {
	pRepo repository.PaymentRepository
}

func NewPaymentService(pRepo repository.PaymentRepository) PaymentService {
	return &paymentService{
		pRepo: pRepo,
	}
}

func (pServ *paymentService) ProcessPayment(
	paymentRequestBody body.EncryptedPaymentData,
	username string,
) (dto.PaymentResponse, error) {
	var paymentResponse dto.PaymentResponse

	decryptedData, err := helpers.DecryptPaymentData(paymentRequestBody, username)
	if err != nil {
		return paymentResponse, err
	}

	tx, err := pServ.pRepo.CreateTransaction()

	if !helpers.ValidateCardNumber(decryptedData.CardCredentials.CardNumber) {
		pServ.pRepo.RollBackTransaction(tx)
		return paymentResponse, utils.ErrInvalidCardNumber
	}

	if err := pServ.CreatePayment(decryptedData); err != nil {
		pServ.pRepo.RollBackTransaction(tx)
		return paymentResponse, err
	}

	if err := pServ.ProcessTransaction(decryptedData.PaymentId); err != nil {
		pServ.pRepo.RollBackTransaction(tx)
		return paymentResponse, err
	}

	orderId, err := pServ.CreateOrder(*decryptedData, &paymentResponse)
	if err != nil {
		pServ.pRepo.RollBackTransaction(tx)
		return paymentResponse, err
	}

	if err := pServ.SendOrderReceipt(username, orderId, decryptedData.UserTimezone); err != nil {
		pServ.pRepo.RollBackTransaction(tx)
		return paymentResponse, err
	}

	pServ.pRepo.CommitTransaction(tx)
	return paymentResponse, err
}

func (pServ *paymentService) SendOrderReceipt(
	username string,
	orderId int64,
	userTimezone string,
) error {
	serviceFee := os.Getenv("SERVICE_FEE")
	paymentReceipt, err := pServ.pRepo.GetPaymentReceipt(orderId, serviceFee)
	if err != nil {
		return err
	}

	html, err := filereader.GetHTMLTempalate("order_receipt.html")
	if err != nil {
		return err
	}
	html = helpers.FormatHtml(html, username, paymentReceipt, userTimezone)

	return mailer.SendEmail(paymentReceipt.CustomerEmail, "Order receipt", html, "text/html")
}

func (pServ *paymentService) CreateOrder(
	decryptedData body.DecryptedPaymentData,
	paymentResponse *dto.PaymentResponse,
) (int64, error) {
	orderId, err := pServ.pRepo.AddOrder(decryptedData)
	if err != nil {
		return 0, utils.ErrOrderCreationFailed
	}
	paymentResponse.OrderId = orderId
	paymentResponse.Success = true

	return orderId, nil
}

func (pServ *paymentService) CreatePayment(
	decryptedData *body.DecryptedPaymentData,
) error {
	encryptedCardNum, err := aes_encryption.Encrypt((*decryptedData).CardCredentials.CardNumber)
	(*decryptedData).CardCredentials.CardNumber = encryptedCardNum

	paymendId, err := pServ.pRepo.AddPayment(*decryptedData)
	if err != nil || paymendId == 0 {
		return utils.ErrPaymentCreationFailed
	}
	decryptedData.PaymentId = paymendId

	return nil
}

func (pServ *paymentService) ProcessTransaction(
	paymentId int64,
) error {
	success, err := helpers.DoPayPalPayment()
	if err != nil || !success {
		_, _ = pServ.pRepo.UpdatePaymentStatus(enum.Failed, paymentId)

		return utils.ErrPayPalPaymentFailed
	}

	success, err = pServ.pRepo.UpdatePaymentStatus(enum.Completed, paymentId)
	if err != nil || !success {
		return utils.ErrPaymentStatusUpdateFailed
	}

	return nil
}
