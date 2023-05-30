#!/bin/bash/env bash

# Check if docker is installed
if ! command -v docker > /dev/null 2>&1; then
    echo "docker is not installed. Please install docker and try again."
    exit 1
fi

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