# Platform-specific build variables
GOOS_LINUX      := linux
GOOS_WINDOWS    := windows
GOARCH          := amd64

# Directory paths for build output and archives
BUILD_DIR       := build
BIN_DIR         := $(BUILD_DIR)/bin
ARCHIVE_DIR     := $(BUILD_DIR)/archives
LINUX_BIN_DIR   := $(BIN_DIR)/testgen_linux_amd64
WINDOWS_BIN_DIR := $(BIN_DIR)/testgen_windows_amd64

# Binary output paths
LINUX_BIN       := $(LINUX_BIN_DIR)/testgen
WINDOWS_BIN     := $(WINDOWS_BIN_DIR)/testgen.exe

# Version and Build information (from VERSION file and git)
VERSION_FILE    := VERSION
VERSION         := $(shell cat $(VERSION_FILE))
BUILD           := $(shell git rev-parse --short HEAD)

# Default target: clean, build, and zip the binaries
all: clean build zip

# Build the project for both Linux and Windows
build: $(LINUX_BIN) $(WINDOWS_BIN)

$(LINUX_BIN):
	@echo "Building for Linux..."
	@mkdir -p $(LINUX_BIN_DIR)  # Ensure the Linux build directory exists
	GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH) go build -o $(LINUX_BIN) \
		-ldflags "-X 'github.com/Atrolide/go-testgen/pkg/version.Version=$(VERSION)' -X 'github.com/Atrolide/go-testgen/pkg/version.Build=$(BUILD)'" \
		cmd/testgen/main.go

$(WINDOWS_BIN):
	@echo "Building for Windows..."
	@mkdir -p $(WINDOWS_BIN_DIR)  # Ensure the Windows build directory exists
	GOOS=$(GOOS_WINDOWS) GOARCH=$(GOARCH) go build -o $(WINDOWS_BIN) \
		-ldflags "-X 'github.com/Atrolide/go-testgen/pkg/version.Version=$(VERSION)' -X 'github.com/Atrolide/go-testgen/pkg/version.Build=$(BUILD)'" \
		cmd/testgen/main.go

# Package the binaries into zip archives
zip: $(ARCHIVE_DIR)/testgen_$(VERSION)_linux_amd64.zip $(ARCHIVE_DIR)/testgen_$(VERSION)_windows_amd64.zip

$(ARCHIVE_DIR)/testgen_$(VERSION)_linux_amd64.zip: $(LINUX_BIN)
	@echo "Creating zip archive for Linux..."
	@mkdir -p $(ARCHIVE_DIR)  # Ensure the archive directory exists
	@zip $@ -j $(LINUX_BIN)

$(ARCHIVE_DIR)/testgen_$(VERSION)_windows_amd64.zip: $(WINDOWS_BIN)
	@echo "Creating zip archive for Windows..."
	@mkdir -p $(ARCHIVE_DIR)  # Ensure the archive directory exists
	@zip $@ -j $(WINDOWS_BIN)

# Clean up the build artifacts and directories
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)

# Display available Makefile commands and their descriptions
help:
	@echo "Makefile Commands:"
	@echo "  make build        - Build the project for both Linux and Windows"
	@echo "  make clean        - Clean up build artifacts"
	@echo "  make zip          - Create zip archives for both Linux and Windows binaries"
	@echo "  make help         - Show this help message"
