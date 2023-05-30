
#!/bin/bash/env bash

# Check if docker-compose is installed
if ! command -v docker-compose > /dev/null 2>&1; then
    echo "docker-compose is not installed. Please install docker-compose and try again."
    exit 1
fi

# Check if docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Docker is not running. Please start docker and try again."
    exit 1
fi

echo "===> Starting API Docs"

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
