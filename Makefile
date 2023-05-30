# Makefile for the 'calculation-service' project.
.DEFAULT_GOAL := help

include tasks/api-docs.mk
include tasks/deploy-docker.mk

# init - Initialize the project.
.PHONY: init
init:
	./scripts/init.sh

# build - Build the project. It is not intended to be used directly.
# Use 'build/local' or start instead.
.PHONY: build
build:
	./scripts/build.sh

# build/local - Build the project for local development.
.PHONY: build/local
build/local:
	./scripts/build_local.sh

# build/temp - Build the project for live-reload development.
# It is not intended to be used directly. Use 'start/dev' instead.
.PHONY: build/temp
build/temp:
	./scripts/build_temp.sh

# start - Start the project.
.PHONY: start
start:
	./scripts/start.sh

# start/dev - Start the project for local development with live-reload.
.PHONY: start/dev
start/dev:
	./scripts/start_dev.sh

# docker/build - Build the project for Docker.
.PHONY: docker/build
docker/build:
	./scripts/docker_build.sh

# docker_deploy - Display help for docker deployments.
.PHONY: deploy_docker
deploy_docker:
	$(MAKE) deploy_docker/help


.PHONY: api_docs
api_docs:
	@echo "Starting API documentation..."
	$(MAKE) api_docs/help