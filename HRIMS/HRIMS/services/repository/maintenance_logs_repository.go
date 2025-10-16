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

// MaintenanceConn Initializes connection to DB
type MaintenanceConn struct {
	conn *pgxpool.Pool
}

// NewMaintenance Connects to DB
func NewMaintenance() *MaintenanceConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &MaintenanceConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *MaintenanceConn) Create(e *entity.Maintenance) (int32, error) {
	var id int32
	query := `INSERT INTO maintenance 
		(vehicle_id, maintenance_type, description, maintenance_date, service_provider, cost, next_service_due, created_by, created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING maintenance_id`

	err := con.conn.QueryRow(context.Background(), query,
		e.VehicleID,
		e.MaintenanceType,
		e.Description,
		e.MaintenanceDate,
		e.ServiceProvider,
		e.Cost,
		e.NextServiceDue,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckMaintenance Checks if maintenance exists for a vehicle on specific date
func (con *MaintenanceConn) CheckMaintenance(vehicleID int32, date time.Time) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM maintenance WHERE vehicle_id=$1 AND maintenance_date=$2)"
	err := con.conn.QueryRow(context.Background(), query, vehicleID, date).Scan(&exists)
	return exists, err
}

func maintenanceSelectQuery() string {
	return `SELECT
				maintenance_id,
				vehicle_id,
				maintenance_type,
				description,
				maintenance_date,
				service_provider,
				cost,
				next_service_due,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM maintenance WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *MaintenanceConn) List() ([]*entity.Maintenance, error) {
	var id, vehicleID, createdBy, updatedBy, deletedBy pgtype.Int4
	var maintenanceType, description, serviceProvider pgtype.Text
	var maintenanceDate, nextServiceDue, createdAt, updatedAt, deletedAt pgtype.Timestamp
	var cost pgtype.Numeric

	var items []*entity.Maintenance
	query := maintenanceSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Maintenance{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id, &vehicleID, &maintenanceType, &description, &maintenanceDate, &serviceProvider,
			&cost, &nextServiceDue, &createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
		); err != nil {
			return []*entity.Maintenance{}, err
		}

		var costValue float64
		cost.AssignTo(&costValue)

		item := &entity.Maintenance{
			MaintenanceID:   id.Int,
			VehicleID:       vehicleID.Int,
			MaintenanceType: maintenanceType.String,
			Description:     description.String,
			MaintenanceDate: maintenanceDate.Time.Local(),
			ServiceProvider: serviceProvider.String,
			Cost:            costValue,
			NextServiceDue:  nextServiceDue.Time.Local(),
			CreatedBy:       createdBy.Int,
			UpdatedBy:       updatedBy.Int,
			DeletedBy:       deletedBy.Int,
			CreatedAt:       createdAt.Time.Local(),
			UpdatedAt:       updatedAt.Time.Local(),
			DeletedAt:       deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *MaintenanceConn) Get(id int32) (*entity.Maintenance, error) {
	var vehicleID, createdBy, updatedBy, deletedBy pgtype.Int4
	var maintenanceType, description, serviceProvider pgtype.Text
	var maintenanceDate, nextServiceDue, createdAt, updatedAt, deletedAt pgtype.Timestamp
	var cost pgtype.Numeric

	var item *entity.Maintenance
	query := maintenanceSelectQuery() + ` AND maintenance_id=$1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id, &vehicleID, &maintenanceType, &description, &maintenanceDate, &serviceProvider,
		&cost, &nextServiceDue, &createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return item, err
	}

	var costValue float64
	cost.AssignTo(&costValue)

	item = &entity.Maintenance{
		MaintenanceID:   id,
		VehicleID:       vehicleID.Int,
		MaintenanceType: maintenanceType.String,
		Description:     description.String,
		MaintenanceDate: maintenanceDate.Time.Local(),
		ServiceProvider: serviceProvider.String,
		Cost:            costValue,
		NextServiceDue:  nextServiceDue.Time.Local(),
		CreatedBy:       createdBy.Int,
		UpdatedBy:       updatedBy.Int,
		DeletedBy:       deletedBy.Int,
		CreatedAt:       createdAt.Time.Local(),
		UpdatedAt:       updatedAt.Time.Local(),
		DeletedAt:       deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *MaintenanceConn) Update(e *entity.Maintenance) (int32, error) {
	query := `UPDATE maintenance SET 
				maintenance_type=$1, description=$2, maintenance_date=$3, service_provider=$4, cost=$5, next_service_due=$6, updated_by=$7, updated_at=$8
				WHERE maintenance_id=$9`
	_, err := con.conn.Exec(context.Background(), query,
		e.MaintenanceType,
		e.Description,
		e.MaintenanceDate,
		e.ServiceProvider,
		e.Cost,
		e.NextServiceDue,
		e.UpdatedBy,
		time.Now().Local(),
		e.MaintenanceID,
	)
	return e.MaintenanceID, err
}

// SoftDelete Softly delete single record by ID
func (con *MaintenanceConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE maintenance SET deleted_by=$1, deleted_at=$2 WHERE maintenance_id=$3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete single record by ID
func (con *MaintenanceConn) HardDelete(id int32) error {
	query := "DELETE FROM maintenance WHERE maintenance_id=$1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
