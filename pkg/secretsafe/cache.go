package secretsafe

import (
	"sync"
	"time"
)

type Cache struct {
	secrets map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	value   string
	expires time.Time
}

// NewCache initializes a new Cache
func NewCache() *Cache {
	return &Cache{
		secrets: make(map[string]cacheEntry),
	}
}

// SetCache stores a secret in cache with an expiration
func (c *Cache) SetCache(namespace, key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheKey := namespace + ":" + key
	c.secrets[cacheKey] = cacheEntry{
		value:   value,
		expires: time.Now().Add(ttl),
	}
}

// GetCache retrieves a secret from cache if it's still valid
func (c *Cache) GetCache(namespace, key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheKey := namespace + ":" + key
	entry, exists := c.secrets[cacheKey]

	if !exists || time.Now().After(entry.expires) {
		return "", false
	}

	return entry.value, true
}

// DeleteCache removes a secret from the cache
func (c *Cache) DeleteCache(namespace, key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheKey := namespace + ":" + key
	delete(c.secrets, cacheKey)
}
