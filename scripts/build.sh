#!/bin/bash/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Remove existing builds
echo "Removing existing builds..."
rm -rf bin/*
rm -rf pkg/*
mkdir -p bin


 # Download dependencies
echo "===> Downloading dependencies..."
go mod download

# Build!
echo "===>  Building..."
go build -o bin/ ./...


# Done!
echo
echo "==> Results:"
ls -hl bin/
