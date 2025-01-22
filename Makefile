APP_NAME := manaegement system
SRC_DIR := ./cmd
BUILD_DIR := ./build
BIN_DIR := $(BUILD_DIR)/bin
SRC_TEST := ./database

GO := go
GO_RUN := $(GO) run
GO_BUILD := $(GO) build
GO_CLEAN := $(GO) clean
GO_TEST := $(GO) test
GO_MOD := $(GO) mod

.PHONY: all build clean test fmt

all: build

build:
	@mkdir -p $(BIN_DIR)
	$(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)/main.go
run:
	$(GO_MOD) tidy
	$(GO_RUN) $(SRC_DIR)/main.go
clean: 
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR)
