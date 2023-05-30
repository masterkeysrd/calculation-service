#!/bin/bash/env bash

# Stop the API docs container
echo "===> Stopping API docs container"
docker-compose -f ./deployments/docker/api_docs/docker-compose.yml down

# Result!
echo "===> API Docs stopped successfully"
docker compose -f ./deployments/docker/api_docs/docker-compose.yml ps
echo "API Docs is stopped"