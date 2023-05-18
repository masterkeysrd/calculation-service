# Makefile - gvm
# This makefile contains the tasks for the gvm is
# an abstraction layer for the gvm commands in scripts/gvm/Makefile.

# gvm displays the gvm help.
.PHONY: gvm
gvm: gvm/help

# gvm/init installs the go version and initializes the go environment.
.PHONY: gvm/init
gvm/init:
	$(MAKE) -C scripts/gvm gvm/init

# gvm/use sets the go version and initializes the go environment.
.PHONY: gvm/use
gvm/use:
	$(MAKE) -C scripts/gvm gvm/use

# gvm/help prints the gvm help message.
.PHONY: gvm/help
gvm/help:
	$(MAKE) -C scripts/gvm gvm/help
