package repository

import (
	"context"
	"fmt"
	"os"
	"time"
	"training-backend/services/database"
	"training-backend/services/entity"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

// EmployeesConn handles DB operations for employees table
type EmployeesConn struct {
	conn *pgxpool.Pool
}

// NewEmployees initializes DB connection
func NewEmployees() *EmployeesConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil // Return nil if connection fails
	}
	return &EmployeesConn{
		conn: conn,
	}
}

// Create inserts a new employee into the database
func (con *EmployeesConn) Create(e *entity.Employees) (int32, error) {
	var id int32
	query := `INSERT INTO employees (first_name, last_name, email, phone_number, department, position, hire_date, created_by, created_at)
			  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Department,
		e.Position,
		e.HireDate,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckEmployee checks if employee exists by email
func (con *EmployeesConn) CheckEmployee(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM employees WHERE email = $1)"
	err := con.conn.QueryRow(context.Background(), query, email).Scan(&exists)
	return exists, err
}

func employeesSelectQuery() string {
	return `SELECT 
				id,
				first_name,
				last_name,
				email,
				phone_number,
				department,
				position,
				hire_date,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM employees WHERE deleted_at IS NULL`
}

// List returns all employees
func (con *EmployeesConn) List() ([]*entity.Employees, error) {
	var id pgtype.Int4
	var firstName, lastName, email, department, position pgtype.Text
	var phoneNumber pgtype.Int4
	var hireDate pgtype.Date
	var createdBy, updatedBy, deletedBy pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.Employees

	query := employeesSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Employees{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&firstName,
			&lastName,
			&email,
			&phoneNumber,
			&department,
			&position,
			&hireDate,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Employees{}, err
		}

		item := &entity.Employees{
			ID:          id.Int,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Email:       email.String,
			PhoneNumber: phoneNumber.Int,
			Department:  department.String,
			Position:    position.String,
			HireDate:    hireDate.Time,
			CreatedBy:   createdBy.String,
			UpdatedBy:   updatedBy.String,
			DeletedBy:   deletedBy.String,
			CreatedAt:   createdAt.Time,
			UpdatedAt:   updatedAt.Time,
			DeletedAt:   deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get retrieves a single employee by ID
func (con *EmployeesConn) Get(id int32) (*entity.Employees, error) {
	var dbID pgtype.Int4
	var firstName, lastName, email, department, position, createdBy, updatedBy, deletedBy pgtype.Text
	var phoneNumber pgtype.Int4
	var hireDate pgtype.Date
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.Employees

	query := employeesSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&dbID,
		&firstName,
		&lastName,
		&email,
		&phoneNumber,
		&department,
		&position,
		&hireDate,
		&createdBy,
		&updatedBy,
		&deletedBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return item, err
	}

	item = &entity.Employees{
		ID:          dbID.Int,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Email:       email.String,
		PhoneNumber: phoneNumber.Int,
		Department:  department.String,
		Position:    position.String,
		HireDate:    hireDate.Time,
		CreatedBy:   createdBy.String,
		UpdatedBy:   updatedBy.String,
		DeletedBy:   deletedBy.String,
		CreatedAt:   createdAt.Time,
		UpdatedAt:   updatedAt.Time,
		DeletedAt:   deletedAt.Time,
	}
	return item, err
}

// Update modifies an existing employee
func (con *EmployeesConn) Update(e *entity.Employees) (int32, error) {
	query := `UPDATE employees SET 
				first_name = $1,
				last_name = $2,
				email = $3,
				phone_number = $4,
				department = $5,
				position = $6,
				hire_date = $7,
				updated_by = $8,
				updated_at = $9
			WHERE id = $10`
	_, err := con.conn.Exec(context.Background(), query,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Department,
		e.Position,
		e.HireDate,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	return e.ID, err
}

// SoftDelete marks an employee as deleted
func (con *EmployeesConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE employees SET deleted_by=$1, deleted_at=$2 WHERE id=$3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently removes an employee
func (con *EmployeesConn) HardDelete(id int32) error {
	query := "DELETE FROM employees WHERE id=$1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
