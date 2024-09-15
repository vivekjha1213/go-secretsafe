package secretsafe

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStorage is a mock implementation of the Storage interface
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Save(namespace, key, value string) error {
	args := m.Called(namespace, key, value)
	return args.Error(0)
}

func (m *MockStorage) Load(namespace, key string) (string, error) {
	args := m.Called(namespace, key)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Delete(namespace, key string) error {
	args := m.Called(namespace, key)
	return args.Error(0)
}

// MockCache is a mock implementation of the Cache interface
type MockCache struct {
	mock.Mock
}

func (m *MockCache) SetCache(namespace, key, value string, duration time.Duration) {
	m.Called(namespace, key, value, duration)
}

func (m *MockCache) GetCache(namespace, key string) (string, bool) {
	args := m.Called(namespace, key)
	return args.String(0), args.Bool(1)
}

func (m *MockCache) DeleteCache(namespace, key string) {
	m.Called(namespace, key)
}

func TestSetSecret(t *testing.T) {
	storage := new(MockStorage)
	cache := new(MockCache)
	sm := NewSecretManager(storage, cache)

	// Setup expectations
	storage.On("Save", "namespace1", "key1", mock.Anything).Return(nil)
	cache.On("SetCache", "namespace1", "key1", "value1", 5*time.Minute).Return()

	// Call the method
	err := sm.SetSecret("namespace1", "key1", "value1")

	// Assert expectations
	assert.NoError(t, err)
	storage.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestGetSecret(t *testing.T) {
	storage := new(MockStorage)
	cache := new(MockCache)
	sm := NewSecretManager(storage, cache)

	// Setup expectations
	cache.On("GetCache", "namespace1", "key1").Return("", false)
	storage.On("Load", "namespace1", "key1").Return("encryptedValue", nil)
	Encrypt = func(value string) (string, error) {
		return value, nil
	}
	Decrypt = func(encrypted string) (string, error) {
		return encrypted, nil
	}
	cache.On("SetCache", "namespace1", "key1", "value1", 5*time.Minute).Return()

	// Call the method
	value, err := sm.GetSecret("namespace1", "key1")

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, "value1", value)
	storage.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestDeleteSecret(t *testing.T) {
	storage := new(MockStorage)
	cache := new(MockCache)
	sm := NewSecretManager(storage, cache)

	// Setup expectations
	storage.On("Delete", "namespace1", "key1").Return(nil)
	cache.On("DeleteCache", "namespace1", "key1").Return()

	// Call the method
	err := sm.DeleteSecret("namespace1", "key1")

	// Assert expectations
	assert.NoError(t, err)
	storage.AssertExpectations(t)
	cache.AssertExpectations(t)
}
