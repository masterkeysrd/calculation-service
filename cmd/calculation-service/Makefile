DEFAULT_GOAL := calculation-service

ROOT_DIR := $(CURDIR)/../..
BIN_DIR := $(ROOT_DIR)/bin
CONFIG_DIR := $(ROOT_DIR)/config
TMP_DIR := $(ROOT_DIR)/tmp

# calculation-service - Displays the calculation-service help.
.PHONY: calculation-service
calculation-service: calculation-service/help

# build - Builds the calculation-service server.
.PHONY: calculation-service/build
calculation-service/build:
	go build -o $(BIN_DIR)/calculation-service ./main.go

# build_dev - Builds the calculation-service server for development.
.PHONY: calculation-service/build_dev
calculation-service/build_dev:
	go build -o $(TMP_DIR)/calculation-service ./main.go

.PHONY: calculation-service/serve
calculation-service/serve:
	go run main.go

.PHONY: calculation-service/serve_dev
calculation-service/serve_dev:
	@air -c $(CONFIG_DIR)/air/calculation-service.air.toml

# calculation-service/clean - Cleans the calculation-service server.	
.PHONY: calculation-service/clean
calculation-service/clean:
	rm -f $(BIN_DIR)/calculation-service
	rm -f $(TMP_DIR)/calculation-service

# calculation-service/help - Prints the calculation-service help message.
.PHONY: calculation-service/help
calculation-service/help:
	@echo "Usage: make calculation-service/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  build      - Builds the calculation-service server."
	@echo "  build_dev  - Builds the calculation-service server for development."
	@echo "  serve      - Serves the calculation-service server."
	@echo "  serve_dev  - Serves the calculation-service server for development."
	@echo "  clean      - Cleans the calculation-service server."
	@echo "  help       - Prints the calculation-service help message."
	@echo ""
	@echo "Examples:"
	@echo "  make calculation-service/build"
	@echo "  make calculation-service/build_dev"
	@echo "  make calculation-service/serve"
	@echo "  make calculation-service/serve_dev"
	@echo "  make calculation-service/clean"
	@echo "  make calculation-service/help"
	@echo ""

