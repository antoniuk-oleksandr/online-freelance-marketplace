package model

type UserSessionData struct {
	Authenticated bool   `json:"authenticated" db:"authenticated"`
	MasterKey     []byte `json:"masterKey" db:"master_key"`
}
