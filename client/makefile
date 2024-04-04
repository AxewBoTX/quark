# Variables
GO := go
AIR := air
MAKE := make
TEMPL := templ
BUN := bun
BINARY_NAME := quark_client
BUILD_DIR := build
SRC_DIR := base_src

dev: 
	@$(AIR) -c .air.toml
air:
	@$(MAKE) generate
	@$(GO) build -o ./tmp/main .

dev-tw:
	@$(BUN) tailwindcss -o ./tailwind/input.css -o ./public/lib/index.css --watch

generate:
	@$(TEMPL) generate

build:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@rm -rf $(BUILD_DIR)
clean-all:
	@rm -rf $(BUILD_DIR)
	@rm -rf $(SRC_DIR)

.PHONY: 
	dev dev-tw generate air build clean clean-all
