package aes_encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(text string) (string, error) {
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

func GenerateAESIV() ([]byte, error) {
	nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nonce, err
    }
    
    return nonce, nil
}

func EncryptWithKey(plaintext, key string, initialVector []byte) (string, error) {
    keyBytes := sha256.Sum256([]byte(key))
    
    block, err := aes.NewCipher(keyBytes[:])
    if err != nil {
        return "", err
    }
    
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    ciphertext := aesGCM.Seal(nil, initialVector, []byte(plaintext), nil)
    finalCiphertext := append(initialVector, ciphertext...)

    return base64.StdEncoding.EncodeToString(finalCiphertext), nil
}


func DecryptWithKey(encryptedText, key string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(encryptedText)
    if err != nil {
        return "", err
    }

    keyBytes := sha256.Sum256([]byte(key))
    
    block, err := aes.NewCipher(keyBytes[:])
    if err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := aesGCM.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]

    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
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


func GenerateMasterKey() ([]byte, error) {
    masterKey := make([]byte, 32)
    _, err := rand.Read(masterKey)
    if err != nil {
        return nil, err
    }
    return masterKey, nil
}