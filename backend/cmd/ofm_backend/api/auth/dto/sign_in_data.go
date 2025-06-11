package dto

type SignInData struct {
	UserData     *UserData              `json:"userData"`
	ChatPartners []ChatPartnerPublicKey `json:"chatPartners"`
}
