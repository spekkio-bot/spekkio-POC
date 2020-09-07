#!/bin/bash
BASE_DIR="$(cd "$(dirname "$0" )" && pwd )"
SRC_DIR=$BASE_DIR/src
BUILD_DIR=$BASE_DIR/bin

cd $SRC_DIR
echo "NOTE: App being built targets Linux/AMD64."
echo "Building app to $BUILD_DIR..."
env GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/spekkio 2> /dev/null
rc=$?
if [[ $rc -ne 0 ]]; then
    echo "Build failed with exit status $rc! Aborting deploy."
    exit 1
else
    echo "Done!"
fi
