.DEFAULT_GOAL := gvm

# gvm displays the gvm help
.PHONY: gvm
gvm: gvm/help

# gvm/init installs the go version and initializes the go environment
.PHONY: gvm/init
gvm/init:
	bash $(CURDIR)/gvm_init.sh

# gvm/use sets the go version and initializes the go environment
.PHONY: gvm/use
gvm/use:
	bash $(CURDIR)/gvm_use.sh

# gvm/help prints the gvm help message
.PHONY: gvm/help
gvm/help:
	@echo "gvm/init installs the go version and initializes the go environment"
	@echo "gvm/use sets the go version and initializes the go environment"
	@echo "gvm/help prints the gvm help message"
	@echo ""
	@echo "Examples:"
	@echo "  make gvm/init"
	@echo "  make gvm/use"
	@echo "  make gvm/help"
