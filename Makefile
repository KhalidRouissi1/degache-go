# Makefile for degache-go

.PHONY: help build test test-verbose test-coverage clean fmt vet lint run-examples install-deps

# Default target
help:
	@echo "Available targets:"
	@echo "  build         - Build the project"
	@echo "  test          - Run tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golint (if available)"
	@echo "  run-examples  - Run example code"
	@echo "  install-deps  - Install development dependencies"

# Build the project
build:
	@echo "Building degache-go..."
	go build ./...

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	go clean ./...
	rm -f coverage.out coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run golint (if available)
lint:
	@echo "Running golint..."
	@if command -v golint >/dev/null 2>&1; then \
		golint ./...; \
	else \
		echo "golint not found. Install with: go install golang.org/x/lint/golint@latest"; \
	fi

# Run example code
run-examples:
	@echo "Running examples..."
	cd examples && go run main.go

# Install development dependencies
install-deps:
	@echo "Installing development dependencies..."
	go mod tidy
	@echo "Installing golint..."
	go install golang.org/x/lint/golint@latest

# Run all checks
check: fmt vet test
	@echo "All checks passed!"

# Initialize go module (for first time setup)
init:
	@echo "Initializing go module..."
	go mod init github.com/degache-go/degache
	go mod tidy

# Update dependencies
update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Benchmark tests
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. ./...

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060
	@echo "Documentation server started at http://localhost:6060"

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build -o bin/degache-linux-amd64 ./examples
	GOOS=darwin GOARCH=amd64 go build -o bin/degache-darwin-amd64 ./examples
	GOOS=windows GOARCH=amd64 go build -o bin/degache-windows-amd64.exe ./examples
	@echo "Binaries built in bin/ directory"

# Create release
release: clean fmt vet test build-all
	@echo "Release build completed!"
