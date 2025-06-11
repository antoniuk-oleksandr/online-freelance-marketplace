package model

type UserData struct {
	Id             int `json:"user_id" db:"user_id"`
	Username       string `json:"username" db:"username"`
	Avatar         string `json:"avatar" db:"avatar"`
	PrivateKey     string `json:"private_key" db:"private_key"`
	PrivateKeyIV   string `json:"private_key_iv" db:"private_key_iv"`
	PrivateKeySalt string `json:"private_key_salt" db:"private_key_salt"`
	MasterKey      string `json:"master_key" db:"master_key"`
}
