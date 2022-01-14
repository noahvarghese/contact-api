#!/bin/bash

build_env_json_str() {
    envsubst < ./test/env.template.json
}

build_env_json_file() {
    env_json=$(build_env_json_str)
    echo "$env_json" > .env.json
}

check_failed() {
    exit_code=$1

    if [ $exit_code -gt 0 ]; then
        exit $exit_code
    fi
}

get_test_data_json() {
    body_json="'$(cat ./test/body.template.json)'"
    echo $body_json
}

integration_test() {
    # Build the cli version so we don't have the dependency of the api gateway
	GOOS=linux GOARCH=amd64 go build -o cli ./cmd/cli/main.go
        
    # Execute cmd with args 
    ./cli --host=owd.noahvarghese.me --json=./test/body.template.json

    if [ $? -gt 0 ]; then
        printf "\n\nINTEGRATION TEST FAILED\n\n"

        exit 255
    fi

    exit 0
}

e2e_test() {
    build_env_json_file

    # Build Serverless API and Function
    sam build --use-container
    check_failed $?

    # Start Serverless Application Model locally in background process
    sam local start-api --env-vars ./.env.json
    SAM_PID=$!

    # Test
    result=$(curl -d $(get_test_data_json) -H "Content-Type: application/json" -X POST http://localhost:3000/contact)

    kill -2 $SAM_PID

    # Confirm response
    if [[ $result == $(printf "{ \"message\": \"Sent\" }\n") ]]; then
        exit 0
    else
        printf "\n\nE2E TEST FAILED\n\n"
        exit 255
    fi
}


ARGS=(--e2e,-e,--integration,-i)


# check_args $ARGS || exit $?
TEST=""

# Get from cli args whether end to end or integration test
if [[ $1 == --e2e ]] || [[ $1 == '-e' ]]; then
    TEST="e2e_test"
elif [[ $1 == --integration ]] || [[ $1 == '-i' ]]; then
    TEST="integration_test"
else
    printf "Invalid option, must be one of ['--e2e', '-e'] or ['--integration', '-i']\n"
    exit 1
fi

# Setup environment
source ./scripts/env.sh

if ! envup; then
    exit 1
fi

$TEST