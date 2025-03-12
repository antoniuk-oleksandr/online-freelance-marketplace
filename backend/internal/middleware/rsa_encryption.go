package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var privateKey *rsa.PrivateKey

func generateKeys() error {
	fmt.Println("Generating new key pair...")
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	return saveKeyPair(privateKey)
}

func TryToLoadRSAKeys() {
	var err error
	if privateKey, err = loadRSAKey(); err != nil {
		generateKeys()
	}
}

func loadRSAKey() (*rsa.PrivateKey, error) {
	if _, err := os.Stat(os.Getenv("RSA_PRIVATE_KEY_PATH")); err != nil {
		return nil, errors.New("private key file not found")
	}

	data, err := os.ReadFile(os.Getenv("RSA_PRIVATE_KEY_PATH"))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func saveKeyPair(privateKey *rsa.PrivateKey) error {
	// Save private key
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}
	if err := os.WriteFile(os.Getenv("RSA_PRIVATE_KEY_PATH"), pem.EncodeToMemory(privPem), 0600); err != nil {
		return err
	}

	// Save public key
	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	pubPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	return os.WriteFile(os.Getenv("RSA_PUBLIC_KEY_PATH"), pem.EncodeToMemory(pubPem), 0644)
}

func GetPublicKey() ([]byte, error) {
	return os.ReadFile(os.Getenv("RSA_PUBLIC_KEY_PATH"))
}

func DecryptRSAData[T any](data string) (*T, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(decryptedData, &result)

	return &result, err
}
