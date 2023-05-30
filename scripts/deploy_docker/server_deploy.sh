#!/bin/bash/env bash

# Check if docker is installed
bash ./scripts/deploy_docker/check_docker.sh

echo "===> Deploying server..."

# Stop and remove old containers
echo "Stopping and removing old containers..."
docker-compose -f ./deployments/docker/server/docker-compose.yml down

# Build!
echo "Building..."
docker-compose -f ./deployments/docker/server/docker-compose.yml build

# Deploy!
echo "Deploying..."
docker-compose -f ./deployments/docker/server/docker-compose.yml up -d

# Done!
echo
echo "Server deployment is complete!"
docker-compose -f ./deployments/docker/server/docker-compose.yml ps
echo "Server is running at http://localhost:8080"
