#!/bin/bash

GOOS=""
GOARCH=""
OUTPUT=""
INPUT=""

# Get from cli args whether end to end or integration test
if [[ $1 == --serverless-prod ]] || [[ $1 == '-sp' ]]; then
    GOOS="linux"
    GOARCH="amd64"
    OUTPUT="./build/main"
    INPUT="./cmd/serverless/main.go"
    echo "Building for Prod Serverless"
elif [[ $1 == --serverless-local ]] || [[ $1 == '-sl' ]]; then
    GOOS="linux"
    GOARCH="amd64"
    OUTPUT="./build/serverless"
    INPUT="./cmd/serverless/main.go"
    echo "Building for Local Serverless"
elif [[ $1 == --cli ]] || [[ $1 == '-c' ]]; then
    GOOS="linux"
    GOARCH="amd64"
    OUTPUT="./build/cli"
    INPUT="./cmd/cli/main.go"
    echo "Building for CLI test"
else
    printf "Invalid option, must be one of ['--cli', '-c'], ['--serverless-prod', '-sp'] or ['--serverless-local', '-sl']\n"
    exit 1
fi

# export GOOS
# export GOARCH
# export OUTPUT
# export INPUT

go build -o $OUTPUT $INPUT