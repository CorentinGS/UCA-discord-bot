APP_NAME = uca_bot

OUTPUT_DIR = bin

ci:
	./scripts/ci.sh

test:
	go test -v ./...

build:
	go build -o $(OUTPUT_DIR)/$(APP_NAME) -v .

run:
	go run ./...

clean:
	rm -rf bin/*
