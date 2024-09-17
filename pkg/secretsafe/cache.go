// pkg/secretsafe/cache.go

package secretsafe

import (
	"sync"
)

type Cache interface {
	Set(namespace, key, value string)
	Get(namespace, key string) (string, bool)
	Delete(namespace, key string)
}

type InMemoryCache struct {
	data map[string]map[string]string
	mu   sync.RWMutex
}

func NewCache() Cache {
	return &InMemoryCache{
		data: make(map[string]map[string]string),
	}
}

func (c *InMemoryCache) Set(namespace, key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[namespace]; !ok {
		c.data[namespace] = make(map[string]string)
	}
	c.data[namespace][key] = value
}

func (c *InMemoryCache) Get(namespace, key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if ns, ok := c.data[namespace]; ok {
		if value, ok := ns[key]; ok {
			return value, true
		}
	}
	return "", false
}

func (c *InMemoryCache) Delete(namespace, key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ns, ok := c.data[namespace]; ok {
		delete(ns, key)
	}
}