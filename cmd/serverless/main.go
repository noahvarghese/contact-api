package main

import (
	"contact-api/pkg/getter"
	"contact-api/pkg/storage"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	headers map[string]string
}

func Handler(ctx context.Context, event *Event) (string, error) {
	fmt.Println(ctx)
	fmt.Println(event)

	db := storage.Connection()

	if event == nil {
		fmt.Println("Body not set")
		os.Exit(1)
	}

	if event.headers == nil {
		fmt.Println("Headers not set")
		os.Exit(1)
	}

	// get url from lambda request
	hostName := event.headers["host"]

	// Get the host from url
	host := &getter.Host{
		Url: hostName,
	}
	host.Read(db)

	// If no host return 403
	if host.Id < 1 {
		return "", errors.New("Forbidden: Invalid url " + hostName)
	}

	// If no schema return 404
	template := &getter.Template{}
	template.Read(db, host)

	if template.Id < 1 {
		return "", errors.New("Template not found for " + host.Url)
	}

	// If host.has_images store images in s3
	return "Success", nil
}

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" &&
		os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(Handler)
	} else {
		Handler(context.TODO(), nil)
	}
}
