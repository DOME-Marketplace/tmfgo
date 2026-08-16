# Makefile for isbetmf

# Binary names
BINARY_NAME=isbetmf
REPORTING_BINARY_NAME=tmf-reporting

# Build directories
BIN_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt

# Main entry points
MAIN_SRC=main.go
REPORTING_SRC=cmd/reporting/main.go

.PHONY: all build build-reporting test clean run help fmt lint vet

all: build build-reporting

# Build the main server binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_SRC)

# Build the reporting tool binary
build-reporting:
	@echo "Building $(REPORTING_BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(REPORTING_BINARY_NAME) $(REPORTING_SRC)

# Run the main server locally (default profile)
run:
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run $(MAIN_SRC) -run isbedev

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Vet code
vet:
	@echo "Vetting code..."
	$(GOVET) ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Run 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest' to install."; \
		exit 1; \
	fi

# Help target
help:
	@echo "Available targets:"
	@echo "  all              - Build both the server and reporting tool"
	@echo "  build            - Build the main server binary ($(BINARY_NAME))"
	@echo "  build-reporting  - Build the reporting tool binary ($(REPORTING_BINARY_NAME))"
	@echo "  run              - Run the server locally with default profile"
	@echo "  test             - Run all tests"
	@echo "  fmt              - Format code"
	@echo "  vet              - Run go vet"
	@echo "  lint             - Run golangci-lint (if installed)"
