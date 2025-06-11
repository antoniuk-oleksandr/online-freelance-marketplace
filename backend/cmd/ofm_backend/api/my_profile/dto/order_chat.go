package dto

type OrderChat struct {
	Messages    []ChatMessage `json:"messages"`
	ChatPartner ChatPartner   `json:"chatPartner"`
}
