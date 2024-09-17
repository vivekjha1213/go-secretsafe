// tests/storage_test.go
package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vivekjha1213/go-secretsafe/pkg/secretsafe" // Adjust this path if necessary
)

func TestFileStorage(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "secretsafe-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a new FileStorage instance
	storage, err := secretsafe.NewStorage(tempDir)
	if err != nil {
		t.Fatalf("Failed to create FileStorage: %v", err)
	}

	// Test Set and Get
	t.Run("Set and Get", func(t *testing.T) {
		err := storage.Set("test-namespace", "test-key", "test-value")
		if err != nil {
			t.Fatalf("Failed to set value: %v", err)
		}

		value, err := storage.Get("test-namespace", "test-key")
		if err != nil {
			t.Fatalf("Failed to get value: %v", err)
		}

		if value != "test-value" {
			t.Errorf("Got %q, want %q", value, "test-value")
		}
	})

	// Test Get non-existent key
	t.Run("Get non-existent key", func(t *testing.T) {
		_, err := storage.Get("test-namespace", "non-existent-key")
		if err != secretsafe.ErrKeyNotFound {
			t.Errorf("Expected ErrKeyNotFound when getting non-existent key, but got: %v", err)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		err := storage.Set("test-namespace", "delete-key", "delete-value")
		if err != nil {
			t.Fatalf("Failed to set value: %v", err)
		}

		err = storage.Delete("test-namespace", "delete-key")
		if err != nil {
			t.Fatalf("Failed to delete value: %v", err)
		}

		_, err = storage.Get("test-namespace", "delete-key")
		if err != secretsafe.ErrKeyNotFound {
			t.Errorf("Expected ErrKeyNotFound when getting deleted key, but got: %v", err)
		}
	})

	// Test file persistence
	t.Run("File persistence", func(t *testing.T) {
		err := storage.Set("persist-namespace", "persist-key", "persist-value")
		if err != nil {
			t.Fatalf("Failed to set value: %v", err)
		}

		// Create a new storage instance with the same path
		newStorage, err := secretsafe.NewStorage(tempDir)
		if err != nil {
			t.Fatalf("Failed to create new FileStorage: %v", err)
		}

		value, err := newStorage.Get("persist-namespace", "persist-key")
		if err != nil {
			t.Fatalf("Failed to get persisted value: %v", err)
		}

		if value != "persist-value" {
			t.Errorf("Got %q, want %q", value, "persist-value")
		}
	})

	// Test file permissions
	t.Run("File permissions", func(t *testing.T) {
		err := storage.Set("perm-namespace", "perm-key", "perm-value")
		if err != nil {
			t.Fatalf("Failed to set value: %v", err)
		}

		filePath := filepath.Join(tempDir, "perm-namespace.json")
		info, err := os.Stat(filePath)
		if err != nil {
			t.Fatalf("Failed to get file info: %v", err)
		}

		if info.Mode().Perm() != 0600 {
			t.Errorf("Expected file permissions 0600, got %v", info.Mode().Perm())
		}
	})
}
