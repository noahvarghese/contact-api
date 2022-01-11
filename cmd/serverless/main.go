package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	headers map[string]string
}

func getHost(e *Event) (string, error) {
	if e == nil {
		return "", errors.New("body not set")
	}

	if e.headers == nil {
		return "", errors.New("headers not set")
	}

	// get url from lambda request
	h := e.headers["host"]
	return h, nil
}

func Handler(ctx context.Context, event *Event) (map[string]interface{}, error) {
	fmt.Println(ctx)
	fmt.Println(event)

	_, err := getHost(event)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": err.Error(),
		}, nil
	}

	// db := storage.Connection()

	// // Get the host from url
	// host := &getter.Host{
	// 	Url: hostName,
	// }
	// host.Read(db)

	// // If no host return 403
	// if host.Id < 1 {
	// 	return "", errors.New("Forbidden: Invalid url " + hostName)
	// }

	// // If no schema return 404
	// template := &getter.Template{}
	// template.Read(db, host)

	// if template.Id < 1 {
	// 	return "", errors.New("Template not found for " + host.Url)
	// }

	// // If host.has_images store images in s3
	// return "Success", nil
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
