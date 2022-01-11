.PHONY: build

build:
	sam build

build-ContactFunction:
	GOOS=linux GOARCH=amd64 go build -o contact ./cmd/serverless/main.go
	cp ./contact $(ARTIFACTS_DIR)