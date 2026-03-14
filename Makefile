BINARY_NAME := file_renamer

BUILD_DIR := bin

.DEFAULT_GOAL := build

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete. Executable located at $(BUILD_DIR)/$(BINARY_NAME)"

run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# TODO: Add Tests
# test:
# 	@echo "Running tests..."
# 	@go test ./...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

.PHONY: build run clean
