package body

type CardCredentials struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpiryDate     string `json:"expiryDate"`
}
