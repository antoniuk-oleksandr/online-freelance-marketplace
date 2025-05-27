package body

type SignUpBody struct {
	Email          string `json:"email" db:"email"`
	FirstName      string `json:"firstName" db:"first_name"`
	Password       string `json:"password" db:"password"`
	Surname        string `json:"surname" db:"surname"`
	Username       string `json:"username" db:"username"`
	PrivateKey     []byte `json:"privateKey" db:"private_key"`
	PrivateKeyIV   []byte `json:"privateKeyIV" db:"private_key_iv"`
	PrivateKeySalt []byte `json:"privateKeySalt" db:"private_key_salt"`
	PublicKey      []byte `json:"publicKey" db:"public_key"`
	MasterKey      []byte `json:"masterKey" db:"master_key"`
}
