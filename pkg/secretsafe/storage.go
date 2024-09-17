// pkg/secretsafe/storage.go

package secretsafe

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Define a custom error for "key not found"
var ErrKeyNotFound = fmt.Errorf("key not found")

type Storage interface {
	Set(namespace, key, value string) error
	Get(namespace, key string) (string, error)
	Delete(namespace, key string) error
}

type FileStorage struct {
	path string
	mu   sync.RWMutex
}

func NewStorage(path string) (Storage, error) {
	if path == "" {
		path = filepath.Join(os.TempDir(), "secretsafe")
	}

	err := os.MkdirAll(path, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &FileStorage{path: path}, nil
}

func (fs *FileStorage) Set(namespace, key, value string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	data := make(map[string]string)
	filePath := fs.getFilePath(namespace)

	// Read existing data
	if fileContent, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(fileContent, &data)
	}

	// Update data
	data[key] = value

	// Write updated data
	fileContent, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(filePath, fileContent, 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (fs *FileStorage) Get(namespace, key string) (string, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data := make(map[string]string)
	filePath := fs.getFilePath(namespace)

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal data: %w", err)
	}

	value, ok := data[key]
	if !ok {
		return "", ErrKeyNotFound // Use the custom error
	}

	return value, nil
}

func (fs *FileStorage) Delete(namespace, key string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	data := make(map[string]string)
	filePath := fs.getFilePath(namespace)

	// Read existing data
	if fileContent, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(fileContent, &data)
	}

	// Delete key
	delete(data, key)

	// Write updated data
	fileContent, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(filePath, fileContent, 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (fs *FileStorage) getFilePath(namespace string) string {
	return filepath.Join(fs.path, namespace+".json")
}
