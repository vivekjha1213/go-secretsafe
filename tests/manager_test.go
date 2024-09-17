// tests/manager_test.go

package tests

import (
	"testing"

	"github.com/vivekjha1213/go-secretsafe/pkg/secretsafe"
)

type MockStorage struct {
	data map[string]map[string]string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		data: make(map[string]map[string]string),
	}
}

func (m *MockStorage) Set(namespace, key, value string) error {
	if _, ok := m.data[namespace]; !ok {
		m.data[namespace] = make(map[string]string)
	}
	m.data[namespace][key] = value
	return nil
}

func (m *MockStorage) Get(namespace, key string) (string, error) {
	if ns, ok := m.data[namespace]; ok {
		if value, ok := ns[key]; ok {
			return value, nil
		}
	}
	return "", secretsafe.ErrKeyNotFound
}

func (m *MockStorage) Delete(namespace, key string) error {
	if ns, ok := m.data[namespace]; ok {
		delete(ns, key)
	}
	return nil
}

func TestSecretManager(t *testing.T) {
	storage := NewMockStorage()
	cache := secretsafe.NewCache()
	manager := secretsafe.NewSecretManager(storage, cache)

	// Test SetSecret
	err := manager.SetSecret("test", "key1", "value1")
	if err != nil {
		t.Errorf("SetSecret failed: %v", err)
	}

	// Test GetSecret
	value, err := manager.GetSecret("test", "key1")
	if err != nil {
		t.Errorf("GetSecret failed: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %s", value)
	}

	// Test DeleteSecret
	err = manager.DeleteSecret("test", "key1")
	if err != nil {
		t.Errorf("DeleteSecret failed: %v", err)
	}

	// Verify deletion
	_, err = manager.GetSecret("test", "key1")
	if err != secretsafe.ErrKeyNotFound {
		t.Errorf("Expected ErrKeyNotFound after deletion, got %v", err)
	}
}