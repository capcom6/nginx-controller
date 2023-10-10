project_name = nginx-controller

init:
	go mod download

run:
	go run cmd/$(project_name)/main.go

test:
	go test -cover ./...

install:
	go install ./cmd/$(project_name)

docker-dev:
	docker-compose -f deployments/docker-compose.dev.yml up --build

.PHONY: init run test install docker-dev