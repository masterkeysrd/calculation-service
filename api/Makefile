.DEFAULT_GOAL := api-docs/help

COMPOSE_FILE ?= docker-compose.yml
COMPOSE_FILE_PATH = $(CURDIR)/../$(COMPOSE_FILE)
COMPOSE_SERVICE ?= api-docs
DOCKER_COMPOSE_CMD ?= docker compose
DOCKER_COMPOSE_CMD_ARGS ?= --file $(COMPOSE_FILE_PATH)

# api-docs/serve starts the api docs server
.PHONY: api-docs/serve
api-docs/serve:
	@echo "Starting API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) up -d $(COMPOSE_SERVICE)

# api-docs/stop stops the api docs server
.PHONY: api-docs/stop
api-docs/stop:
	@echo "Stopping API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) stop $(COMPOSE_SERVICE)

# api-docs/clean cleans the api docs server
.PHONY: api-docs/clean
api-docs/clean:
	@echo "Cleaning API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) rm $(COMPOSE_SERVICE)
	
# api-docs/help prints the help message
.PHONY: api-docs/help
api-docs/help:
	@echo "Usage: make [target] [DOCKER_COMPOSE_CMD=<docker-compose-cmd>]  [COMPOSE_FILE=<compose-file>] [COMPOSE_SERVICE=<compose-service>]"
	@echo ""
	@echo "Targets:"
	@echo "  api-docs/serve - Starts the API docs server"
	@echo "  api-docs/stop - Stops the API docs server"
	@echo "  api-docs/clean - Cleans the API docs server"
	@echo "  api-docs/help - Prints this help message"
	@echo ""
	@echo "Variables:"
	@echo "  COMPOSE_FILE - The compose file to use (default: docker-compose.yml)"
	@echo "  COMPOSE_SERVICE - The compose service to use (default: api-docs)"
	@echo "  DOCKER_COMPOSE_CMD - The docker compose command to use (default: docker compose)"
	@echo ""
	@echo "Examples:"
	@echo "  make api-docs/serve"
	@echo "  make api-docs/stop"
	@echo "  make api-docs/clean"
	@echo "  make api-docs/help"
	@echo "  make api-docs/serve DOCKER_COMPOSE_CMD=docker-compose COMPOSE_FILE=docker-compose.yml COMPOSE_SERVICE=api-docs"
	@echo "  make api-docs/stop DOCKER_COMPOSE_CMD=docker-compose COMPOSE_FILE=docker-compose.yml COMPOSE_SERVICE=api-docs"
	@echo "  make api-docs/clean DOCKER_COMPOSE_CMD=docker-compose COMPOSE_FILE=docker-compose.yml COMPOSE_SERVICE=api-docs"
	@echo "  make api-docs/help DOCKER_COMPOSE_CMD=docker-compose COMPOSE_FILE=docker-compose.yml COMPOSE_SERVICE=api-docs"