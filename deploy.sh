#!/bin/bash
BASE_DIR="$(cd "$(dirname "$0" )" && pwd )"
SRC_DIR=$BASE_DIR/src
BUILD_DIR=$BASE_DIR/bin
SERVERLESS_CONFIG=$BASE_DIR/serverless.yml

cd $SRC_DIR
echo "Building app to $BUILD_DIR..."
env GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/spekkio 2> /dev/null
rc=$?
if [[ $rc -ne 0 ]]; then
    echo "Build failed with exit status $rc! Aborting deploy."
    exit 1
fi
#   if [ -d $BUILD_DIR/queries ]; then
#       echo "No queries directory found in $BUILD_DIR."
#       echo "Creating queries directory in $BUILD_DIR..."
#       mkdir $BUILD_DIR/queries
#   fi
#   echo "Copying SQL queries to $BUILD_DIR/queries/sql..."
#   cp -r $SRC_DIR/queries/sql $BUILD_DIR/queries/sql
echo "Deploying via serverless framework..."
cd $BASE_DIR
serverless deploy 2> /dev/null
rc=$?
if [[ $rc -eq 0 ]]; then
    echo "Done!"
else
    echo "Deploy failed with exit status $rc!"
    exit 1
fi
