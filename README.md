# go-secretsafe

![go-secretsafe Logo](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSeWcxkvj2f16XpcwuDGxwnUqnB7miVZykFiA&s)

[![Go Version](https://img.shields.io/github/go-mod/go-version/vivekjha1213/go-secretsafe)](https://github.com/vivekjha1213/go-secretsafe)
[![Build Status](https://img.shields.io/github/actions/workflow/status/vivekjha1213/go-secretsafe/build.yml)](https://github.com/vivekjha1213/go-secretsafe/actions)
[![License](https://img.shields.io/github/license/vivekjha1213/go-secretsafe)](https://github.com/vivekjha1213/go-secretsafe/blob/main/LICENSE)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://github.com/vivekjha1213/go-secretsafe/blob/main/CONTRIBUTING.md)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/vivekjha1213/go-secretsafe/pulls)
[![Issues Open](https://img.shields.io/github/issues/vivekjha1213/go-secretsafe)](https://github.com/vivekjha1213/go-secretsafe/issues)
[![Forks](https://img.shields.io/github/forks/vivekjha1213/go-secretsafe)](https://github.com/vivekjha1213/go-secretsafe/network/members)
[![Stars](https://img.shields.io/github/stars/vivekjha1213/go-secretsafe)](https://github.com/vivekjha1213/go-secretsafe/stargazers)
[![Downloads](https://img.shields.io/github/downloads/vivekjha1213/go-secretsafe/total)](https://github.com/vivekjha1213/go-secretsafe/releases)

`go-secretsafe` is a robust Go package for managing secrets securely. It provides encryption, storage, caching, and versioning capabilities for handling sensitive information in your applications.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Architecture Diagram](#architecture-diagram)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Reference](#api-reference)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Features

- Secure encryption and decryption of secrets
- Local file storage for secrets (encrypted at rest)
- In-memory caching for fast retrieval of frequently used secrets
- Versioning support for tracking changes to secrets
- Custom error types for better error handling and debugging
- Simple and intuitive API for managing secrets
- Configurable options for encryption algorithms, storage locations, and cache settings

## Project Structure

```
go-secretsafe/
├── cmd/
│   └── secretsafe/
│       └── main.go         # Main application entry point
├── pkg/
│   └── secretsafe/
│       ├── cache.go        # In-memory cache implementation
│       ├── encryption.go   # Encryption and decryption logic
│       ├── manager.go      # Secret management operations
│       ├── storage.go      # Local file storage for secrets
│       ├── utils/          # Utility functions and helpers
│       └── versioning.go   # Version control for secrets
├── tests/                  # Test files
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```

## Architecture Diagram

![Screenshot 2024-09-16 at 1 50 19 AM](https://github.com/user-attachments/assets/25c1d84d-1a46-45c4-934d-64c879dd7921)


## Installation

To install `go-secretsafe`, use the following command:

```bash
go get github.com/vivekjha1213/go-secretsafe
```

## Usage

Here's a basic example of how to use `go-secretsafe`:

```go
package main

import (
    "fmt"
    "github.com/vivekjha1213/go-secretsafe/pkg/secretsafe"
)

func main() {
    manager, err := secretsafe.NewManager()
    if err != nil {
        fmt.Printf("Error creating manager: %v\n", err)
        return
    }

    // Set a secret
    err = manager.SetSecret("database", "password", "mysecretpassword")
    if err != nil {
        fmt.Printf("Error setting secret: %v\n", err)
        return
    }

    // Get a secret
    password, err := manager.GetSecret("database", "password")
    if err != nil {
        fmt.Printf("Error getting secret: %v\n", err)
        return
    }

    fmt.Printf("Retrieved password: %s\n", password)

    // Delete a secret
    err = manager.DeleteSecret("database", "password")
    if err != nil {
        fmt.Printf("Error deleting secret: %v\n", err)
        return
    }

    fmt.Println("Secret deleted successfully")
}
```

## Configuration

`go-secretsafe` can be configured with custom options:

```go
manager, err := secretsafe.NewManager(
    secretsafe.WithEncryptionAlgorithm("aes256"),
    secretsafe.WithStoragePath("/path/to/secrets"),
    secretsafe.WithCacheSize(100),
)
```

## API Reference

### `NewManager(options ...Option) (*Manager, error)`
Creates a new instance of Manager with optional configuration.

### `(m *Manager) SetSecret(namespace, key, value string) error`
Sets a secret value for the given namespace and key.

### `(m *Manager) GetSecret(namespace, key string) (string, error)`
Retrieves a secret value for the given namespace and key.

### `(m *Manager) DeleteSecret(namespace, key string) error`
Deletes a secret value for the given namespace and key.

### `(m *Manager) ListSecrets(namespace string) ([]string, error)`
Lists all keys in the given namespace.

### `(m *Manager) SecretExists(namespace, key string) bool`
Checks if a secret exists for the given namespace and key.

## Testing

To run the tests for `go-secretsafe`, use the following command:

```bash
go test ./...
```

## Contributing

We welcome contributions from the community. To contribute:

1. Fork this repository
2. Create a new branch for your feature or bug fix
3. Commit your changes
4. Push to your fork
5. Open a pull request

Please make sure to write appropriate unit tests and follow the existing code style.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

For any questions, issues, or support requests, please:

1. Check the [documentation]([https://docs.google.com/document/d/19YywMvbiWT63gN8pEeRd7dJPgDey3JAA3hlryC4u0hc/edit#heading=h.h73squfqw4e7](https://docs.google.com/document/d/19YywMvbiWT63gN8pEeRd7dJPgDey3JAA3hlryC4u0hc/edit?usp=sharing))
2. Search for existing [issues](https://github.com/vivekjha1213/go-secretsafe/issues)
3. Open a new [issue](https://github.com/vivekjha1213/go-secretsafe/issues/new) if needed
4. Contact [vivekjha1213@gmail.com](mailto:vivekjha1213@gmail.com) for further assistance

---

Thank you for using go-secretsafe! We hope it helps you manage your secrets securely and efficiently.
