package model

type ChatPartnerPublicKey struct {
	UserId    int    `json:"user_id" db:"user_id"`
	PublicKey string `json:"public_key" db:"public_key"`
}
