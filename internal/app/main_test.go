package app

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const EnvPath = "../../.env"

func TestValidateSuccess(t *testing.T) {
	godotenv.Load(EnvPath)
	body := map[string]interface{}{
		"hostname": "test.com",
		"data": map[string]string{
			"name":  "test",
			"email": "test@test.com",
		},
	}

	err := validate(body)

	assert.Nil(t, err)
}

func TestValidateMissingHost(t *testing.T) {
	godotenv.Load(EnvPath)
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
	godotenv.Load(EnvPath)
	body := map[string]interface{}{
		"hostname": "test.com",
	}

	err := validate(body)

	assert.EqualError(t, err, "body not set")
}

func TestHandlerAllData(t *testing.T) {
	godotenv.Load(EnvPath)
	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "test.com",
		"data": map[string]string{
			"required":     "test",
			"not_required": "test@test.com",
		},
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 201, d["statusCode"].(int))
	assert.Equal(t, "Sent", d["body"].(map[string]string)["message"])
}

// TODO: Test adding a script in body

func TestHandlerRequiredData(t *testing.T) {
	godotenv.Load(EnvPath)
	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "test.com",
		"data": map[string]string{
			"required": "test",
		},
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 201, d["statusCode"].(int))
	assert.Equal(t, "Sent", d["body"].(map[string]string)["message"])

}

func TestHandlerMissingRequired(t *testing.T) {
	godotenv.Load(EnvPath)

	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "test.com",
		"data": map[string]string{
			"not_required": "test@test.com",
		},
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 400, d["statusCode"].(int))
	assert.Equal(t, "Missing parameter: required", d["body"].(map[string]string)["message"])
}

func TestHandlerMissingData(t *testing.T) {
	godotenv.Load(EnvPath)
	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "test.com",
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 400, d["statusCode"].(int))
	assert.Equal(t, "body not set", d["body"].(map[string]string)["message"])
}

func TestHandlerMissingHost(t *testing.T) {
	godotenv.Load(EnvPath)
	b, _ := json.Marshal(map[string]interface{}{})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 400, d["statusCode"].(int))
	assert.Equal(t, "hostname not set", d["body"].(map[string]string)["message"])
}

func TestHandlerInvalidHost(t *testing.T) {
	godotenv.Load(EnvPath)
	b, _ := json.Marshal(map[string]interface{}{
		"hostname": "invalid.com",
		"data":     map[string]string{},
	})

	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: string(b),
	})

	assert.Nil(t, err)
	assert.Equal(t, 403, d["statusCode"].(int))
	assert.Equal(t, "Invalid host", d["body"].(map[string]string)["message"])
}

// func TestHandlerMissingTemplate(t *testing.T) {
// 	godotenv.Load(EnvPath)
// 	b, _ := json.Marshal(map[string]interface{}{
// 		"hostname": "missing_template.com",
// 		"data":     map[string]string{},
// 	})

// 	d, err := Handler(context.TODO(), events.APIGatewayProxyRequest{
// 		Body: string(b),
// 	})

// 	assert.Nil(t, err)
// 	assert.Equal(t, 404, d["statusCode"].(int))
// 	assert.Equal(t, "No template found for host missing_template.com", d["body"].(map[string]string)["message"])
// }
