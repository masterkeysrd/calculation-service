#!/bin/bash/env bash

# Stop and remove old containers
echo "===>  Stopping and removing old containers..."
docker-compose -f ./deployments/docker/compose.yml down

# Build!
echo "===>  Building..."
docker-compose -f ./deployments/docker/compose.yml build

# Deploy!
echo "===>  Deploying..."
docker-compose -f ./deployments/docker/compose.yml up

# Done!
echo
echo "==> Results:"
docker-compose -f ./deployments/docker/compose.yml ps
