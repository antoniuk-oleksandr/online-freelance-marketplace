package body

import "github.com/golang-jwt/jwt/v4"

type GoogleJwtClaims struct {
	Email      string `json:"email"`
	PicURL     string `json:"picture"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	jwt.RegisteredClaims
}
