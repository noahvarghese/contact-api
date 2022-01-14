package main

import (
	"contact-api/internal/app"
	"contact-api/pkg/file"
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
)

func load_dotenv(dotenv string) {
	if dotenv != "" && file.Exists(dotenv) {
		godotenv.Load(dotenv)
	}
}

func get_json_data(file string) string {
	if file == "" {
		fmt.Fprintln(os.Stderr, "No data file provided")
		os.Exit(1)
	}

	b, err := os.ReadFile(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "%w", err)
		os.Exit(1)
	}

	data := string(b)

	if data == "" {
		fmt.Fprintln(os.Stderr, "Empty file", file)
		os.Exit(1)
	}

	return data
}

func main() {
	dotenv := flag.String("env", "", "Path to .env")
	file := flag.String("data", "", "Path to the JSON file containing the data")

	flag.Parse()

	load_dotenv(*dotenv)
	data := get_json_data(*file)

	app.Handler(context.TODO(), events.APIGatewayProxyRequest{
		Body: data,
	})
}
