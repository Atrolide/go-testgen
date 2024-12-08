# Variables for platform and output paths
GOOS_LINUX = linux
GOARCH = amd64
GOOS_WINDOWS = windows
BUILD_DIR = build
LINUX_DIR = $(BUILD_DIR)/testgen_linux_amd64
WINDOWS_DIR = $(BUILD_DIR)/testgen_windows_amd64
LINUX_BIN = $(LINUX_DIR)/testgen
WINDOWS_BIN = $(WINDOWS_DIR)/testgen.exe

# Default target
all: clean build 

# Build project for both Linux and Windows platforms
build:
	@echo "Building for Linux..."
	@mkdir -p $(LINUX_DIR)  # Create the linux directory if it doesn't exist
	GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH) go build -o $(LINUX_BIN) cmd/testgen/main.go
	@echo "Building for Windows..."
	@mkdir -p $(WINDOWS_DIR)  # Create the windows directory if it doesn't exist
	GOOS=$(GOOS_WINDOWS) GOARCH=$(GOARCH) go build -o $(WINDOWS_BIN) cmd/testgen/main.go

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Show help
help:
	@echo "Makefile commands:"
	@echo "  make build        - Build the project for both Linux and Windows"
	@echo "  make clean        - Clean up build artifacts"
	@echo "  make help         - Show this help message"
