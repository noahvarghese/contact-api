package main

import (
	"contact-api/internal/app"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" &&
		os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(app.Handler)
	} else {
		fmt.Println("Not running in lambda environment")
	}
}
