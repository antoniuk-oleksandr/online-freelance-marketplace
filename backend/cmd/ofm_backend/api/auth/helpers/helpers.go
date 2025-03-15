package helpers

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateResetLink(token string) string {
	host := os.Getenv("FRONTEND_HOST")
	port := os.Getenv("FRONTEND_PORT")
	return fmt.Sprintf("http://%s:%s/reset-password?token=%s", host, port, token)
}

func ExchangeAuthCodeForToken(code string) (*body.TokenResponse, error) {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURI := os.Getenv("GOOGLE_REDIRECT_URI")

	if googleClientID == "" || googleClientSecret == "" || redirectURI == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	data := fmt.Sprintf(
		"code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code",
		code, googleClientID, googleClientSecret, redirectURI,
	)

	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send token request: %v", err)
	}
	defer resp.Body.Close()

	var tokenResponse body.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %v", err)
	}

	return &tokenResponse, nil
}

func ParseJWT(tokenString string) (*body.GoogleJwtClaims, error) {
	var claims body.GoogleJwtClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid kid in token header")
		}

		return getPublicKey(kid)
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func getPublicKey(kid string) (*rsa.PublicKey, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Google's public keys: %v", err)
	}
	defer resp.Body.Close()

	var certResp body.GoogleCert
	if err := json.NewDecoder(resp.Body).Decode(&certResp); err != nil {
		return nil, fmt.Errorf("failed to decode cert response: %v", err)
	}

	for _, key := range certResp.Keys {
		if key.Kid == kid {
			n, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, fmt.Errorf("failed to decode modulus (n): %v", err)
			}

			e, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, fmt.Errorf("failed to decode exponent (e): %v", err)
			}

			return &rsa.PublicKey{
				N: new(big.Int).SetBytes(n),
				E: int(new(big.Int).SetBytes(e).Int64()),
			}, nil
		}
	}

	return nil, fmt.Errorf("public key not found for kid: %s", kid)
}
