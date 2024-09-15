Here’s a sample open-source documentation for your `go-secretsafe` project. This will help users understand how to use the tool and contribute to it.

### `README.md`

```markdown
# go-secretsafe

`go-secretsafe` is a lightweight, command-line tool for managing secrets securely. It provides a simple and effective way to store, retrieve, and delete secrets using encryption. It supports in-memory caching to improve performance.

## Features

- **CLI Commands**: Set, get, and delete secrets from the command line.
- **Encryption**: Securely encrypt and decrypt secrets.
- **Caching**: Optionally cache secrets in memory for faster access.
- **Versioning**: Manage different versions of secrets (planned for future versions).

## Installation

To install `go-secretsafe`, clone the repository and build the project:

```bash
git clone https://github.com/vivekjha1213/go-secretsafe.git
cd go-secretsafe
go build ./cmd/secretsafe
```

Alternatively, you can use the following command to install it directly:

```bash
go install github.com/vivekjha1213/go-secretsafe/cmd/secretsafe@latest
```

## Usage

### Set a Secret

To store a secret, use the `set` command:

```bash
secretsafe set <namespace> <key> <value>
```

**Example:**

```bash
secretsafe set dev db_password supersecretpassword
```

### Get a Secret

To retrieve a stored secret, use the `get` command:

```bash
secretsafe get <namespace> <key>
```

**Example:**

```bash
secretsafe get dev db_password
```

### Delete a Secret

To delete a stored secret, use the `delete` command:

```bash
secretsafe delete <namespace> <key>
```

**Example:**

```bash
secretsafe delete dev db_password
```

## Testing

To run the tests for `go-secretsafe`, use the following command:

```bash
go test ./pkg/secretsafe
```

Make sure you have the `testify` package installed:

```bash
go get github.com/stretchr/testify
```

## Contributing

Contributions are welcome! If you have any ideas for improvements or new features, please open an issue or submit a pull request. Follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Commit your changes and push to your fork.
5. Open a pull request against the main repository.

## License

`go-secretsafe` is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, you can reach out to Vivek Kumar Jha at [vivekjha1213@gmail.com](mailto:vivekjha1213@gmail.com).

---

Thank you for using `go-secretsafe`!
```

### Notes

- **Installation Instructions**: Provides both manual build and direct install methods.
- **Usage Examples**: Clear instructions on how to use each CLI command.
- **Testing**: Instructions for running tests and installing dependencies.
- **Contributing**: Guidelines for contributing to the project.
- **License**: Information about the project's license.
- **Contact**: Provides a way for users to get in touch for support.

Feel free to adjust this documentation according to the specific details and needs of your project.