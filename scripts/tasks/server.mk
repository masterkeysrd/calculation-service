## Makefile - server tasks
## This makefile contains the tasks for the server is
## an abstraction layer for the server commands in cmd/server/Makefile.

# server - displays the server help.
.PHONY: server
server: server/help

# server/build - builds the server server.
.PHONY: server/build
server/build:
	$(MAKE) -C cmd/server server/build

# server/build_dev - builds the server for development.
.PHONY: server/build_dev
server/build_dev:
	$(MAKE) -C cmd/server server/build_dev

# server/start - starts the server.
.PHONY: server/start
server/start:
	$(MAKE) -C cmd/server server/start

# server/start_dev - starts the server for development with live reload.
.PHONY: server/start_dev
server/start_dev:
	$(MAKE) -C cmd/server server/start_dev

# server/clean - cleans the server binary and temporary files.
.PHONY: server/clean
server/clean:
	$(MAKE) -C cmd/server server/clean

# server/help - prints the server help message.
.PHONY: server/help
server/help:
	$(MAKE) -C cmd/server server/help