#!/bin/bash

set -e

REPO="MichelDiz/tree2"
VERSION="${1:-latest}"

# Detect OS and ARCH
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize arch names
if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
  ARCH="arm64"
fi

# Detect version
if [ "$VERSION" = "latest" ]; then
  VERSION=$(curl --silent "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
fi

echo "Installing tree2 version $VERSION for $OS/$ARCH..."

# Build download URL
BINARY_URL="https://github.com/$REPO/releases/download/$VERSION/tree2-${OS}-${ARCH}"

# Download binary
curl -L -o tree2 "$BINARY_URL"
chmod +x tree2

# Install
sudo mv tree2 /usr/local/bin/tree2

echo "tree2 installed successfully!"
tree2 --help
