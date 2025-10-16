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

// SportsEquipmentConn Initializes connection to DB
type SportsEquipmentConn struct {
	conn *pgxpool.Pool
}

// NewSportsEquipment Connects to DB
func NewSportsEquipment() *SportsEquipmentConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &SportsEquipmentConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *SportsEquipmentConn) Create(e *entity.SportsEquipment) (int32, error) {
	var id int32
	query := `INSERT INTO sports_equipment 
				(name, quantity, condition, location, last_maintenance_date, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Quantity,
		e.Condition,
		e.Location,
		e.LastMaintenanceDate,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckSportsEquipment Checks if record exists in DB by name
func (con *SportsEquipmentConn) CheckSportsEquipment(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM sports_equipment WHERE name = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

func sportsEquipmentSelectQuery() string {
	return `SELECT
				id,
				name,
				quantity,
				condition,
				location,
				last_maintenance_date,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM sports_equipment WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *SportsEquipmentConn) List() ([]*entity.SportsEquipment, error) {
	var id pgtype.Int4
	var name pgtype.Text
	var quantity pgtype.Int4
	var condition pgtype.Text
	var location pgtype.Text
	var lastMaintenance pgtype.Timestamp
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.SportsEquipment

	query := sportsEquipmentSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.SportsEquipment{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&quantity,
			&condition,
			&location,
			&lastMaintenance,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.SportsEquipment{}, err
		}

		var lastMaintenancePtr *time.Time
		if lastMaintenance.Status == pgtype.Present {
			lastMaintenancePtr = &lastMaintenance.Time
		}

		item := &entity.SportsEquipment{
			ID:                  id.Int,
			Name:                name.String,
			Quantity:            quantity.Int,
			Condition:           condition.String,
			Location:            location.String,
			LastMaintenanceDate: lastMaintenancePtr,
			CreatedBy:           createdBy.Int,
			UpdatedBy:           updatedBy.Int,
			DeletedBy:           deletedBy.Int,
			CreatedAt:           createdAt.Time.Local(),
			UpdatedAt:           updatedAt.Time.Local(),
			DeletedAt:           deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *SportsEquipmentConn) Get(id int32) (*entity.SportsEquipment, error) {
	var name pgtype.Text
	var quantity pgtype.Int4
	var condition pgtype.Text
	var location pgtype.Text
	var lastMaintenance pgtype.Timestamp
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.SportsEquipment

	query := sportsEquipmentSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&name,
		&quantity,
		&condition,
		&location,
		&lastMaintenance,
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

	var lastMaintenancePtr *time.Time
	if lastMaintenance.Status == pgtype.Present {
		lastMaintenancePtr = &lastMaintenance.Time
	}

	item = &entity.SportsEquipment{
		ID:                  id,
		Name:                name.String,
		Quantity:            quantity.Int,
		Condition:           condition.String,
		Location:            location.String,
		LastMaintenanceDate: lastMaintenancePtr,
		CreatedBy:           createdBy.Int,
		UpdatedBy:           updatedBy.Int,
		DeletedBy:           deletedBy.Int,
		CreatedAt:           createdAt.Time.Local(),
		UpdatedAt:           updatedAt.Time.Local(),
		DeletedAt:           deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *SportsEquipmentConn) Update(e *entity.SportsEquipment) (int32, error) {
	query := `UPDATE sports_equipment SET 
				name = $1,
				quantity = $2,
				condition = $3,
				location = $4,
				last_maintenance_date = $5,
				updated_by = $6,
				updated_at = $7
			  WHERE id = $8`
	_, err := con.conn.Exec(context.Background(), query,
		e.Name,
		e.Quantity,
		e.Condition,
		e.Location,
		e.LastMaintenanceDate,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *SportsEquipmentConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE sports_equipment SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *SportsEquipmentConn) HardDelete(id int32) error {
	query := "DELETE FROM sports_equipment WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
