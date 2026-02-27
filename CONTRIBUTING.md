# Contributing to balafetch

First off, thanks for taking the time to contribute! Whether it's a bug report, a feature suggestion, or a code contribution, your help is greatly appreciated. This document provides guidelines to help you contribute effectively to the balafetch project.

The following is a set of guidelines for contributing to balafetch. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Your First Code Contribution](#your-first-code-contribution)
- [Development Setup](#development-setup)
- [Development Workflow](#development-workflow)
- [Building the Project](#building-the-project)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Style Guidelines](#style-guidelines)

## Code of Conduct

Be respectful, be kind, and have fun. That's it.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** (commands you ran, error messages, logs from corresponding balafetch.log file of your OS)
- **Describe the behavior you observed** and what you expected
- **Include your environment details:**
  - OS and version
  - Terminal emulator
  - fastfetch version (`fastfetch --version`)
  - balafetch version

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the suggested enhancement
- **Explain why this enhancement would be useful** to most balafetch users
- **List any alternative solutions** you've considered

### Your First Code Contribution?

Unsure where to begin? Look for issues labeled:
- `good first issue` - Simple issues perfect for newcomers
- `help wanted` - Issues where we'd appreciate community help

## Development Setup

### Prerequisites

- **Go 1.21 or higher** - [Install Go](https://go.dev/doc/install)
- **Git** - [Install Git](https://git-scm.com/downloads)
- **fastfetch** - Required for testing. [Install fastfetch](https://github.com/fastfetch-cli/fastfetch)

### Setup Steps

1. **Fork the repository** on GitHub

2. **Clone your fork:**
   ```bash
   git clone https://github.com/YOUR-USERNAME/balafetch.git
   cd balafetch
   ```

3. **Add the upstream repository:**
   ```bash
   git remote add upstream https://github.com/gitmobkab/balafetch.git
   ```

4. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

5. **quickly test balafetch:**
   ```bash
   go run ./cmd/balafetch/
   ./balafetch
   ```


### Package Descriptions

- **`cmd/balafetch`** - Main application entry point
- **`internal/api`** - Handles communication with the Balatro fandom API
- **`internal/imageutil`** - Handles image downloading and cleanup
- **`internal/model`** - Defines data structures used across the project
- **`internal/random`** - Random selection utilities for picking categories and cards
- **`internal/run`** - Executes fastfetch with the downloaded card image

## Development Workflow

1. **Create a new branch** for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following the [style guidelines](#style-guidelines)

3. **Test your changes:**
   ```bash
   go test ./...
   go run ./cmd/balafetch  # Test the actual program
   ```

4. **Update documentation if needed:**
   - Update `README.md` for user-facing changes
   - Update `docs/errors_codes.md` if you added new exit codes
   - Update `docs/troubleshooting.md` if you added new error scenarios

5. **Commit your changes:**
   ```bash
   git add .
   git commit -m "feat: add awesome feature"
   ```
   
   Use conventional commit messages:
   - `feat:` - New feature
   - `fix:` - Bug fix
   - `docs:` - Documentation changes
   - `refactor:` - Code refactoring
   - `test:` - Adding tests
   - `chore:` - Maintenance tasks

6. **Push to your fork:**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Open a Pull Request** on GitHub



## Building the Project

### Development Build

For quick testing during development:
```bash
go run ./cmd/balafetch/
```

**DO NOT** use `go build` for development builds, as it does not include the necessary build-time information (version, commit hash, build time) that is included in the release builds.

**DO NOT** use `compile.sh` or push a commit on `dist/`
as these are reserved for release builds only. Using them for development builds can lead to confusion and potential issues with versioning and distribution.


## Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./internal/random/
```

### Error Handling

- Always handle errors explicitly
- Log errors appropriately (using the project's logging setup)
- Return meaningful error messages
- Use the established exit codes (see `docs/errors_codes.md`)

### Exit Codes

If you need to add a new exit code:
1. Update `docs/errors_codes.md` with the new code, type, reason, category, and fix
2. Update error handling logic in the code
3. Update `docs/troubleshooting.md` if the error needs special troubleshooting steps

## Questions?

Feel free to open an issue with the `question` label if you need help or clarification on anything!

---

**Happy contributing! 🎴🚀**