#!/bin/sh

OS=$(uname)
ARCH=$(uname -m)

if [ "$OS" = "Darwin" ]; then
    if [ "$ARCH" = "arm64" ]; then
        BIN="ghosthub-darwin-arm64"
    else
        BIN="ghosthub-darwin-amd64"
    fi
elif [ "$OS" = "Linux" ]; then
    BIN="ghosthub-linux-amd64"
else
    echo "❌ Operating system not supported: $OS"
    exit 1
fi

if [ ! -f "build/$BIN" ]; then
    echo "❌ Binary not found: build/$BIN"
    echo "Run ./build.sh first"
    exit 1
fi

echo "📦 Installing ghosthub..."
sudo cp "build/$BIN" /usr/local/bin/ghosthub
sudo chmod +x /usr/local/bin/ghosthub

echo "✅ Installation completed!"
echo "🚀 Use 'ghosthub' to start"