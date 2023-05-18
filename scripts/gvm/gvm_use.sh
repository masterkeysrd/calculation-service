#!/bin/bash

echo "Starting gvm use script"


# Source gvm script
echo "Source gvm"
source $GVM_ROOT/scripts/gvm

# Get root dir
echo "Initialzing variables"
ROOT_DIR=$(pwd)/../..
echo "Root dir: $ROOT_DIR"

# Read .go-version and .go-pkgset files
echo "Reading .go-version and .go-pkgset files"
GO_VERSION=$(cat $ROOT_DIR/.go-version)
GO_PKGSET=$(cat $ROOT_DIR/.go-pkgset)

# Use go version and pkgset
echo "Use go version and pkgset"
gvm use $GO_VERSION
gvm pkgset use $GO_PKGSET