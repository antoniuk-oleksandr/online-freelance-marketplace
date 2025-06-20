package utils

import (
	"encoding/base64"
	b64 "encoding/base64"
)

func EncodeString(text string) string {
	return b64.StdEncoding.EncodeToString([]byte(text))
}

func DecodeString(encodedText string) (string, error) {
	decodedBytes, err := b64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}
	
	return string(decodedBytes), nil
}

func ConvertToBase64(data []byte) string {
	return b64.StdEncoding.EncodeToString(data)
}

func ConvertBase64ToBytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}