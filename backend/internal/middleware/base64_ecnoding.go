package middleware

import b64 "encoding/base64"

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