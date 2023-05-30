#!/bin/bash/env bash

# Check if docker is installed
bash ./scripts/deploy_docker/check_docker.sh

echo "===> Deploying database..."

# Stop and remove old containers
echo "Stopping and removing old containers..."
docker-compose -f ./deployments/docker/database/docker-compose.yml down

# Build!
echo "Building..."
docker-compose -f ./deployments/docker/database/docker-compose.yml build

# Deploy!
echo "Deploying..."
docker-compose -f ./deployments/docker/database/docker-compose.yml up -d

# Done!
echo
echo "Database deployment is complete!"
docker-compose -f ./deployments/docker/database/docker-compose.yml ps
echo "Database is running at localhost:5432"
echo "Database pgadmin is running at http://localhost:5050"