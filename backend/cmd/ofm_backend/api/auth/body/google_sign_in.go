package body

type GoogleSignInBody struct {
	AccessToken  string `json:"accessToken"`
	KeepSignedIn bool   `json:"keepSignedIn"`
}
