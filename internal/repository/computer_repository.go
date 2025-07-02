package repository

import (
	"database/sql"
	"greenbone-computer-inventory/internal/models"
)

type ComputerRepository struct {
	db *sql.DB
}

func NewComputerRepository(db *sql.DB) *ComputerRepository {
	return &ComputerRepository{db: db}
}

func (r *ComputerRepository) Create(computer *models.Computer) error {
	query := `
		INSERT INTO computers (computer_name, ip_address, mac_address, employee_abbreviation, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, computer.ComputerName, computer.IPAddress, computer.MACAddress,
		computer.EmployeeAbbreviation, computer.Description).Scan(
		&computer.ID, &computer.CreatedAt, &computer.UpdatedAt)
}

func (r *ComputerRepository) GetByID(id string) (*models.Computer, error) {
	computer := &models.Computer{}
	query := `
		SELECT id, computer_name, ip_address, mac_address, employee_abbreviation, description, created_at, updated_at
		FROM computers WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&computer.ID, &computer.ComputerName, &computer.IPAddress, &computer.MACAddress,
		&computer.EmployeeAbbreviation, &computer.Description, &computer.CreatedAt, &computer.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return computer, nil
}

func (r *ComputerRepository) GetAll() ([]*models.Computer, error) {
	query := `
		SELECT id, computer_name, ip_address, mac_address, employee_abbreviation, description, created_at, updated_at
		FROM computers ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers []*models.Computer
	for rows.Next() {
		computer := &models.Computer{}
		err := rows.Scan(
			&computer.ID, &computer.ComputerName, &computer.IPAddress, &computer.MACAddress,
			&computer.EmployeeAbbreviation, &computer.Description, &computer.CreatedAt, &computer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		computers = append(computers, computer)
	}
	return computers, nil
}

func (r *ComputerRepository) GetByEmployee(employeeAbbr string) ([]*models.Computer, error) {
	query := `
		SELECT id, computer_name, ip_address, mac_address, employee_abbreviation, description, created_at, updated_at
		FROM computers WHERE employee_abbreviation = $1 ORDER BY created_at DESC`

	rows, err := r.db.Query(query, employeeAbbr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers []*models.Computer
	for rows.Next() {
		computer := &models.Computer{}
		err := rows.Scan(
			&computer.ID, &computer.ComputerName, &computer.IPAddress, &computer.MACAddress,
			&computer.EmployeeAbbreviation, &computer.Description, &computer.CreatedAt, &computer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		computers = append(computers, computer)
	}
	return computers, nil
}

func (r *ComputerRepository) Update(computer *models.Computer) error {
	query := `
		UPDATE computers 
		SET computer_name = $2, ip_address = $3, mac_address = $4, employee_abbreviation = $5, description = $6, updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at`

	return r.db.QueryRow(query, computer.ID, computer.ComputerName, computer.IPAddress,
		computer.MACAddress, computer.EmployeeAbbreviation, computer.Description).Scan(&computer.UpdatedAt)
}

func (r *ComputerRepository) Delete(id string) error {
	query := `DELETE FROM computers WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ComputerRepository) CountByEmployee(employeeAbbr string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM computers WHERE employee_abbreviation = $1`
	err := r.db.QueryRow(query, employeeAbbr).Scan(&count)
	return count, err
}
