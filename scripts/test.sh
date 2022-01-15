#!/bin/bash

unit_test() {
    go clean -testcache
    go test ./...
}

# Setup environment
source ./scripts/env.sh

if ! envup; then
    exit 1
fi

unit_test