# Greenbone Computer Inventory

## Prerequisites

Install the migrate CLI tool with PostgreSQL support:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Setup

1. Start PostgreSQL: `make postgres`
2. Start notification service: `make notification-service`
3. Run migrations: `make migrate-up`
4. Build and run: `make run`

## Available Commands

- `make postgres` - Start PostgreSQL container
- `make notification-service` - Start notification service container
- `make migrate-up` - Run database migrations
- `make db-tables` - List tables in database
- `make db-computers` - List all computers in database
- `make db-clear` - Clear all computers from database
- `make build` - Build the application
- `make run` - Build and run the application

## API Endpoints

- `POST /api/computers` - Add new computer
- `GET /api/computers` - Get all computers
- `GET /api/computers/:id` - Get single computer by ID
- `PUT /api/computers/:id` - Update/reassign computer
- `DELETE /api/computers/:id` - Delete computer
- `GET /api/employees/:employee/computers` - Get computers by employee

## API Usage Examples

**Add a computer:**
```bash
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-LAPTOP-001",
    "ip_address": "192.168.1.100",
    "mac_address": "00:11:22:33:44:55",
    "employee_abbreviation": "mmu",
    "description": "Max Mustermann development laptop"
  }'
```

**Get all computers:**
```bash
curl http://localhost:3000/api/computers
```

**Get computers by employee:**
```bash
curl http://localhost:3000/api/employees/mmu/computers
```

## Services

- **Computer Inventory API**: `http://localhost:3000`
- **Notification Service**: `http://localhost:8080` (Greenbone service)

### NOTES: 

As this is a job test, I am including .env file in the repository, to make it easy to run the code without any issues.
Normally, this would be excluded from the repository by the .gitignore file.

I am using basic Golang db calls here, but in production I would prefer GORM or pgx for database operations.

## Some notes I took as I wrote the implementation:
If this was in a production environment, we should also check:
* Unassigned devices (where there is no abbreviation): Computers with no employee abbreviation could indicate rogue/unauthorized devices on the network
  If you want to follow CIS Controls, unassigned devices should trigger immediate security notifications rather than being silently stored 
* When it comes to device lifecycle management, we should probably track all statuses: active/inactive/returned

