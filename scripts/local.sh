#!/bin/bash

build_env_json_str() {
    envsubst < ./test/env.template.json
}

build_env_json_file() {
    env_json=$(build_env_json_str)
    echo "$env_json" > .env.json
}

# Loads environment variables into memory and build json file for development api to use
source ./env.sh && envup && build_env_json_file

sam build --use-container
sam local start-api -n ./.env.json