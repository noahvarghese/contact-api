.PHONY: build

build:
	sam build --use-container

build-ContactFunction:
	./scripts/build.sh --serverless-prod
	cp ./build/main $(ARTIFACTS_DIR)