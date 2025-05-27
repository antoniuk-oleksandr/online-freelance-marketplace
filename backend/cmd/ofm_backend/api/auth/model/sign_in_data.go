package model

type SignInData struct {
	Id             string `json:"id" db:"user_id"`
	Username       string `json:"username" db:"username"`
	Avatar         string `json:"avatar" db:"avatar"`
	PrivateKey     []byte `json:"privateKey" db:"private_key"`
	PrivateKeyIV   []byte `json:"privateKeyIV" db:"private_key_iv"`
	PrivateKeySalt []byte `json:"privateKeySalt" db:"private_key_salt"`
	MasterKey      []byte `json:"masterKey" db:"master_key"`
}
