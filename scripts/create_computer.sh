#!/bin/bash
curl -X POST http://localhost:3000/api/computers \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-LAPTOP-001",
    "ip_address": "192.168.1.100",
    "mac_address": "00:11:22:33:44:55",
    "employee_abbreviation": "mmu",
    "description": "Max Mustermann development laptop"
  }'