package main

import (
	db "contact-api/src/config"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(ctx context.Context, body map[string]string) {
	_ = db.Init()

	// Get the host
	// If no host return 403
	// If no schema return 404
	// If host.has_images store images in s3
}

func main() {
	lambda.Start(LambdaHandler)
}
