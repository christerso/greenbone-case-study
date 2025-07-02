package tests

import (
	"greenbone-computer-inventory/internal/handlers"
	"greenbone-computer-inventory/internal/models"
	"testing"
)

func TestComputerValidation(t *testing.T) {
	handler := &handlers.ComputerHandler{}

	tests := []struct {
		name      string
		computer  models.Computer
		wantError bool
	}{
		{
			name: "valid computer",
			computer: models.Computer{
				ComputerName: "TEST-LAPTOP",
				IPAddress:    "192.168.1.100",
				MACAddress:   "00:11:22:33:44:55",
			},
			wantError: false,
		},
		{
			name: "missing computer name",
			computer: models.Computer{
				IPAddress:  "192.168.1.100",
				MACAddress: "00:11:22:33:44:55",
			},
			wantError: true,
		},
		{
			name: "invalid IP address",
			computer: models.Computer{
				ComputerName: "TEST-LAPTOP",
				IPAddress:    "invalid-ip",
				MACAddress:   "00:11:22:33:44:55",
			},
			wantError: true,
		},
		{
			name: "invalid MAC address",
			computer: models.Computer{
				ComputerName: "TEST-LAPTOP",
				IPAddress:    "192.168.1.100",
				MACAddress:   "invalid-mac",
			},
			wantError: true,
		},
		{
			name: "invalid employee abbreviation length",
			computer: models.Computer{
				ComputerName:         "TEST-LAPTOP",
				IPAddress:            "192.168.1.100",
				MACAddress:           "00:11:22:33:44:55",
				EmployeeAbbreviation: stringPtr("toolong"),
			},
			wantError: true,
		},
		{
			name: "valid employee abbreviation",
			computer: models.Computer{
				ComputerName:         "TEST-LAPTOP",
				IPAddress:            "192.168.1.100",
				MACAddress:           "00:11:22:33:44:55",
				EmployeeAbbreviation: stringPtr("abc"),
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.ValidateComputer(&tt.computer)
			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
