package models

import (
	"time"
)

type Computer struct {
	ID                   string    `json:"id"`
	ComputerName         string    `json:"computer_name"`
	IPAddress            string    `json:"ip_address"`
	MACAddress           string    `json:"mac_address"`
	EmployeeAbbreviation *string   `json:"employee_abbreviation"`
	Description          *string   `json:"description"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
