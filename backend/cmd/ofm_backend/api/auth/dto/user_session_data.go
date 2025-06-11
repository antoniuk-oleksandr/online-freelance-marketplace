package dto

type UserSessionData struct {
	MasterKey    []byte                 `json:"masterKey"`
	ChatPartners []ChatPartnerPublicKey `json:"chatPartners"`
}
