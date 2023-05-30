
#!/bin/bash/env bash

# Check if docker is installed
bash ./scripts/deploy_docker/check_docker.sh

echo "===> Deploying API Docs..."

# Remove existing containers
echo "Removing existing containers"
docker-compose -f ./deployments/docker/api_docs/docker-compose.yml down

# Start containers
echo "Starting containers"
docker-compose -f ./deployments/docker/api_docs/docker-compose.yml up -d

# Check if containers are running
echo "Checking if containers are running"
if ! docker ps | grep api-docs > /dev/null 2>&1; then
    echo "Containers are not running. Please check the logs and try again."
    exit 1
fi

# Result!
echo "===> API Docs started successfully"
docker compose -f ./deployments/docker/api_docs/docker-compose.yml ps
echo "API Docs is running at http://localhost:8090"
