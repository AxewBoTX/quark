# Variables
GO := go
BINARY_NAME := quark_server
BUILD_DIR := build
DB_DIR := database

dev: 
	@air -c .air.toml

build:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@rm -rf $(BUILD_DIR)
clean-all:
	@rm -rf $(BUILD_DIR)
	@rm -rf $(DB_DIR)

.PHONY: 
	dev build clean clean-all
