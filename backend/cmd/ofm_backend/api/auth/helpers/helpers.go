package helpers

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/auth/body"
	"ofm_backend/cmd/ofm_backend/api/auth/model"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateResetLink(token string) string {
	host := os.Getenv("FRONTEND_HOST")
	return fmt.Sprintf("%s/reset-password?token=%s", host, token)
}

func GetGoogleUserInfo(accessToken string) (*model.GoogleUserInfo, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var userInfo model.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &userInfo, nil
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

func CreateSignCookies(ctx *fiber.Ctx, signResponse *model.SignResponse, keepSignedIn bool) {
	ctx.Cookie(&fiber.Cookie{
		Name:        "accessToken",
		Value:       signResponse.AccessToken,
		Expires:     time.Now().Add(time.Duration(signResponse.AccessTokenExpiresAt) * time.Minute),
		HTTPOnly:    true,
		Secure:      false,
		SameSite:    fiber.CookieSameSiteStrictMode,
		SessionOnly: !keepSignedIn,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:        "refreshToken",
		Value:       signResponse.RefreshToken,
		Expires:     time.Now().Add(time.Duration(signResponse.RefreshTokenExpiresAt) * time.Minute),
		HTTPOnly:    true,
		Secure:      false,
		SameSite:    fiber.CookieSameSiteStrictMode,
		SessionOnly: !keepSignedIn,
	})
}

func ClearCookies(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})
}
