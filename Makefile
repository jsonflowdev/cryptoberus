# Build for the current OS and architecture
build:
	go build -ldflags="-s -w" -o ./build/cryptoberus ./cmd/cryptoberus-binance

# Build for all supported platforms
build-all: build-linux build-windows build-mac

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./build/trader-linux ./cmd/cryptoberus-binance

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./build/trader-windows.exe ./cmd/cryptoberus-binance

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./build/trader-mac ./cmd/cryptoberus-binance

run:
	@./build/cryptoberus

test:
	@go test -v -cover ./...

clean:
	rm -rf build/*

.PHONY: build build-all build-linux build-windows build-mac run test clean