.PHONY: build

build:
	sam build

build-ContactFunction:
	GOOS=linux GOARCH=amd64 go build -o main ./cmd/serverless/main.go
	cp ./main $(ARTIFACTS_DIR)