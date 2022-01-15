#!/bin/bash

GOOS="linux"
GOARCH="amd64"
INPUT="./cmd/serverless/main.go"
OUTPUT="./build/main"

go build -o $OUTPUT $INPUT