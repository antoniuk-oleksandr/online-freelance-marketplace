package body

type DecryptedPaymentData struct {
	CardCredentials CardCredentials `json:"cardCredentials"`
	PackageId       int             `json:"packageId"`
	ServiceId       int             `json:"serviceId"`
	PaymentId       int64           `json:"paymentId"`
	Username        string          `json:"username"`
	UserTimezone    string          `json:"userTimezone"`
}
