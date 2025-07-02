CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS computers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    computer_name VARCHAR(255) NOT NULL,
    ip_address INET NOT NULL,
    mac_address VARCHAR(17) NOT NULL UNIQUE,
    employee_abbreviation VARCHAR(3),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_computers_mac_address ON computers(mac_address);
CREATE INDEX idx_computers_employee_abbreviation ON computers(employee_abbreviation);
CREATE INDEX idx_computers_ip_address ON computers(ip_address);