#!/bin/bash/env bash

# Check if docker is installed
bash ./scripts/deploy_docker/check_docker.sh

echo "===> Stopping server..."

# Stop and remove old containers
echo "Stopping and removing old containers..."
docker-compose -f ./deployments/docker/server/docker-compose.yml down

# Done!
echo
echo "Server is stopped!"
docker-compose -f ./deployments/docker/server/docker-compose.yml ps
