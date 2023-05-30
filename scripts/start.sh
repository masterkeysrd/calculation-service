#!/bin/bash/env bash

APP_ENV=local

if [ -z "$1" ]; then
    echo "No environment specified, defaulting to local"
else
    APP_ENV=$1
fi

echo "==> Starting $APP_ENV server..."

# source ./scripts/gvm_install.sh
echo "===>  Using go version manager..."
source scripts/gvm_use.sh

# Run!
echo "===>  Running..."
go run ./cmd/server/main.go