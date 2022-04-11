#!/bin/bash

source ./env.sh && envup
sam build --use-container
sam deploy --no-confirm-changeset --no-fail-on-empty-changeset --stack-name contact-api --s3-bucket $S3_BUCKET --s3-prefix $S3_PREFIX --capabilities CAPABILITY_IAM --region ca-central-1