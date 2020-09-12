#!/bin/bash
BASE_DIR="$(cd "$(dirname "$0" )" && pwd )"
SRC_DIR=$BASE_DIR/src
BUILD_DIR=$BASE_DIR/bin
SERVERLESS_CONFIG=$BASE_DIR/serverless.yml

function build {
    echo "NOTE: App being built targets Linux/AMD64."
    echo -n "Building app to $BUILD_DIR... "
    env GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/spekkio 2> /dev/null
    rc=$?
    if [[ $rc -ne 0 ]]; then
        echo "ERR!"
        echo "Build failed with exit status $rc! Aborting deploy."
        exit $rc
    fi
    echo "done"
}

function deploy {
    echo "Deploying via serverless framework... "
    cd $BASE_DIR
    serverless deploy 2> /dev/null
    rc=$?
    if [[ $rc -ne 0 ]]; then
        echo "ERR!"
        echo "Deploy failed with exit status $rc!"
        exit $rc
    fi
}

cd $SRC_DIR

if [[ $# -eq 0 ]]; then
    go run main.go dev
elif [[ $# -eq 1 ]]; then
    case $1 in
    b | build)
        build
        rc=$?
        if [[ $rc -ne 0 ]]; then
            exit 1
        fi
        ;;
    d | deploy)
        build
        rc=$?
        if [[ $rc -ne 0 ]]; then
            exit $rc
        fi
        deploy
        rc=$?
        if [[ $rc -ne 0 ]]; then
            exit $rc
        fi
        ;;
    t | test)
        stamp=$(date +"%s")
        output=/tmp/$stamp.out
        go test -v ./... -coverprofile=$output
        go tool cover -html=$output
        ;;
    *)
        echo "Invalid options."
        exit 1
        ;;
    esac
fi
echo "Done!"
