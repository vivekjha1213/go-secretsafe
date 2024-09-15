package secretsafe

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Encrypt encrypts a secret using AES-256 GCM
func Encrypt(value string) (string, error) {
	key := []byte("a very very very very secret key") // 32 bytes for AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(value), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts an encrypted secret
func Decrypt(encryptedValue string) (string, error) {
	key := []byte("a very very very very secret key") // 32 bytes for AES-256

	data, err := hex.DecodeString(encryptedValue)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
