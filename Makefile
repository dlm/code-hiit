# code-hiit Makefile
# Builds binaries for multiple platforms

# Get version from git tag, default to dev
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

# Binary name
BINARY := code-hiit

# Build directories
DIST_DIR := dist

# Platforms to build for
PLATFORMS := \
	linux/amd64 \
	darwin/arm64

.PHONY: all clean build test linux/amd64 darwin/arm64 install

all: clean build

# Build for all platforms
build: $(PLATFORMS)

# Build for specific platforms
linux/amd64:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY)-linux-amd64

darwin/arm64:
	@echo "Building for macOS (Apple Silicon)..."
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY)-darwin-arm64

# Build for current platform only
local:
	@echo "Building for current platform..."
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY)

# Install to /usr/local/bin
install: local
	@echo "Installing to /usr/local/bin..."
	sudo cp $(BINARY) /usr/local/bin/$(BINARY)
	@echo "Installed successfully!"

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(DIST_DIR)
	@rm -f $(BINARY)
	@echo "Clean complete!"

# Display version
version:
	@echo $(VERSION)

# Help
help:
	@echo "code-hiit build system"
	@echo ""
	@echo "Targets:"
	@echo "  all              - Clean and build for all platforms"
	@echo "  build            - Build for all platforms (Linux AMD64, macOS ARM64)"
	@echo "  local            - Build for current platform only"
	@echo "  linux/amd64      - Build for Linux (AMD64)"
	@echo "  darwin/arm64     - Build for macOS (Apple Silicon)"
	@echo "  install          - Build and install to /usr/local/bin"
	@echo "  test             - Run tests"
	@echo "  clean            - Remove build artifacts"
	@echo "  version          - Display version"
	@echo "  help             - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  VERSION          - Set build version (default: git describe or 'dev')"
