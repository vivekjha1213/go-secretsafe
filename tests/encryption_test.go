package tests

import (
	"testing"

	"github.com/vivekjha1213/go-secretsafe/pkg/secretsafe"
)



func TestEncryptDecrypt(t *testing.T) {
	// Example key for encryption/decryption
	key := "jshafqhsfiq12345649765hruwrjb&*hbfsbfsurg@"

	testCases := []struct {
		name      string
		plaintext string
	}{
		{"Empty string", ""},
		{"Short string", "Hello, World!"},
		{"Long string", "This is a much longer string that we'll use to test our encryption and decryption functions to ensure they work correctly with various input lengths."},
		{"Special characters", "!@#$%^&*()_+{}[]|\\:;\"'<>,.?/~`"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt the plaintext
			ciphertext, err := secretsafe.Encrypt(tc.plaintext, key)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			// Ensure the ciphertext is different from the plaintext
			if ciphertext == tc.plaintext {
				t.Errorf("Ciphertext is identical to plaintext")
			}

			// Decrypt the ciphertext
			decrypted, err := secretsafe.Decrypt(ciphertext, key)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			// Check if the decrypted text matches the original plaintext
			if decrypted != tc.plaintext {
				t.Errorf("Decrypted text does not match original plaintext. Got %q, want %q", decrypted, tc.plaintext)
			}
		})
	}
}

func TestDecryptInvalidInput(t *testing.T) {
	key := "my-strong-secret-key"

	testCases := []struct {
		name       string
		ciphertext string
	}{
		{"Empty string", ""},
		{"Invalid base64", "This is not valid base64!"},
		{"Too short after decoding", "SGVsbG8="}, // "Hello" in base64, which is too short for a valid ciphertext
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := secretsafe.Decrypt(tc.ciphertext, key)
			if err == nil {
				t.Errorf("Expected an error when decrypting invalid input, but got none")
			}
		})
	}
}
