#!/bin/bash

echo "Creating first computer..."
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-LAPTOP-001",
    "ip_address": "192.168.1.100",
    "mac_address": "00:11:22:33:44:55",
    "employee_abbreviation": "mmu",
    "description": "Max laptop"
  }'

echo -e "\n\nCreating second computer..."
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-PHONE-002",
    "ip_address": "192.168.1.101",
    "mac_address": "00:aa:bb:cc:dd:ee",
    "employee_abbreviation": "mmu",
    "description": "Max phone"
  }'

echo -e "\n\nCreating third computer (should trigger notification)..."
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-TABLET-003",
    "ip_address": "192.168.1.102",
    "mac_address": "00:ff:ee:dd:cc:bb",
    "employee_abbreviation": "mmu",
    "description": "Max tablet - should trigger warning"
  }'

echo -e "\n\nDone! Check notification service logs with: docker logs greenbone-notification"