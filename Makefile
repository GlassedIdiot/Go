
# Variables
GO=go
BINARY_NAME=rango
SRC_DIR=./src

# Default target
all: build

# Build the binary
build:
	$(GO) build main.go -o $(BINARY_NAME).exe

# Run the binary
run: 
	$(GO) run main.go

# Clean the build
clean:
	rm -f $(BINARY_NAME)

# Format the code
fmt:
	$(GO) fmt $(SRC_DIR)/...

# Lint the code
lint:
	golangci-lint run $(SRC_DIR)/...

# Test the code
test:
	$(GO) test $(SRC_DIR)/...

# Phony targets
.PHONY: all build run clean fmt lint test