package aes_encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(text string) (string, error){
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		return "", errors.New("Invalid encryption key")
	}
	
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}
	
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encodedText string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		return "", errors.New("invalid encryption key")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}