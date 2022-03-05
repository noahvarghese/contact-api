package app

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

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func validate(body map[string]interface{}) error {
	var err error = nil

	if body["hostname"] == nil {
		err = errors.New("hostname not set")
	}

	if body["data"] == nil && err == nil {
		err = errors.New("body not set")
	}

	return err
}

func response(b string, status int) map[string]interface{} {
	body, _ := json.Marshal(map[string]string{"message": b})

	res := map[string]interface{}{
		"isBase64Encoded": false,
		"statusCode":      status,
		"headers": map[string]string{
			"content-type": "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		"body": string(body),
	}

	return res
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (map[string]interface{}, error) {
	b := make(map[string]interface{})
	json.Unmarshal([]byte(event.Body), &b)

	// Check that the expected fields of 'data' and 'hostname' are set in request body
	err := validate(b)

	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}

	hostname := fmt.Sprintf("%v", b["hostname"])
	data := b["data"].(map[string]interface{})

	db, err := storage.Connection()

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	// If no host return 403
	host := &getter.Host{
		Url: hostname,
	}
	err = host.Read(db)

	if err != nil {
		return response("Invalid host", http.StatusForbidden), nil
	}

	// Set Email template
	t := &getter.Template{}
	t.Read(db, host.ID)

	// If no template return 404
	if t.ID < 1 {
		return response("No template found for host "+host.Url, http.StatusNotFound), nil
	}

	// Get all fields for template from db
	fields, err := (&getter.Field{}).GetAll(db, host.ID)

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	// check that all required fields exist
	for i := 0; i < len(fields); i++ {
		f := fields[i]

		_, ok := data[f.Name]

		if f.Required && !ok {
			return response("Missing parameter: "+f.Name, http.StatusBadRequest), nil
		}
	}

	// Store string copy of JSON request body
	jsonStr, err := json.Marshal(data)

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	message, err := setter.NewMessage(db, string(jsonStr), host.ID)

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	// Create email
	mail, err := email.Bind(t, data)

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	err = email.Send(mail, host)

	if err != nil {
		return response(err.Error(), http.StatusInternalServerError), nil
	}

	message.SetSent()

	return response("Sent", http.StatusCreated), nil
}
