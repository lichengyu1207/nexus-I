.PHONY: dev build test clean docker-up docker-down

APP_NAME := nexus-I
GO_VERSION := 1.21

dev:
	cd backend && go run cmd/gateway/main.go

build:
	cd backend && go build -o ../bin/$(APP_NAME) ./cmd/gateway

worker:
	cd backend && go run cmd/worker/main.go

test:
	cd backend && go test ./...

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate:
	cd backend && go run cmd/gateway/main.go migrate
