DEFAULT_GOAL := server

ROOT_DIR := $(CURDIR)/../..
BIN_DIR := $(ROOT_DIR)/bin
CONFIG_DIR := $(ROOT_DIR)/config
TMP_DIR := $(ROOT_DIR)/tmp

# server - displays the server help.
.PHONY: server
server: server/help

# server/build - builds the server server.
.PHONY: server/build
server/build:
	@echo "Building server..."
	go build -o $(BIN_DIR)/server ./main.go

# server/build_dev - builds the server for development.
.PHONY: server/build_dev
server/build_dev:
	@echo "Building development server..."
	go build -o $(TMP_DIR)/server ./main.go

# server/start - starts the server.
.PHONY: server/start
server/start:
	go run server.go

# server/start_dev - starts the server for development with live reload.
.PHONY: server/start_dev
server/start_dev:
	@echo "Starting development server..."
	air -c $(CONFIG_DIR)/air/server.toml

# server/clean - cleans the server binary and temporary files.
.PHONY: server/clean
server/clean:
	rm -f $(BIN_DIR)/server
	rm -f $(TMP_DIR)/server

# server/help - prints the server help message.
.PHONY: server/help
server/help:
	@echo "Usage: make server/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  build        builds the server server."
	@echo "  build_dev    builds the server for development."
	@echo "  start        starts the server."
	@echo "  start_dev    starts the server for development with live reload."
	@echo "  clean        cleans the server binary and temporary files."
	@echo "  help         prints the server help message."
	@echo ""
	@echo "Examples:"
	@echo "  make server/build"
	@echo "  make server/build_dev"
	@echo "  make server/start"
	@echo "  make server/start_dev"
	@echo "  make server/clean"
	@echo "  make server/help"
	@echo ""
