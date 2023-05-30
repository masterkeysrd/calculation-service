#!/bin/bash/env bash
SERVICE_NAME="calculation-service"
VERSION="1.0.0"
REGISTRY="localhost"

if [ $1 == "help" ]; then
    echo "Usage: ./build_docker.sh [service_name] [version] [registry]"
    echo "Example: ./build_docker.sh calculation-service 1.0.0 localhost"
    exit 0
fi

if [ -z "$1" ]; then
    echo "No service name provided. Using default: $SERVICE_NAME"
else

    SERVICE_NAME=$1
fi

if [ -z "$2" ]; then
    echo "No version provided. Using default: $VERSION"
else
    VERSION=$2
fi

if [ -z "$3" ]; then
    echo "No registry provided. Using default: $REGISTRY"
else
    REGISTRY=$3
fi

echo "==> Building $SERVICE_NAME:$VERSION..."

# Build!
echo "===>  Building..."
docker build \
    -t $REGISTRY/$SERVICE_NAME:$VERSION \
    -f ./cmd/server/Dockerfile \
    .