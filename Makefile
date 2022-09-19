.PHONY: help build build-local up down logs ps

DOCKER_TAG := latest
build: ## build docker image to release
	docker build -t teru-0529/gotodo:${DOCKER_TAG} --target deploy ./

build-local: ## build docker image to local development
	docker compose build --no-cache

up: ## do docker compose up with hot release
	docker compose up -d

down: ## do docker compose down
	docker compose down

logs: ## tail docker compose logs
	docker compose logs -f

ps: ## check container status
	docker compose ps

test: ## execute tests
	go test -v ./...
