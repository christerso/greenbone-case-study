# Greenbone Computer Inventory

## Prerequisites

Install the migrate CLI tool:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Build, Test and Run

1. Start PostgreSQL: `make postgres`
2. Start notification service: `make notification-service`  
3. Run migrations: `make migrate-up`
4. Build: `make build`
5. Run: `make run`
6. Test: `go test ./tests/`

## API Usage

Add a computer:
```bash
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-LAPTOP-001",
    "ip_address": "192.168.1.100", 
    "mac_address": "00:11:22:33:44:55",
    "employee_abbreviation": "mmu",
    "description": "Development laptop"
  }'
```

Get all computers:
```bash
curl http://localhost:3000/api/computers
```

Get computers by employee:
```bash
curl http://localhost:3000/api/employees/mmu/computers
```

## Services

- Computer Inventory API: http://localhost:3000
- Notification Service: http://localhost:8080

## Notes and Amendments

I included the .env file in the repository for easy testing. In production this would be excluded.

Security considerations for production use which I thought of during development:
- Unassigned devices should trigger security alerts per CIS Controls
- Device lifecycle management (active/inactive/returned status)
I am sure there are plenty more things to do like: audit logging for all device assignments etc...

By the way, the notification service README has an inconsistency: https://github.com/greenbone/exercise-admin-notification
It mentions /api/notification in the description but the actual endpoint is /api/notify.

Thank you for the opportunity to do the test. I hope you find the code and documentation clear and easy to follow.
Looking forward to your feedback!

Kind regards,
Christer Soederlund
