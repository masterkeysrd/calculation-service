# Title: Live Reload Makefile
.DEFAULT_GOAL := live_reload

# live_reload - Displays the live reload help.
.PHONY: live_reload
live_reload: live_reload/help

# live_reload/install - Installs the live reload tool.
.PHONY: live_reload/install
live_reload/install:
	@echo "Installing live reload..."
	bash $(CURDIR)/install.sh
	@echo ""

# live_reload/help - Prints the live reload help message.
.PHONY: live_reload/help
live_reload/help:
	@echo "Usage: make live_reload/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  install - Installs the live reload tool."
	@echo "  help    - Prints the live reload help message."
	@echo ""
	@echo "Examples:"
	@echo "  make live_reload/install"
	@echo "  make live_reload/help"
	@echo ""