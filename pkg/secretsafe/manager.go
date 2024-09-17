// pkg/secretsafe/manager.go

package secretsafe

import (
	"sync"
)

type SecretManager struct {
	storage Storage
	cache   Cache
	mu      sync.RWMutex
}

func NewSecretManager(storage Storage, cache Cache) *SecretManager {
	return &SecretManager{
		storage: storage,
		cache:   cache,
	}
}

func (sm *SecretManager) SetSecret(namespace, key, value string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	err := sm.storage.Set(namespace, key, value)
	if err != nil {
		return err
	}

	sm.cache.Set(namespace, key, value)
	return nil
}

func (sm *SecretManager) GetSecret(namespace, key string) (string, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Try to get from cache first
	if value, found := sm.cache.Get(namespace, key); found {
		return value, nil
	}

	// If not in cache, get from storage
	value, err := sm.storage.Get(namespace, key)
	if err != nil {
		return "", err
	}

	// Update cache
	sm.cache.Set(namespace, key, value)
	return value, nil
}

func (sm *SecretManager) DeleteSecret(namespace, key string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	err := sm.storage.Delete(namespace, key)
	if err != nil {
		return err
	}

	sm.cache.Delete(namespace, key)
	return nil
}