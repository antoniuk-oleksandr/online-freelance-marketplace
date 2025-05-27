package model

type SignResponse struct {
	AccessToken string
	RefreshToken string
	AccessTokenExpiresAt int
	RefreshTokenExpiresAt int
}