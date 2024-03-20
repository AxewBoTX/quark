# Variables
GO := go
BINARY_NAME := quark_server
BUILD_DIR := build
SRC_DIR := base_src

dev: 
	@air -c .air.toml

build:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@rm -rf $(BUILD_DIR)
clean-all:
	@rm -rf $(BUILD_DIR)
	@rm -rf $(SRC_DIR)

.PHONY: 
	dev build clean clean-all
