package main

import (
	"contact-api/pkg/getter"
	"contact-api/pkg/storage"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	hostname string
	body     map[string]string
}

func validate(ev *Event) error {
	var err error = nil

	if ev.body == nil {
		err = errors.New("body not set")
	}

	if ev.hostname == "" {
		err = errors.New("hostname not set")
	}

	return err
}

func Handler(ctx context.Context, event *Event) (map[string]interface{}, error) {
	fmt.Println(ctx)
	fmt.Println(event)

	err := validate(event)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": err.Error(),
		}, nil
	}

	db, err := storage.Connection()

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": err.Error(),
		}, nil
	}

	// Get the host from url
	host := &getter.Host{
		Url: event.hostname,
	}
	host.Read(db)

	// If no host return 403
	if host.Id < 1 {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": "Invalid host",
		}, nil
	}

	// If no schema return 404
	template := &getter.Template{}
	template.Read(db, host)

	if template.Id < 1 {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": "Template not found for host " + host.Url,
		}, nil
	}

	// If host.has_images store images in s3

	return map[string]interface{}{
		"statusCode": fmt.Sprint(http.StatusCreated),
		"headers": map[string]string{
			"content-type": "application/json",
		},
	}, nil
}

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" &&
		os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(Handler)
	} else {
		Handler(context.TODO(), nil)
	}
}
