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

unit_test() {
    go clean -testcache
    go test ./...
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
    # THIS NEEDS TO BE FIXED AS WELL
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
elif [[ $1 == --unit ]] || [[ $1 == '-u' ]]; then
    TEST="unit_test"
else
    printf "Invalid option, must be one of ['--e2e', '-e'] or ['--unit', '-u']\n"
    exit 1
fi

# Setup environment
source ./scripts/env.sh

if ! envup; then
    exit 1
fi

$TEST