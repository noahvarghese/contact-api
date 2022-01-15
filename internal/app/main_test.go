package app

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidateSuccess(t *testing.T) {
	body := map[string]interface{}{
		"hostname": "owd.noahvarghese.me",
		"data": map[string]string{
			"name":  "test",
			"email": "test@test.com",
		},
	}

	err := validate(body)

	assert.Nil(t, err)
}

func TestValidateMissingHost(t *testing.T) {
	body := map[string]interface{}{
		"data": map[string]string{
			"name":  "test",
			"email": "test@test.com",
		},
	}

	err := validate(body)

	assert.EqualError(t, err, "hostname not set")
}

func TestValidateMissingData(t *testing.T) {
	body := map[string]interface{}{
		"hostname": "owd.noahvarghese.me",
	}

	err := validate(body)

	assert.EqualError(t, err, "body not set")
}

func TestHandlerPass(t *testing.T) {
	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "owd.noahvarghese.me",
		"data": map[string]string{
			"name":  "test",
			"email": "test@test.com",
		},
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, d["statusCode"].(int), 201)
	assert.Equal(t, d["body"].(map[string]string)["message"], "Sent")
}
