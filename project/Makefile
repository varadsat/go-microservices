SHELL := /bin/bash

## up: start containers
up:
	@echo "Starting Docker containers..."
	docker compose up -d

## up_build: rebuild and start
up_build:
	@echo "Stopping containers..."
	docker compose down
	@echo "Building and starting containers..."
	docker compose up --build -d

## down: stop containers
down:
	@echo "Stopping containers..."
	docker compose down

## logs
logs:
	docker compose logs -f
