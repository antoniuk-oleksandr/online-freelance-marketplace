package model

type SignInData struct {
	UserData     UserData               `json:"user_data" db:"user_data"`
	ChatPartners []ChatPartnerPublicKey `json:"chat_partners" db:"chat_partners"`
}
