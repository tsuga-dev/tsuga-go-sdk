.PHONY: test fmt build clean

# Default target
all: fmt test build

# Format go code
fmt:
	go fmt ./...

# Run tests
test:
	go test -v ./...

# Build the SDK (verify it compiles)
build:
	go build -v ./...

# Clean up
clean:
	go clean
