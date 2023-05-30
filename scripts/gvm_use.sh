#!/bin/bash/env bash

# Check if gvm is installed first
if ! [ -x "$(command -v gvm)" ]; then
  echo "====> gvm is not installed. Please run scripts/gvm_init.sh"
  exit 1
fi

# Source gvm and use the correct go version and pkgset 
echo "====> Sourcing gvm..."
source $GVM_ROOT/scripts/gvm

echo "====> Using go version and pkgset..."
gvm use $(cat .go-version)
gvm pkgset use $(cat .go-pkgset)
