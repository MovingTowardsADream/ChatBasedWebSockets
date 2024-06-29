include .env

.PHONY:

local-launch:
	go run ./cmd/app/main.go --config=./configs/local.yaml

dev-launch:
	go run ./cmd/app/main.go --config=./configs/dev.yaml

launch:
	go run ./cmd/app/main.go --config=${CONFIG_PATH}

docker-postgres-build:
	docker run -d \
      --name chatBasedWebSocket \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=admin \
      -e POSTGRES_DB=chat \
      -p 5432:5432 \
      postgres