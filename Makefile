APP_NAME=garantex-usdt-rates

.PHONY: build test run lint docker-build

build:
	go build -o app ./cmd/app

test:
	go test ./...

run:
	go run ./cmd/app

lint:
	golangci-lint run

docker-build:
	docker build -t $(APP_NAME) .

