#!/bin/bash

build_env_json_str() {
    envsubst < ./test/env.template.json
}

build_env_json_file() {
    env_json=$(build_env_json_str)
    echo "$env_json" > .env.json
}

sam build --use-container
# Need to build .env.json
sam local start-api -n ./.env.json