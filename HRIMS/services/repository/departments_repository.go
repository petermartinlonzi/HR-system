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

// DepartmentConn Initializes connection to DB
type DepartmentConn struct {
	conn *pgxpool.Pool
}

// NewDepartment Connects to DB
func NewDepartment() *DepartmentConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &DepartmentConn{
		conn: conn,
	}
}

// Create Inserts new department record to DB
func (con *DepartmentConn) Create(e *entity.Department) (int32, error) {
	var id int32
	query := `INSERT INTO departments (name, description, directorate_id, created_by, created_at)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Description,
		e.DirectorateID,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// departmentSelectQuery returns base SELECT query
func departmentSelectQuery() string {
	return `SELECT
				id,
				name,
				description,
				directorate_id,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM departments WHERE deleted_at IS NULL`
}

// List Lists all department records
func (con *DepartmentConn) List() ([]*entity.Department, error) {
	var id, directorateID, createdBy, updatedBy, deletedBy pgtype.Int4
	var name, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.Department

	query := departmentSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Department{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&description,
			&directorateID,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Department{}, err
		}

		item := &entity.Department{
			ID:            id.Int,
			Name:          name.String,
			Description:   description.String,
			DirectorateID: directorateID.Int,
			CreatedBy:     createdBy.Int,
			UpdatedBy:     updatedBy.Int,
			DeletedBy:     deletedBy.Int,
			CreatedAt:     createdAt.Time.Local(),
			UpdatedAt:     &updatedAt.Time,
			DeletedAt:     &deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single department record by ID
func (con *DepartmentConn) Get(id int32) (*entity.Department, error) {
	var directorateID, createdBy, updatedBy, deletedBy pgtype.Int4
	var name, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.Department

	query := departmentSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&name,
		&description,
		&directorateID,
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

	item = &entity.Department{
		ID:            id,
		Name:          name.String,
		Description:   description.String,
		DirectorateID: directorateID.Int,
		CreatedBy:     createdBy.Int,
		UpdatedBy:     updatedBy.Int,
		DeletedBy:     deletedBy.Int,
		CreatedAt:     createdAt.Time.Local(),
		UpdatedAt:     &updatedAt.Time,
		DeletedAt:     &deletedAt.Time,
	}
	return item, err
}

// Update Updates single department record
func (con *DepartmentConn) Update(e *entity.Department) (int32, error) {
	query := `UPDATE departments SET
				name = $1,
				description = $2,
				directorate_id = $3,
				updated_by = $4,
				updated_at = $5
			WHERE id = $6`
	_, err := con.conn.Exec(context.Background(), query,
		e.Name,
		e.Description,
		e.DirectorateID,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// SoftDelete Softly delete department record
func (con *DepartmentConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE departments SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete department record
func (con *DepartmentConn) HardDelete(id int32) error {
	query := "DELETE FROM departments WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
