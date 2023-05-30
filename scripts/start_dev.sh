#!/bin/bash/env bash
echo "===>  Starting development server..."

# source ./scripts/gvm_install.sh
echo "===>  Using go version manager..."
source scripts/gvm_use.sh

# Install the live reload utility
echo "====> Installing live reload utility..." 
go install github.com/cosmtrek/air@latest

# Run!
echo "===>  Running..."
air -c ./config/air/server.toml
