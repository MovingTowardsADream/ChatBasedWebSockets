include .env

.PHONY:

local-launch-type-local:
	go run ./cmd/app/main.go --config=./configs/local.yaml

local-launch-type-dev:
	go run ./cmd/app/main.go --config=./configs/dev.yaml

local-launch:
	go run ./cmd/app/main.go --config=${CONFIG_PATH}

docker-postgres-build:
	docker run -d \
      --name chatBasedWebSocket \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=admin \
      -e POSTGRES_DB=chat \
      -p 5432:5432 \
      postgres

migrations-down:
	migrate -path ./migrations -database 'postgres://postgres:admin@localhost:5432/chat?sslmode=disable' down