package dto

type PaymentCredentials struct {
	CardNumber     string `json:"cardNumber"`
	CardHolderName string `json:"cardHolderName"`
	ExpiryDate     string `json:"expiryDate"`
	SecurityCode   string `json:"securityCode"`
}
