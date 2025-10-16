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

// DriverConn Initializes connection to DB
type DriverConn struct {
	conn *pgxpool.Pool
}

// NewDriver Connects to DB
func NewDriver() *DriverConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &DriverConn{
		conn: conn,
	}
}

func driverSelectQuery() string {
	return `SELECT
				driver_id,
				first_name,
				last_name,
				license_number,
				license_expiry,
				phone_number,
				employment_status,
				assigned_vehicle_id,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM driver WHERE deleted_at IS NULL`
}

// Create Inserts new Driver record to DB
func (con *DriverConn) Create(e *entity.Driver) (int32, error) {
	var id int32
	query := `INSERT INTO driver 
				(first_name, last_name, license_number, employment_status, created_by, created_at) 
				VALUES($1, $2, $3, $4, $5, $6) RETURNING driver_id`

	err := con.conn.QueryRow(context.Background(), query,
		e.FirstName,
		e.LastName,
		e.LicenseNumber,
		e.EmploymentStatus,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// List Lists all Driver records
func (con *DriverConn) List() ([]*entity.Driver, error) {
	var driverID pgtype.Int4
	var firstName pgtype.Text
	var lastName pgtype.Text
	var licenseNumber pgtype.Text
	var licenseExpiry pgtype.Timestamp
	var phoneNumber pgtype.Text
	var employmentStatus pgtype.Text
	var assignedVehicleID pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Driver

	query := driverSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Driver{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&driverID,
			&firstName,
			&lastName,
			&licenseNumber,
			&licenseExpiry,
			&phoneNumber,
			&employmentStatus,
			&assignedVehicleID,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Driver{}, err
		}

		item := &entity.Driver{
			DriverID:          driverID.Int,
			FirstName:         firstName.String,
			LastName:          lastName.String,
			LicenseNumber:     licenseNumber.String,
			LicenseExpiry:     &licenseExpiry.Time,
			PhoneNumber:       phoneNumber.String,
			EmploymentStatus:  employmentStatus.String,
			AssignedVehicleID: assignedVehicleID.Int,
			CreatedBy:         createdBy.Int,
			UpdatedBy:         updatedBy.Int,
			DeletedBy:         deletedBy.Int,
			CreatedAt:         createdAt.Time.Local(),
			UpdatedAt:         &updatedAt.Time,
			DeletedAt:         &deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single Driver by ID
func (con *DriverConn) Get(id int32) (*entity.Driver, error) {
	var firstName pgtype.Text
	var lastName pgtype.Text
	var licenseNumber pgtype.Text
	var licenseExpiry pgtype.Timestamp
	var phoneNumber pgtype.Text
	var employmentStatus pgtype.Text
	var assignedVehicleID pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Driver

	query := driverSelectQuery() + ` AND driver_id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&firstName,
		&lastName,
		&licenseNumber,
		&licenseExpiry,
		&phoneNumber,
		&employmentStatus,
		&assignedVehicleID,
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

	item = &entity.Driver{
		DriverID:          id,
		FirstName:         firstName.String,
		LastName:          lastName.String,
		LicenseNumber:     licenseNumber.String,
		LicenseExpiry:     &licenseExpiry.Time,
		PhoneNumber:       phoneNumber.String,
		EmploymentStatus:  employmentStatus.String,
		AssignedVehicleID: assignedVehicleID.Int,
		CreatedBy:         createdBy.Int,
		UpdatedBy:         updatedBy.Int,
		DeletedBy:         deletedBy.Int,
		CreatedAt:         createdAt.Time.Local(),
		UpdatedAt:         &updatedAt.Time,
		DeletedAt:         &deletedAt.Time,
	}
	return item, err
}

// Update Updates single Driver record by ID
func (con *DriverConn) Update(e *entity.Driver) (int32, error) {
	query := `UPDATE driver SET 
				first_name = $1,
				last_name = $2,
				license_number = $3,
				license_expiry = $4,
				phone_number = $5,
				employment_status = $6,
				assigned_vehicle_id = $7,
				updated_by = $8,
				updated_at = $9
			  WHERE driver_id = $10`
	_, err := con.conn.Exec(context.Background(), query,
		e.FirstName,
		e.LastName,
		e.LicenseNumber,
		e.LicenseExpiry,
		e.PhoneNumber,
		e.EmploymentStatus,
		e.AssignedVehicleID,
		e.UpdatedBy,
		time.Now().Local(),
		e.DriverID,
	)
	if err != nil {
		return e.DriverID, err
	}
	return e.DriverID, err
}

// SoftDelete Softly delete single Driver by ID
func (con *DriverConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE driver SET deleted_by = $1, deleted_at = $2 WHERE driver_id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete single Driver by ID
func (con *DriverConn) HardDelete(id int32) error {
	query := "DELETE FROM driver WHERE driver_id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
