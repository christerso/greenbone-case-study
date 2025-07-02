#!/bin/bash
# Replace COMPUTER_ID with actual UUID from create response
curl -X PUT http://localhost:3000/api/computers/$1 \
  -H "Content-Type: application/json" \
  -d '{
    "computer_name": "DEV-LAPTOP-001-UPDATED",
    "ip_address": "192.168.1.100",
    "mac_address": "00:11:22:33:44:55",
    "employee_abbreviation": "abc",
    "description": "Reassigned to another employee"
  }'