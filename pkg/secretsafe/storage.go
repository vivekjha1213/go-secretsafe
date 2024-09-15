package secretsafe

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Storage struct {
	filePath string
}

// NewStorage initializes storage
func NewStorage(path string) (*Storage, error) {
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".secretsafe")
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, err
	}
	return &Storage{filePath: filepath.Join(path, "secrets.json")}, nil
}

// Save stores a secret
func (s *Storage) Save(namespace, key, value string) error {
	secrets, err := s.loadSecrets()
	if err != nil {
		return err
	}

	if secrets[namespace] == nil {
		secrets[namespace] = make(map[string]string)
	}
	secrets[namespace][key] = value

	return s.saveSecrets(secrets)
}

// Load retrieves a secret
func (s *Storage) Load(namespace, key string) (string, error) {
	secrets, err := s.loadSecrets()
	if err != nil {
		return "", err
	}

	if val, exists := secrets[namespace][key]; exists {
		return val, nil
	}

	return "", errors.New("secret not found")
}

// Delete removes a secret
func (s *Storage) Delete(namespace, key string) error {
	secrets, err := s.loadSecrets()
	if err != nil {
		return err
	}

	delete(secrets[namespace], key)
	return s.saveSecrets(secrets)
}

// Helper functions to read/write from the file
func (s *Storage) loadSecrets() (map[string]map[string]string, error) {
	file, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]map[string]string), nil
		}
		return nil, err
	}

	var secrets map[string]map[string]string
	if err := json.Unmarshal(file, &secrets); err != nil {
		return nil, err
	}
	return secrets, nil
}

func (s *Storage) saveSecrets(secrets map[string]map[string]string) error {
	fileData, err := json.MarshalIndent(secrets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, fileData, 0600)
}
