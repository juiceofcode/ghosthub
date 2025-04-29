#!/bin/sh

rm -rf build/

mkdir -p build/

# Linux
GOOS=linux GOARCH=amd64 go build -o build/ghosthub-linux-amd64

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o build/ghosthub-darwin-amd64

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o build/ghosthub-darwin-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o build/ghosthub-windows-amd64.exe

echo "✅ Builds completed!" 