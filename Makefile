# constants
ENTORY_POINT=./main.go
BUILD_DIR=./build/

# Build source
.PHONY: build
build: clean-built
	@echo "> building a go application..."
	test -e $(BUILD_DIR) || mkdir $(BUILD_DIR)
	go build -o $(BUILD_DIR)/show-config $(ENTORY_POINT)

# Remove files under build directory
.PHONY: clean-built
clean-built:
	@echo "> cleaning binaries under $(BUILD_DIR)..."
	test ! -e $(BUILD_DIR) || rm -rf $(BUILD_DIR)

# Format go files
.PHONY: fmt
fmt:
	@echo "> formatting go files..."
	go fmt ./...
