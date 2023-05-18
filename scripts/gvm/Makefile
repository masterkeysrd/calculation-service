.DEFAULT_GOAL := gvm

# gvm - displays the gvm help.
.PHONY: gvm
gvm: gvm/help

# gvm/init - installs the go version and initializes the go environment.
.PHONY: gvm/init
gvm/init:
	@echo "Installing go version and initializing the go environment"
	bash $(CURDIR)/gvm_init.sh
	@echo ""

# gvm/use - sets the go version and initializes the go environment.
.PHONY: gvm/use
gvm/use:
	@echo "Setting go version and initializing the go environment"
	bash $(CURDIR)/gvm_use.sh
	@echo ""

# gvm/help - prints the gvm help message.
.PHONY: gvm/help
gvm/help:
	@echo "Usage: make gvm/<target>"
	@echo ""
	@echo "Targets:"
	@echo "  init        installs the go version and initializes the go environment"
	@echo "  use         sets the go version and initializes the go environment"
	@echo "  help        prints the gvm help message"
	@echo ""
	@echo "Examples:"
	@echo "  make gvm/init"
	@echo "  make gvm/use"
	@echo "  make gvm/help"
	@echo ""
