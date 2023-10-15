project_name = nginx-controller
image_name = capcom6/$(project_name):latest

extension=
ifeq ($(OS),Windows_NT)
	extension = .exe
endif

init:
	go mod download

init-dev: init
	go install github.com/cosmtrek/air@latest \
		&& go install github.com/swaggo/swag/cmd/swag@latest

air:
	air

run:
	go run cmd/$(project_name)/main.go

test:
	go test -cover ./...

build:
	go build ./cmd/$(project_name)
	
install:
	go install ./cmd/$(project_name)

docker-dev:
	docker-compose -f deployments/docker-compose.dev.yml up --build

api-docs:
	swag fmt -g ./cmd/$(project_name)/main.go \
		&& swag init -g ./cmd/$(project_name)/main.go -o ./api

view-docs:
	php -S 127.0.0.1:8080 -t ./api

.PHONY: init init-dev air run test install api-docs docker-dev view-docs