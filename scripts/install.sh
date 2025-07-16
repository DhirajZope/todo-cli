#!/usr/bin/env bash
set -euo pipefail

REPO="DhirajZope/todo-cli"
VER=${1:-latest}
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
[[ "$ARCH" == "x86_64" ]] && ARCH=amd64
[[ "$ARCH" =~ arm64|aarch64 ]] && ARCH=arm64

# find the correct asset URL
if [[ "$VER" == "latest" ]]; then
  URL=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" \
    | grep "browser_download_url" \
    | grep "${OS}-${ARCH}.tar.gz" \
    | head -n1 | cut -d '"' -f4)
else
  URL="https://github.com/$REPO/releases/download/$VER/todo_${VER}_${OS}-${ARCH}.tar.gz"
fi

echo "Downloading $URL"
curl -Lo todo.tar.gz "$URL"
tar -xzf todo.tar.gz
chmod +x todo

INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
  echo "Requires sudo to install to $INSTALL_DIR"
  sudo mv todo "$INSTALL_DIR/todo"
else
  mv todo "$INSTALL_DIR/todo"
fi

echo "✅ Installed todo to $INSTALL_DIR/todo"
