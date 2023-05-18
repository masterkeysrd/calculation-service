# Makefile: dev
# This makefile contains the tasks for the development.

# dev - displays the help for the development commands.
.PHONY: dev
dev: dev/help

# dev/all - initializes the go development environment and installs the dependencies, and starts the server.
.PHONY: dev/all
dev/all: dev/init dev/deps

# dev/init - initializes the go development environment and installs the dependencies.
.PHONY: dev/init
dev/init: gvm/init dev/deps

# dev/deps - installs the dependencies.
.PHONY: dev/deps
dev/deps: deps gvm/use live-reload/install

# dev/help - prints the help message for the development commands.
.PHONY: dev/help
dev/help:
	@echo "Usage: make dev/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  all                initializes the go development environment and installs the dependencies, and starts the server"
	@echo "  init               initializes the go development environment and installs the dependencies"
	@echo "  deps               installs the dependencies"
	@echo ""
	@echo "Examples:"
	@echo "  make dev/all"
	@echo "  make dev/init"
	@echo "  make dev/deps"
	@echo ""