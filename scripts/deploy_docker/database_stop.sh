#!/bin/bash/env bash

# Check if docker is installed
bash ./scripts/deploy_docker/check_docker.sh

# Stop and remove old containers
echo "===> Stopping database..."
docker-compose -f ./deployments/docker/database/docker-compose.yml down

# Done!
echo "Database is stopped!"
docker-compose -f ./deployments/docker/database/docker-compose.yml ps
