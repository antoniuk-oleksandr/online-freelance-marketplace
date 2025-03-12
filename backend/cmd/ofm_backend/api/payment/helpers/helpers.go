package helpers

import (
	"fmt"
	"ofm_backend/cmd/ofm_backend/api/payment/body"
	"ofm_backend/cmd/ofm_backend/api/payment/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/internal/middleware"
	"regexp"
	"strings"
	"time"
)

func FormatHtml(
	html string,
	username string,
	paymentReceipt *model.PaymentReceipt,
	userTimezone string,
) string {
	orderId := fmt.Sprintf("#%d", paymentReceipt.OrderId)
	price := fmt.Sprintf("%.2f$", paymentReceipt.Price)
	serviceFee := fmt.Sprintf("%.2f$", paymentReceipt.ServiceFee)
	total := fmt.Sprintf("%.2f$", paymentReceipt.Total)
	orderDate := formatOrderDate(paymentReceipt.Date, userTimezone)

	html = strings.Replace(html, "{username}", username, -1)
	html = strings.Replace(html, "{order_id}", orderId, -1)
	html = strings.Replace(html, "{service_name}", paymentReceipt.ServiceName, -1)
	html = strings.Replace(html, "{package_name}", paymentReceipt.PackageName, -1)
	html = strings.Replace(html, "{price}", price, -1)
	html = strings.Replace(html, "{service_fee}", serviceFee, -1)
	html = strings.Replace(html, "{total}", total, -1)
	html = strings.Replace(html, "{order_date}", orderDate, -1)
	html = strings.Replace(html, "{year}", utils.IntToString(time.Now().Year()), -1)

	return html
}

func formatOrderDate(date time.Time, userTimezone string) string {
	location, _ := time.LoadLocation(userTimezone)
	localTime := date.In(location)

	return localTime.Format("2006-01-02 15:04:05")
}

func DecryptPaymentData(
	paymentRequestBody body.EncryptedPaymentData,
	username string,
) (*body.DecryptedPaymentData, error) {
	decryptedData, err := middleware.DecryptRSAData[body.DecryptedPaymentData](paymentRequestBody.Data)
	if err != nil {
		return nil, utils.ErrDecryptionFailed
	}
	(*decryptedData).Username = username

	return decryptedData, nil
}

func ValidateCardNumber(
	cardNumber string,
) bool {
	if len(cardNumber) > 19 || len(cardNumber) < 13 {
		return false
	}

	match, err := regexp.MatchString("^[0-9]+$", cardNumber)
	if err != nil || !match {
		return false
	}

	sum := 0
	shouldBeDoubled := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if shouldBeDoubled {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		shouldBeDoubled = !shouldBeDoubled
		sum += digit
	}

	return sum%10 == 0
}

func DoPayPalPayment() (bool, error) {
	// cannot be implemented cuz I've got no Business PayPal account
	return true, nil
}
