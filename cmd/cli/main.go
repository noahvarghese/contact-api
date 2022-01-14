package main

import (
	"contact-api/internal/app"
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func main() {
	app.Handler(context.TODO(), events.APIGatewayProxyRequest{})
}
