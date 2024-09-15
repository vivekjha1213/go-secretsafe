package secretsafe

import (
    "fmt"
    "time"
)

// Versioning handles versioning of secrets
type Versioning struct {
    Namespace string
    Key       string
    Version   string
    Timestamp time.Time
}

// NewVersioning creates a new version entry
func NewVersioning(namespace, key, version string) *Versioning {
    return &Versioning{
        Namespace: namespace,
        Key:       key,
        Version:   version,
        Timestamp: time.Now(),
    }
}

// String returns a string representation of the versioning entry
func (v *Versioning) String() string {
    return fmt.Sprintf("Namespace: %s, Key: %s, Version: %s, Timestamp: %s",
        v.Namespace, v.Key, v.Version, v.Timestamp.Format(time.RFC3339))
}
