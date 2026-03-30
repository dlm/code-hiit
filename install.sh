#!/bin/sh
set -e

# code-hiit installer
# Usage: curl -fsSL https://raw.githubusercontent.com/dlm/code-hiit/main/install.sh | sh
#
# Environment variables:
#   INSTALL_DIR - Directory to install binary (default: auto-detect)
#
# Examples:
#   curl -fsSL ... | sh                           # Auto-detect location
#   curl -fsSL ... | INSTALL_DIR=~/.local/bin sh  # Custom directory
#   curl -fsSL ... | INSTALL_DIR=~/bin sh         # Another custom location

REPO="dlm/code-hiit"
BINARY_NAME="code-hiit"

# Determine install directory (default to user's local bin)
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    printf "${GREEN}==>${NC} %s\n" "$1"
}

log_error() {
    printf "${RED}Error:${NC} %s\n" "$1" >&2
}

log_warn() {
    printf "${YELLOW}Warning:${NC} %s\n" "$1"
}

# Detect OS
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "darwin";;
        *)
            log_error "Unsupported operating system: $(uname -s)"
            log_error "Supported: Linux, macOS"
            exit 1
            ;;
    esac
}

# Detect architecture
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64)   echo "amd64";;
        aarch64|arm64)  echo "arm64";;
        *)
            log_error "Unsupported architecture: $(uname -m)"
            log_error "Supported: x86_64/amd64, aarch64/arm64"
            exit 1
            ;;
    esac
}

# Get latest release version (including pre-releases)
get_latest_version() {
    curl -fsSL "https://api.github.com/repos/${REPO}/releases" | \
        grep '"tag_name":' | \
        head -1 | \
        sed -E 's/.*"([^"]+)".*/\1/'
}

# Download and install
main() {
    log_info "Installing ${BINARY_NAME}..."

    # Detect platform
    OS=$(detect_os)
    ARCH=$(detect_arch)
    log_info "Detected platform: ${OS}/${ARCH}"

    # Validate platform combination
    if [ "$OS" = "linux" ] && [ "$ARCH" != "amd64" ]; then
        log_error "Only Linux AMD64 (x86_64) is supported"
        log_error "Your system: ${OS}/${ARCH}"
        exit 1
    fi
    if [ "$OS" = "darwin" ] && [ "$ARCH" != "arm64" ]; then
        log_error "Only macOS Apple Silicon (ARM64) is supported"
        log_error "Your system: ${OS}/${ARCH}"
        exit 1
    fi

    # Get latest version
    log_info "Fetching latest release..."
    VERSION=$(get_latest_version)
    if [ -z "$VERSION" ]; then
        log_error "Failed to fetch latest release version"
        exit 1
    fi
    log_info "Latest version: ${VERSION}"

    # Construct download URL
    BINARY_FILE="${BINARY_NAME}-${OS}-${ARCH}"
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY_FILE}"

    # Create temporary directory
    TMP_DIR=$(mktemp -d)
    trap "rm -rf ${TMP_DIR}" EXIT

    # Download binary
    log_info "Downloading from ${DOWNLOAD_URL}..."
    if ! curl -fsSL -o "${TMP_DIR}/${BINARY_NAME}" "${DOWNLOAD_URL}"; then
        log_error "Failed to download binary"
        exit 1
    fi

    # Make executable
    chmod +x "${TMP_DIR}/${BINARY_NAME}"

    # Verify binary works
    if ! "${TMP_DIR}/${BINARY_NAME}" --version >/dev/null 2>&1; then
        log_error "Downloaded binary failed to execute"
        exit 1
    fi

    # Ensure install directory exists
    if [ ! -d "$INSTALL_DIR" ]; then
        log_info "Creating directory: ${INSTALL_DIR}"
        if mkdir -p "$INSTALL_DIR" 2>/dev/null; then
            : # Success
        elif command -v sudo >/dev/null 2>&1; then
            log_warn "Need sudo to create ${INSTALL_DIR}"
            sudo mkdir -p "$INSTALL_DIR"
        else
            log_error "Cannot create ${INSTALL_DIR} (no write permission and sudo not available)"
            exit 1
        fi
    fi

    # Install binary
    log_info "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    elif command -v sudo >/dev/null 2>&1; then
        log_warn "Need sudo to write to ${INSTALL_DIR}"
        sudo mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        log_error "Cannot write to ${INSTALL_DIR} (no write permission and sudo not available)"
        exit 1
    fi

    log_info "✓ Installed to ${INSTALL_DIR}/${BINARY_NAME}"

    # Check if directory is in PATH
    case ":$PATH:" in
        *":$INSTALL_DIR:"*)
            log_info ""
            log_info "Run '${BINARY_NAME}' to start!"
            ;;
        *)
            log_info ""
            log_warn "Note: ${INSTALL_DIR} is not in your PATH"
            log_warn "Add this line to your shell rc file (~/.bashrc, ~/.zshrc, etc.):"
            log_warn "  export PATH=\"${INSTALL_DIR}:\${PATH}\""
            log_info ""
            log_info "Or run directly: ${INSTALL_DIR}/${BINARY_NAME}"
            ;;
    esac
}

main
