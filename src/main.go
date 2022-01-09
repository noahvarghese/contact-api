package main

import (
	"context"
	"os"
	"src/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(ctx context.Context, body map[string]string) {
	db := models.Init()
	// get url from lambda request
	// Get the host from url
	var host models.Host
	host.Read(db, "")
	// If no host return 403
	// If no schema return 404
	// If host.has_images store images in s3
}

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" &&
		os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(LambdaHandler)
	} else {
		LambdaHandler(context.TODO(), nil)
	}
}