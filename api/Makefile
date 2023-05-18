.DEFAULT_GOAL := api-docs/help

COMPOSE_FILE ?= docker-compose.yml
COMPOSE_FILE_PATH = $(CURDIR)/../$(COMPOSE_FILE)
COMPOSE_SERVICE ?= api-docs
DOCKER_COMPOSE_CMD ?= docker compose
DOCKER_COMPOSE_CMD_ARGS = --file $(COMPOSE_FILE_PATH)

# api-docs/serve - starts the api docs server.
.PHONY: api-docs/serve
api-docs/serve:
	@echo "Starting API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) up -d $(COMPOSE_SERVICE)

# api-docs/stop - stops the api docs server.
.PHONY: api-docs/stop
api-docs/stop:
	@echo "Stopping API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) stop $(COMPOSE_SERVICE)

# api-docs/clean - cleans the api docs server.
.PHONY: api-docs/clean
api-docs/clean:
	@echo "Cleaning API docs..."
	@$(DOCKER_COMPOSE_CMD) $(DOCKER_COMPOSE_CMD_ARGS) rm $(COMPOSE_SERVICE)
	
# api-docs/help - prints the help message.
.PHONY: api-docs/help
api-docs/help:
	@echo "Usage: make api-docs/<target> [COMPOSE_FILE=<compose-file>] [COMPOSE_SERVICE=<compose-service>] [DOCKER_COMPOSE_CMD=<docker-compose-cmd>] "
	@echo ""
	@echo "Variables:"
	@echo "  COMPOSE_FILE            the compose file to use (default: docker-compose.yml)"
	@echo "  COMPOSE_SERVICE         the compose service to use (default: api-docs)"
	@echo "  DOCKER_COMPOSE_CMD      the docker compose command to use (default: docker compose)"
	@echo ""
	@echo "Targets:"
	@echo "  serve                   starts the api docs server"
	@echo "  stop                    stops the api docs server"
	@echo "  clean                   cleans the api docs server"
	@echo "  help                    prints the help message"
	@echo ""
	@echo "Examples:"
	@echo "  make api-docs/serve"
	@echo "  make api-docs/stop"
	@echo "  make api-docs/clean"
	@echo "  make api-docs/help"
	@echo ""
	
