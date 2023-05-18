# Makefile - api-docs
# This makefile contains the tasks for the api-docs is
# an abstraction layer for the api-docs commands in api/Makefile.

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