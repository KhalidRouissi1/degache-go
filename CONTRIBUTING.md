# Contributing to degache-go

We welcome contributions from the Tunisian developer community and beyond! Whether it's bug fixes, new features, documentation improvements, or test cases, your contributions are valuable.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for using Makefile commands)

### Setting up the development environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/degache-go.git
   cd degache-go
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run tests to ensure everything works:
   ```bash
   make test
   # or
   go test ./...
   ```

## 🛠️ Development Workflow

### Making Changes

1. Create a new branch for your feature/fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the coding standards below

3. Add tests for your changes:
   ```bash
   # Add tests in appropriate *_test.go files
   go test ./... -v
   ```

4. Run all checks:
   ```bash
   make check
   # This runs: fmt, vet, and test
   ```

5. Commit your changes:
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

6. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

7. Create a Pull Request

### Coding Standards

- **Go formatting**: Use `go fmt` or `make fmt`
- **Go vet**: Run `go vet` or `make vet` to catch common issues
- **Comments**: Add meaningful comments for exported functions
- **Error handling**: Always handle errors appropriately
- **Tests**: Write tests for new functionality
- **Documentation**: Update README.md if needed

### Commit Message Format

We follow conventional commits:

- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions/changes
- `refactor:` for code refactoring
- `perf:` for performance improvements

Example: `feat: add postal code validation for new regions`

## 🧪 Testing

### Running Tests

```bash
# Run all tests
make test
# or
go test ./...

# Run tests with verbose output
make test-verbose
# or
go test -v ./...

# Run tests with coverage
make test-coverage
# or
go test -cover ./...
```

### Writing Tests

- Place tests in `*_test.go` files in the same package
- Use table-driven tests when appropriate
- Test both success and failure cases
- Include edge cases
- Add benchmarks for performance-critical functions

Example test structure:
```go
func TestValidateExample(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {"valid input", "valid", true},
        {"invalid input", "invalid", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ValidateExample(tt.input)
            if result != tt.expected {
                t.Errorf("ValidateExample(%q) = %v, want %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## 📝 Documentation

### Code Documentation

- All exported functions must have comments
- Comments should explain what the function does, its parameters, and return values
- Include usage examples in comments when helpful

Example:
```go
// ValidateExample validates an example input
//
// Parameters:
//   - input: The input string to validate
//
// Returns:
//   - bool: true if valid, false otherwise
//
// Example:
//   isValid := ValidateExample("test") // returns true
func ValidateExample(input string) bool {
    // implementation
}
```

### README Updates

If your changes affect the public API or add new features:
- Update the README.md with new examples
- Update the feature list
- Add any new installation or usage instructions

## 🐛 Bug Reports

When reporting bugs, please include:

1. **Description**: Clear description of the issue
2. **Steps to reproduce**: Minimal code example
3. **Expected behavior**: What should happen
4. **Actual behavior**: What actually happens
5. **Environment**: Go version, OS, etc.

## ✨ Feature Requests

For new features:

1. **Use case**: Describe the problem you're trying to solve
2. **Proposed solution**: How you think it should work
3. **Alternatives**: Other solutions you've considered
4. **Tunisian context**: How it relates to Tunisian data/standards

## 🔍 Code Review Process

1. All submissions require review
2. We may ask for changes before merging
3. Please be patient and responsive to feedback
4. Maintain a respectful and constructive tone

## 📋 Checklist for Pull Requests

Before submitting a PR, ensure:

- [ ] Code follows Go conventions and formatting
- [ ] All tests pass
- [ ] New functionality has tests
- [ ] Documentation is updated if needed
- [ ] Commit messages follow conventional format
- [ ] No breaking changes (or clearly documented)
- [ ] Examples work correctly

## 🌍 Tunisian Context

This library is specifically designed for Tunisian developers and data. When contributing:

- Ensure accuracy of Tunisian-specific data (banks, carriers, etc.)
- Consider regional variations within Tunisia
- Respect Arabic language formatting and conventions
- Test with real Tunisian data when possible

## 📞 Getting Help

- Open an issue for questions
- Check existing issues and documentation first
- Be specific about your problem or question

## 🙏 Recognition

Contributors will be recognized in:
- CHANGELOG.md for their contributions
- README.md contributors section (if significant contributions)

Thank you for contributing to degache-go! 🇹🇳
