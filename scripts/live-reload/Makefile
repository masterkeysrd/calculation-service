# Title: Live Reload Makefile
.DEFAULT_GOAL := live-reload

# live-reload - Displays the live reload help.
.PHONY: live-reload
live-reload: live-reload/help

# live-reload/install - Installs the live reload tool.
.PHONY: live-reload/install
live-reload/install:
	@echo "Installing live reload..."
	bash $(CURDIR)/install.sh
	@echo ""

# live-reload/help - Prints the live reload help message.
.PHONY: live-reload/help
live-reload/help:
	@echo "Usage: make live-reload/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  install     installs the live reload tool"
	@echo "  help        prints the live reload help message"
	@echo ""
	@echo "Examples:"
	@echo "  make live-reload/install"
	@echo "  make live-reload/help"
	@echo ""