# Makefile - API documentation
.DEFAULT_GOAL := api_docs/help

# api_docs/start - Start API documentation
.PHONY: api_docs/start
api_docs/start:
	@echo "Starting API documentation..."
	bash ./scripts/api_docs_start.sh

# api_docs/stop - Stop API documentation
.PHONY: api_docs/stop
api_docs/stop:
	@echo "Stopping API documentation..."
	bash ./scripts/api_docs_stop.sh


# api_docs/help - Help for API documentation
.PHONY: api_docs/help
api_docs/help:
	@echo "API documentation help"
	@echo "  make api_docs/start - Start API documentation"
	@echo "  make api_docs/stop - Stop API documentation"
	@echo "  make api_docs/help - Help for API documentation"
