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

// OfficerConn Initializes connection to DB
type OfficerConn struct {
	conn *pgxpool.Pool
}

// NewOfficer Connects to DB
func NewOfficer() *OfficerConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &OfficerConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *OfficerConn) Create(e *entity.Officer) (int32, error) {
	var id int32
	query := `INSERT INTO officers (user_id, department_id, position, phone, designation, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.UserID,
		e.DepartmentID,
		e.Position,
		e.Phone,
		e.Designation,
		e.CreatedBy,
		time.Now().Local()).Scan(&id)
	return id, err
}

// CheckOfficer Checks if officer exists by user_id
func (con *OfficerConn) CheckOfficer(userID int32) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM officers WHERE user_id = $1)"
	err := con.conn.QueryRow(context.Background(), query, userID).Scan(&exists)
	return exists, err
}

func officerSelectQuery() string {
	return `SELECT
				id,
				user_id,
				department_id,
				position,
				phone,
				designation,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM officers WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *OfficerConn) List() ([]*entity.Officer, error) {
	var id, userID, departmentID, createdBy, updatedBy, deletedBy pgtype.Int4
	var position, phone, designation pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.Officer
	query := officerSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Officer{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id, &userID, &departmentID, &position, &phone, &designation,
			&createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
		); err != nil {
			return []*entity.Officer{}, err
		}
		item := &entity.Officer{
			ID:           id.Int,
			UserID:       userID.Int,
			DepartmentID: departmentID.Int,
			Position:     position.String,
			Phone:        phone.String,
			Designation:  designation.String,
			CreatedBy:    createdBy.Int,
			UpdatedBy:    updatedBy.Int,
			DeletedBy:    deletedBy.Int,
			CreatedAt:    createdAt.Time.Local(),
			UpdatedAt:    updatedAt.Time.Local(),
			DeletedAt:    deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *OfficerConn) Get(id int32) (*entity.Officer, error) {
	var userID, departmentID, createdBy, updatedBy, deletedBy pgtype.Int4
	var position, phone, designation pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.Officer
	query := officerSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id, &userID, &departmentID, &position, &phone, &designation,
		&createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return item, err
	}

	item = &entity.Officer{
		ID:           id,
		UserID:       userID.Int,
		DepartmentID: departmentID.Int,
		Position:     position.String,
		Phone:        phone.String,
		Designation:  designation.String,
		CreatedBy:    createdBy.Int,
		UpdatedBy:    updatedBy.Int,
		DeletedBy:    deletedBy.Int,
		CreatedAt:    createdAt.Time.Local(),
		UpdatedAt:    updatedAt.Time.Local(),
		DeletedAt:    deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *OfficerConn) Update(e *entity.Officer) (int32, error) {
	query := `UPDATE officers SET 
				position = $1, 
				phone = $2, 
				designation = $3, 
				updated_by = $4, 
				updated_at = $5 WHERE id = $6`
	_, err := con.conn.Exec(context.Background(), query,
		e.Position,
		e.Phone,
		e.Designation,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID)
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *OfficerConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE officers SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete single record by ID
func (con *OfficerConn) HardDelete(id int32) error {
	query := "DELETE FROM officers WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
