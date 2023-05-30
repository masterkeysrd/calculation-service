#!/bin/bash/env bash

# Removing old builds
echo "===> Removing old builds..."
rm -rf ./tmp

# Building temporal binary
echo "===> Building temporal binary..."
go build -o ./tmp/server ./cmd/server/main.go

# Result
echo "===> Result:"
ls -hl ./tmp