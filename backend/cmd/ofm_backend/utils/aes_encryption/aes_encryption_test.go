package aes_encryption

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptEmptyString(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "12345678901234567890123456789012")
	defer os.Unsetenv("ENCRYPTION_KEY")
	
	plainText := ""
	encryptedtext, err := Encrypt(plainText)
	assert.NoError(t, err, "Encryption should not return an error")
	assert.NotEmpty(t, encryptedtext, "Encrypted text should not be empty")
	
}

func TestEcnryptBrokenKey(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "12345678901234567890123456789012dd")
	defer os.Unsetenv("ENCRYPTION_KEY")
	
	plainText := "TestEcnryptBrokenKey"
	encryptedtext, err := Encrypt(plainText)
	assert.Error(t, err, "Encryption using a broken key should return an error")
	assert.Empty(t, encryptedtext, "Encrypted text should be empty")
}

func TestEncryptDecryptCycle(t *testing.T){
	os.Setenv("ENCRYPTION_KEY", "12345678901234567890123456789012")
	defer os.Unsetenv("ENCRYPTION_KEY")
	
	plainText := "hello world!"
	ecnryptedText, err := Encrypt(plainText)
	assert.NoError(t, err, "Encryption should not return an error")
	
	decryptedText, err := Decrypt(ecnryptedText)
	assert.NoError(t, err, "Decryption should not return an error")
	
	assert.Equal(t, plainText, decryptedText, "Decrypted text should be equal to the original text")
}