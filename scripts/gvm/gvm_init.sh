#!/bin/bash

# Install gvm
echo "Source gvm"
source $GVM_ROOT/scripts/gvm

ROOT_DIR=$(pwd)/../..
echo "Root dir: $ROOT_DIR"

# Read .go-version and .go-pkgset files
echo "Reading .go-version and .go-pkgset files"
GO_VERSION=$(cat $ROOT_DIR/.go-version)
GO_PKGSET=$(cat $ROOT_DIR/.go-pkgset)

# Install go version and create pkgset
echo "Install go version"
gvm install $GO_VERSION
gvm use $GO_VERSION

# Create pkgset
echo "Create pkgset"
gvm pkgset create $GO_PKGSET
gvm pkgset use $GO_PKGSET