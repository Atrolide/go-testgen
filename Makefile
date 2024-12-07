# Variables for platform and output paths
GOOS_LINUX = linux
GOARCH = amd64
GOOS_WINDOWS = windows
BUILD_DIR = build
LINUX_DIR = $(BUILD_DIR)/testgen_linux_amd64
WINDOWS_DIR = $(BUILD_DIR)/testgen_windows_amd64
LINUX_BIN = $(LINUX_DIR)/testgen
WINDOWS_BIN = $(WINDOWS_DIR)/testgen.exe
LINUX_CONFIG = $(LINUX_DIR)/config.ini
WINDOWS_CONFIG = $(WINDOWS_DIR)/config.ini

# Default target
all: clean build config

# Build project for both Linux and Windows platforms
build:
	@echo "Building for Linux..."
	@mkdir -p $(LINUX_DIR)  # Create the linux directory if it doesn't exist
	GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH) go build -o $(LINUX_BIN) cmd/testgen/main.go
	@echo "Building for Windows..."
	@mkdir -p $(WINDOWS_DIR)  # Create the windows directory if it doesn't exist
	GOOS=$(GOOS_WINDOWS) GOARCH=$(GOARCH) go build -o $(WINDOWS_BIN) cmd/testgen/main.go

# Generate the INI config file for both Linux and Windows
config:
	@echo "Generating config file for Linux: $(LINUX_CONFIG)"
	@mkdir -p $(LINUX_DIR)  # Ensure the linux directory exists
	@echo "[suffix]" > $(LINUX_CONFIG)
	@echo "suffix = _test" >> $(LINUX_CONFIG)
	@echo "" >> $(LINUX_CONFIG)

	@echo "[go]" >> $(LINUX_CONFIG)
	@echo "extension = .go" >> $(LINUX_CONFIG)
	@echo "description = Go source code files" >> $(LINUX_CONFIG)
	@echo "" >> $(LINUX_CONFIG)

	@echo "[python]" >> $(LINUX_CONFIG)
	@echo "extension = .py" >> $(LINUX_CONFIG)
	@echo "description = Python source code files" >> $(LINUX_CONFIG)
	@echo "" >> $(LINUX_CONFIG)

	@echo "[lua]" >> $(LINUX_CONFIG)
	@echo "extension = .lua" >> $(LINUX_CONFIG)
	@echo "description = Lua source code files" >> $(LINUX_CONFIG)

	@echo "Generating config file for Windows: $(WINDOWS_CONFIG)"
	@mkdir -p $(WINDOWS_DIR)  # Ensure the windows directory exists
	@echo "[suffix]" > $(WINDOWS_CONFIG)
	@echo "suffix = _test" >> $(WINDOWS_CONFIG)
	@echo "" >> $(WINDOWS_CONFIG)

	@echo "[go]" >> $(WINDOWS_CONFIG)
	@echo "extension = .go" >> $(WINDOWS_CONFIG)
	@echo "description = Go source code files" >> $(WINDOWS_CONFIG)
	@echo "" >> $(WINDOWS_CONFIG)

	@echo "[python]" >> $(WINDOWS_CONFIG)
	@echo "extension = .py" >> $(WINDOWS_CONFIG)
	@echo "description = Python source code files" >> $(WINDOWS_CONFIG)
	@echo "" >> $(WINDOWS_CONFIG)

	@echo "[lua]" >> $(WINDOWS_CONFIG)
	@echo "extension = .lua" >> $(WINDOWS_CONFIG)
	@echo "description = Lua source code files" >> $(WINDOWS_CONFIG)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Show help
help:
	@echo "Makefile commands:"
	@echo "  make build        - Build the project for both Linux and Windows"
	@echo "  make config       - Generate the config file for both Linux and Windows"
	@echo "  make clean        - Clean up build artifacts"
	@echo "  make help         - Show this help message"
