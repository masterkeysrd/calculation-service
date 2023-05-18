# dev/init initialize the go development environment.
.PHONY: dev/init
dev/init: gvm/init deps

# deps installs the dependencies.
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy

# start - starts the server.
.PHONY: start
start:
	@echo "Starting server..."
	go run ./cmd/server/main.go

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

# api-docs displays the api docs help.
.PHONY: api-docs
api-docs: api-docs/help

# api-docs/serve starts the api docs server.
.PHONY: api-docs/serve
api-docs/serve:
	$(MAKE) -C api api-docs/serve

# api-docs/stop stops the api docs server.
.PHONY: api-docs/stop
api-docs/stop:
	$(MAKE) -C api api-docs/stop

# api-docs/clean cleans the api docs server.
.PHONY: api-docs/clean
api-docs/clean:
	$(MAKE) -C api api-docs/clean

# api-docs/help prints the api docs help message.
.PHONY: api-docs/help
api-docs/help:
	$(MAKE) -C api api-docs/help
