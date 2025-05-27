package model

type OrderChat struct {
	Messages    []ChatMessage `json:"messages" db:"messages"`
	ChatPartner ChatPartner  `json:"chat_partner" db:"chat_partner"`
}
