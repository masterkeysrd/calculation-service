.DEFAULT_GOAL := help

-include ./scripts/tasks/*.mk

# deps - installs the dependencies.
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
	go mod vendor

# start - starts the server.
.PHONY: start
start: deps server/start

# start_dev - starts the server in development mode.
.PHONY: start_dev
start_dev: dev/all server/start_dev

# test - runs the tests.
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./... 

# test - runs the tests.
.PHONY: test_coverage
test_coverage:
	@echo "Running tests..."
	go test -v -cover ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# help - prints the help message.
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  deps                           installs the dependencies"
	@echo "  start                          starts the server"
	@echo "  start_dev                      starts the server in development mode"
	@echo "  dev/<sub-commands>             performs the development tasks"
	@echo "  server/<sub-commands>          performs the server tasks"
	@echo "  api-docs/<sub-commands>        performs the api-docs tasks"
	@echo "  gvm/<sub-commands>             performs the gvm tasks"
	@echo "  live-reload/<sub-commands>     performs the live-reload tasks"
	@echo "  help                           prints the help message"
	@echo ""
	@echo "Examples:"
	@echo "  make deps"
	@echo "  make start"
	@echo "  make start_dev"
	@echo "  make dev/all"
	@echo "  make server/start_dev"
	@echo "  make api-docs/serve"
	@echo "  make gvm/init"
	@echo "  make live-reload/install"
	@echo "  make help"
	@echo ""
	@echo "Additional help topics:"
	@echo "  make dev/help"
	@echo "  make server/help"
	@echo "  make api-docs/help"
	@echo "  make gvm/help"
	@echo "  make live-reload/help"
	@echo ""
