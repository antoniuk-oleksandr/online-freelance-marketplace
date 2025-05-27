package body

type GoogleSignUpBody struct {
	AccessToken    string `json:"accessToken"`
	PrivateKey     []byte `json:"privateKey"`
	PublicKey      []byte `json:"publicKey"`
	MasterKey      []byte `json:"masterKey"`
	PrivateKeyIV   []byte `json:"privateKeyIV"`
	PrivateKeySalt []byte `json:"privateKeySalt"`
	KeepSignedIn   bool   `json:"keepSignedIn"`
}
