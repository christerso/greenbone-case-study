#!/bin/bash
# Replace COMPUTER_ID with actual UUID from create response
curl -X DELETE http://localhost:3000/api/computers/$1