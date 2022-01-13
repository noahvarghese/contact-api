package main

import (
	"contact-api/pkg/email"
	"contact-api/pkg/getter"
	"contact-api/pkg/setter"
	"contact-api/pkg/storage"
	"context"
	"encoding/json"
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
	if host.ID < 1 {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "application/json",
			},
			"body": "Invalid host",
		}, nil
	}

	// If no schema return 404
	t := &getter.Template{}
	t.Read(db, host.ID)

	if t.ID < 1 {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusBadRequest),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": "Template not found for host " + host.Url,
		}, nil
	}

	// Get all fields for template from db
	fields, err := (&getter.Field{}).GetAll(db, host.ID)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": err.Error(),
		}, nil
	}

	// check that all required fields exist
	for i := 0; i < len(fields); i++ {
		f := fields[i]

		_, ok := event.body[f.Name]

		if f.Required && !ok {
			return map[string]interface{}{
				"statusCode": fmt.Sprint(http.StatusBadRequest),
				"headers": map[string]string{
					"content-type": "text/plain",
				},
				"body": "Missing parameter: " + f.Name,
			}, nil
		}
	}

	jsonStr, err := json.Marshal(event.body)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": err.Error(),
		}, nil
	}

	message, err := setter.NewMessage(db, string(jsonStr), host.ID)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": err.Error(),
		}, nil
	}

	// Create email
	mail, err := email.Bind(t, event.body)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": err.Error(),
		}, nil
	}

	err = email.Send(mail, host)

	if err != nil {
		return map[string]interface{}{
			"statusCode": fmt.Sprint(http.StatusInternalServerError),
			"headers": map[string]string{
				"content-type": "text/plain",
			},
			"body": err.Error(),
		}, nil
	}

	// Send email
	message.SetSent()

	return map[string]interface{}{
		"statusCode": fmt.Sprint(http.StatusCreated),
		"headers": map[string]string{
			"content-type": "text/plain",
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
