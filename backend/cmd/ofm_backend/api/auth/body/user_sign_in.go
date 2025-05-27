package body

type SignInBody struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
	KeepSignedIn    bool   `json:"keepSignedIn"`
}
