package secretsafe

import (
	"fmt"
	"time"
)

// SecretManager manages secrets
type SecretManager struct {
	storage Storage
	cache   Cache
}

// NewSecretManager creates a new instance of SecretManager
func NewSecretManager(storage Storage, cache Cache) *SecretManager {
	return &SecretManager{
		storage: storage,
		cache:   cache,
	}
}

// SetSecret stores a secret securely
func (sm *SecretManager) SetSecret(namespace, key, value string) error {
	encryptedValue, err := Encrypt(value)
	if err != nil {
		return fmt.Errorf("failed to encrypt secret: %w", err)
	}

	if err := sm.storage.Save(namespace, key, encryptedValue); err != nil {
		return fmt.Errorf("failed to store secret: %w", err)
	}

	sm.cache.SetCache(namespace, key, value, 5*time.Minute)
	return nil
}

// GetSecret retrieves a secret securely
func (sm *SecretManager) GetSecret(namespace, key string) (string, error) {
	// Check the cache first
	if value, found := sm.cache.GetCache(namespace, key); found {
		return value, nil
	}

	// If not found in cache, retrieve from storage
	encryptedValue, err := sm.storage.Load(namespace, key)
	if err != nil {
		return "", fmt.Errorf("failed to load secret: %w", err)
	}

	// Decrypt the secret
	value, err := Decrypt(encryptedValue)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt secret: %w", err)
	}

	// Store in cache
	sm.cache.SetCache(namespace, key, value, 5*time.Minute)
	return value, nil
}

// DeleteSecret removes a secret
func (sm *SecretManager) DeleteSecret(namespace, key string) error {
	if err := sm.storage.Delete(namespace, key); err != nil {
		return fmt.Errorf("failed to delete secret: %w", err)
	}
	sm.cache.DeleteCache(namespace, key)
	return nil
}
