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

# $1 is the title of the test
# $2 is the input file to test
# $3 boolean, false means we are running test to fail, true means test to pass
test() {
    printf $1 

    ./build/cli --data $2 > /tmp/test-output.txt 2>&1

    res=$?

    if [[ $res -eq 0 ]]; then
        if [[ $3 == "true" ]]; then
            printf "\t--\t [ok]\n\n"
        else
            printf "\t--\tTEST FAILED\n\n"
        fi
    else
        if [[ $3 == "true" ]]; then
            printf "\t--\tTEST FAILED\n\n"
        else
            printf "\t--\t [ok]\n\n"
        fi
    fi

    cat /tmp/test-output.txt
    rm /tmp/test-output.txt

    if [[ $res -eq 0 ]] && [[ $3 == "false" ]]; then
        exit 1
    elif [[ $res -gt 0 ]] && [[ $3 == "true" ]]; then
        exit 1
    fi
}

integration_test() {
    # Build the cli version so we don't have the dependency of the api gateway
    ./scripts/build.sh --cli

    test "\nStarting missing host test to fail: " "./test/body.missing_host.json" "false"
    test "\nStarting missing data test to fail: " "./test/body.missing_data.json" "false"
    test "\nStarting test to pass: " "./test/body.valid.json" "true"
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