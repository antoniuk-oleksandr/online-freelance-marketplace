package dto

type ChatPartnerPublicKey struct {
	UserId    int    `json:"userId"`
	PublicKey []byte `json:"publicKey"`
}
