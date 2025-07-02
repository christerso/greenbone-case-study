# Greenbone Computer Inventory

## Prerequisites

Install the migrate CLI tool with PostgreSQL support:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Setup

1. Start PostgreSQL: `make postgres`
2. Run migrations: `make migrate-up`
3. Build and run: `make run`

## Available Commands

- `make postgres` - Start PostgreSQL container
- `make migrate-up` - Run database migrations
- `make db-tables` - List tables in database
- `make db-computers` - List all computers in database
- `make build` - Build the application
- `make run` - Build and run the application

### NOTES: 

As this is a job test, I am including .env file in the repository, to make it easy to run the code without any issues.
Normally, this would be excluded from the repository by the .gitignore file.

I am using basic Golang db calls here, but in production I would prefer GORM or pgx for database operations.
