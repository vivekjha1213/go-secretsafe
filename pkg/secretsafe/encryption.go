package secretsafe

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "errors"
    "io"
)

// createAESKey generates a 32-byte key using SHA-256 to ensure compatibility with AES-256.
func createAESKey(key string) []byte {
    hashedKey := sha256.Sum256([]byte(key)) // 32-byte key from SHA-256
    return hashedKey[:]
}

// Encrypt encrypts the plaintext using AES encryption with a dynamically hashed key.
func Encrypt(plaintext string, key string) (string, error) {
    aesKey := createAESKey(key) // Generate the AES key from the provided key string

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        return "", err
    }

    plaintextBytes := []byte(plaintext)
    ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)

    // Return the encrypted data as a base64-encoded string
    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a base64-encoded ciphertext using AES encryption with a dynamically hashed key.
func Decrypt(ciphertext string, key string) (string, error) {
    aesKey := createAESKey(key) // Generate the AES key from the provided key string

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        return "", err
    }

    decodedCiphertext, err := base64.URLEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    if len(decodedCiphertext) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv := decodedCiphertext[:aes.BlockSize]
    decodedCiphertext = decodedCiphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(decodedCiphertext, decodedCiphertext)

    // Return the decrypted plaintext as a string
    return string(decodedCiphertext), nil
}
