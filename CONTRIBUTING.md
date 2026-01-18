# Contributing to balafetch

First off, thanks for taking the time to contribute! 🎴

The following is a set of guidelines for contributing to balafetch. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Your First Code Contribution](#your-first-code-contribution)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
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
- **Provide specific examples** (commands you ran, error messages, logs from `~/balafetch/balafetch.log`)
- **Describe the behavior you observed** and what you expected
- **Include your environment details:**
  - OS and version
  - Terminal emulator
  - fastfetch version (`fastfetch --version`)
  - balafetch version

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the suggested enhancement
- **Explain why this enhancement would be useful** to most balafetch users
- **List any alternative solutions** you've considered

### Your First Code Contribution

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
   go mod download
   ```

5. **Build and test:**
   ```bash
   go build ./cmd/balafetch/
   ./balafetch
   ```

## Project Structure

```
.
├── cmd/
│   └── balafetch/
│       └── main.go           # Entry point
├── internal/
│   ├── api/                  # Balatro fandom API client
│   │   ├── balatro.go        # API functions
│   │   ├── get_request.go    # HTTP request logic
│   │   └── params.go         # API parameters
│   ├── imageutil/            # Image download utilities
│   │   └── imageutil.go
│   ├── model/                # Data structures
│   │   └── models.go
│   ├── random/               # Random selection logic
│   │   ├── pick.go
│   │   └── pick_test.go
│   └── run/                  # fastfetch execution
│       ├── fastfetch.go
│       └── run.go
├── docs/                     # Documentation
├── imgs/                     # Images for README
├── dist/                     # Compiled binaries (generated)
├── compile.sh               # Build script for all platforms
├── go.mod                   # Go module definition
├── LICENSE
└── README.md
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
   ./balafetch  # Test the actual program
   ```

4. **If you modified code, run the build script:**
   ```bash
   ./compile.sh
   ```
   > [!IMPORTANT]
   > Since we don't have automated releases yet, you **must** run `compile.sh` after code changes to rebuild all platform binaries in `dist/`.

5. **Update documentation if needed:**
   - Update `README.md` for user-facing changes
   - Update `docs/errors_codes.md` if you added new exit codes
   - Update `docs/troubleshooting.md` if you added new error scenarios

6. **Commit your changes:**
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

7. **Push to your fork:**
   ```bash
   git push origin feature/your-feature-name
   ```

8. **Open a Pull Request** on GitHub

## Building the Project

### Development Build

For quick testing during development:
```bash
go build ./cmd/balafetch/ -o balafetch
```

### Cross-Platform Build

**Important:** Before submitting a PR with code changes, you must build for all platforms:

```bash
./compile.sh
```

This script builds balafetch for all supported platforms and places binaries in the `dist/` directory:
- Linux (amd64, arm64, arm)
- macOS/Darwin (amd64, arm64)
- Windows (amd64, arm64)
- FreeBSD (amd64, arm64, arm)

> [!NOTE]
> We don't have GitHub Actions for automated releases yet, so maintainers rely on `compile.sh` for building release binaries.

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

### Manual Testing

After making changes, test the actual program:

1. **Test successful execution:**
   ```bash
   ./balafetch
   ```
   Should display fastfetch with a random Balatro card.

2. **Test error handling:**
   - Disconnect from internet and run `./balafetch` (should fallback gracefully)
   - Check logs at `~/balafetch/balafetch.log`

3. **Test on your platform** (especially if modifying platform-specific code)

### Adding Tests

If you add new functions, especially in `internal/`:
- Add corresponding test files (e.g., `myfunction_test.go`)
- Aim for meaningful test coverage
- Test edge cases and error conditions

## Submitting Changes

### Pull Request Process

1. **Update documentation** - Make sure README, docs/, and code comments are up to date
2. **Run compile.sh** - Ensure all platform binaries build successfully
3. **Test thoroughly** - Run tests and manually test the program
4. **Create a descriptive PR:**
   - Clear title explaining what the PR does
   - Description of changes made
   - Reference any related issues
   - Include before/after examples if applicable

### PR Checklist

Before submitting your PR, verify:

- [ ] Code follows Go style guidelines
- [ ] All tests pass (`go test ./...`)
- [ ] `compile.sh` runs successfully (if code was modified)
- [ ] Documentation is updated (README, docs/, code comments)
- [ ] Commit messages follow conventional commit format
- [ ] No unnecessary files are included (e.g., personal configs, temp files)

## Style Guidelines

### Go Code Style

- Follow standard Go conventions and formatting
- Run `go fmt` on all code:
  ```bash
  go fmt ./...
  ```
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and concise

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

### Documentation Style

- Use clear, concise language
- Include code examples where helpful
- Update all relevant documentation when making changes
- Check for typos and grammatical errors

## Questions?

Feel free to open an issue with the `question` label if you need help or clarification on anything!

---

**Happy contributing! 🎴🚀**