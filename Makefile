.DEFAULT_GOAL := default

###############################################################
## General commands. 										 ##
###############################################################

default: help calculation-service/help gvm/help api-docs/help live_reload/help

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy

###############################################################
## dev - Development commands. 			    				 ##
###############################################################

# dev/init - initialize the go development environment
# and installs the dependencies.
.PHONY: dev/init
dev/init: gvm/init dev/deps

# dev/deps - installs the dependencies.
.PHONY: dev/deps
dev/deps: deps gvm/use live_reload/install
	

###############################################################
## calculation-service - Calculation service commands. 		 ##
###############################################################

.PHONY: calculation-service
calculation-service: calculation-service/help

# build - Builds the calculation-service server.
.PHONY: calculation-service/build
calculation-service/build:
	$(MAKE) -C cmd/calculation-service calculation-service/build

# build_dev - Builds the calculation-service server for development.
.PHONY: calculation-service/build_dev
calculation-service/build_dev:
	$(MAKE) -C cmd/calculation-service calculation-service/build_dev

.PHONY: calculation-service/serve
calculation-service/serve:
	$(MAKE) -C cmd/calculation-service calculation-service/serve

.PHONY: calculation-service/serve_dev
calculation-service/serve_dev:
	$(MAKE) -C cmd/calculation-service calculation-service/serve_dev

# calculation-service/clean - Cleans the calculation-service server.	
.PHONY: calculation-service/clean
calculation-service/clean:
	$(MAKE) -C cmd/calculation-service calculation-service/clean

.PHONY: calculation-service/help
calculation-service/help:
	$(MAKE) -C cmd/calculation-service calculation-service/help

###############################################################
## gvm - Go version manager commands. 	    				 ##
###############################################################

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


###############################################################
## api-docs - API docs commands. 		    				 ##
###############################################################

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

###############################################################
## live_reload - Live reload commands. 		    			 ##
###############################################################

# live_reload - Displays the live reload help.
.PHONY: live_reload
live_reload: live_reload/help

# live_reload/install - installs the live reload tool.
.PHONY: live_reload/install
live_reload/install:
	$(MAKE) -C scripts/live_reload live_reload/install

# live_reload/help - prints the live reload help message.
.PHONY: live_reload/help
live_reload/help:
	$(MAKE) -C scripts/live_reload live_reload/help

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  deps                   - Installs the dependencies."
	@echo "  dev/init               - Initializes the go development environment and installs the dependencies."
	@echo "  dev/deps               - Installs the dependencies."
	@echo ""
	@echo "Examples:"
	@echo "  make deps"
	@echo "  make dev/init"
	@echo "  make dev/deps"
	@echo ""
 
