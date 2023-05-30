#!/bin/bash

echo "===> Running initializing gvm"

# check if gvm is installed 
if ! [ -x "$(command -v gvm)" ]; then
  echo "====> Installing gvm..."
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
fi

# Source gvm
echo "====> Sourcing gvm..."
source $GVM_ROOT/scripts/gvm

# Read .go-version and .go-pkgset files
echo "====> Reading .go-version and .go-pkgset files..."
GO_VERSION=$(cat .go-version)
GO_PKGSET=$(cat .go-pkgset)

# Install go version and create pkgset
echo "====> Installing go version..."
gvm install $GO_VERSION
gvm use $GO_VERSION

# Create pkgset
echo "====> Creating pkgset..."
gvm pkgset create $GO_PKGSET
gvm pkgset use $GO_PKGSET
