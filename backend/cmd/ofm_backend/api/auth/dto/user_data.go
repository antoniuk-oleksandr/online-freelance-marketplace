package dto

type UserData struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Avatar         string `json:"avatar"`
	PrivateKey     []byte `json:"privateKey"`
	PrivateKeyIV   []byte `json:"privateKeyIV"`
	PrivateKeySalt []byte `json:"privateKeySalt"`
	MasterKey      []byte `json:"masterKey"`
}
