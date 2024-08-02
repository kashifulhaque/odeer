if ! command -v go &> /dev/null
then
  echo "Go is not installed. Please install Go and try again."
  exit 1
fi

OS_NAME=$(uname -s | tr '[:upper:]' '[:lower:]')
OS_ARCH=$(uname -m)

if [ "$OS_ARCH" = "x86_64" ]; then
  OS_ARCH="amd64"
elif [ "$OS_ARCH" = "aarch64" ]; then
  OS_ARCH="arm64"
fi

# Build the odeer executable
GOOS=$OS_NAME GOARCH=$OS_ARCH go build -ldflags "-X odeer/internal/config.accountID=$CLOUDFLARE_ACCOUNT_ID -X odeer/internal/config.authToken=$CLOUDFLARE_WORKERS_AI_API_KEY" -o odeer cmd/odeer/main.go
