#!/bin/bash/env bash

# Check if gvm is installed first
if command -v gvm > /dev/null 2>&1; then
    echo "gvm is not installed. Please install gvm and try again."
    exit 1
fi

# Source gvm and use the correct go version and pkgset 
echo "====> Sourcing gvm..."
source $GVM_ROOT/scripts/gvm

echo "====> Using go version and pkgset..."
gvm use $(cat .go-version)
gvm pkgset use $(cat .go-pkgset)
