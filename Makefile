include .env

.PHONY:

local-launch:
	go run ./cmd/app/main.go --config=./configs/local.yaml

dev-launch:
	go run ./cmd/app/main.go --config=./configs/dev.yaml

launch:
	go run ./cmd/app/main.go --config=${CONFIG_PATH}