# Makefile - live reload
# This makefile contains the tasks for the live reload is
# an abstraction layer for the live reload commands in scripts/live-reload/Makefile.

# live-reload - Displays the live reload help.
.PHONY: live-reload
live-reload: live-reload/help

# live-reload/install - installs the live reload tool.
.PHONY: live-reload/install
live-reload/install:
	$(MAKE) -C scripts/live-reload live-reload/install

# live-reload/help - prints the live reload help message.
.PHONY: live-reload/help
live-reload/help:
	$(MAKE) -C scripts/live-reload live-reload/help