package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(ctx context.Context, body map[string]string) {
	fmt.Println("HI")
}

func main() {
	lambda.Start(LambdaHandler)
}
