#!/bin/bash/env bash
echo "===> Building local..."

# source ./scripts/gvm_install.sh
echo "===>  Using go version manager..."
source scripts/gvm_use.sh

# Run the build
echo "====> Running build..."
bash ./scripts/build.sh