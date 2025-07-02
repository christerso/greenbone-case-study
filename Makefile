.PHONY: postgres postgres-stop postgres-clean build run help

include .env
export

postgres:
	@echo "Starting PostgreSQL container..."
	docker run -d \
		--name $(POSTGRES_CONTAINER) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(POSTGRES_PORT):5432 \
		postgres:15-alpine
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 5
	@echo "PostgreSQL is running on port $(POSTGRES_PORT)"
	@echo "Database URL: $(DATABASE_URL)"

postgres-stop:
	@echo "Stopping PostgreSQL container..."
	docker stop $(POSTGRES_CONTAINER) || true
	docker rm $(POSTGRES_CONTAINER) || true

postgres-clean: postgres-stop
	@echo "Cleaning up PostgreSQL data..."
	docker volume rm $$(docker volume ls -q --filter name=postgres) 2>/dev/null || true

build:
	go build -o bin/computer-inventory ./cmd

run: build
	./bin/computer-inventory

# Show help
help:
	@echo "Available targets:"
	@echo "  postgres       - Start PostgreSQL container"
	@echo "  postgres-stop  - Stop PostgreSQL container"
	@echo "  postgres-clean - Stop and clean PostgreSQL data"
	@echo "  build          - Build the application"
	@echo "  run            - Build and run the application"
	@echo "  help           - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  POSTGRES_USER     - PostgreSQL username (default: postgres)"
	@echo "  POSTGRES_PASSWORD - PostgreSQL password (default: password)"
	@echo "  POSTGRES_DB       - PostgreSQL database name (default: greenbone_inventory)"
	@echo "  POSTGRES_PORT     - PostgreSQL port (default: 5432)"