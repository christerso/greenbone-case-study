.PHONY: postgres postgres-stop postgres-clean notification-service notification-stop build run migrate-up migrate-down migrate-create db-tables db-computers help

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

notification-service:
	@echo "Starting notification service..."
	docker run -d \
		--name greenbone-notification \
		-p 8080:8080 \
		greenbone/exercise-admin-notification
	@echo "Notification service running on port 8080"

notification-stop:
	@echo "Stopping notification service..."
	docker stop greenbone-notification || true
	docker rm greenbone-notification || true

build:
	go build -o bin/computer-inventory ./cmd

run: build
	./bin/computer-inventory

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

db-tables:
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "\dt"

db-computers:
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "SELECT * FROM computers;"

db-clear:
	docker exec $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -c "DELETE FROM computers;"

# Show help
help:
	@echo "Available targets:"
	@echo "  postgres            - Start PostgreSQL container"
	@echo "  postgres-stop       - Stop PostgreSQL container"
	@echo "  postgres-clean      - Stop and clean PostgreSQL data"
	@echo "  notification-service - Start notification service container"
	@echo "  notification-stop   - Stop notification service container"
	@echo "  build              - Build the application"
	@echo "  run                - Build and run the application"
	@echo "  migrate-up     - Run database migrations"
	@echo "  migrate-down   - Rollback database migrations"
	@echo "  migrate-create - Create new migration (usage: make migrate-create name=migration_name)"
	@echo "  db-tables      - List all tables in the database"
	@echo "  db-computers   - List all computers in the database"
	@echo "  help           - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  POSTGRES_USER     - PostgreSQL username (default: postgres)"
	@echo "  POSTGRES_PASSWORD - PostgreSQL password (default: password)"
	@echo "  POSTGRES_DB       - PostgreSQL database name (default: greenbone_inventory)"
	@echo "  POSTGRES_PORT     - PostgreSQL port (default: 5432)"