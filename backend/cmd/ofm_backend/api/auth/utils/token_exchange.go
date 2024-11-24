package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"os"
	"strings"
)


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
