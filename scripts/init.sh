#!/bin/bash/env bash
echo "===> Initializing..."

# Initialize GVM
echo "====> Initializing GVM..."
source scripts/gvm_init.sh

# Install the do dependecies
echo "====> Installing go dependencies..."
go mod tidy
go mod download
go mod vendor
