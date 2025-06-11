package model

type UserSessionData struct {
	MasterKey    []byte                 `json:"master_key" db:"master_key"`
	ChatPartners []ChatPartnerPublicKey `json:"chat_partners" db:"chat_partners"`
}
