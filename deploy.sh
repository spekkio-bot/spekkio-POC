#!/bin/bash
BASE_DIR="$(cd "$(dirname "$0" )" && pwd )"
SRC_DIR=$BASE_DIR/src
BUILD_DIR=$BASE_DIR/bin
SERVERLESS_CONFIG=$BASE_DIR/serverless.yml

cd $SRC_DIR
echo "Building app to $BUILD_DIR..."
env GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/spekkio
echo "Deploying via serverless framework..."
serverless deploy
echo "Done!"
