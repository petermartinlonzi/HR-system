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

// RoleConn Initializes connection to DB
type RoleConn struct {
	conn *pgxpool.Pool
}

// NewRole Connects to DB
func NewRole() *RoleConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &RoleConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *RoleConn) Create(e *entity.Role) (int32, error) {
	var id int32
	query := `INSERT INTO role 
				(name, description, created_by, created_at) 
			  VALUES($1,$2,$3,$4) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Description,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckRole Checks if record exists in DB by name
func (con *RoleConn) CheckRole(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM role WHERE name = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

func roleSelectQuery() string {
	return `SELECT
				id,
				name,
				description,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM role WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *RoleConn) List() ([]*entity.Role, error) {
	var id pgtype.Int4
	var name pgtype.Text
	var description pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Role

	query := roleSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Role{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&description,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Role{}, err
		}

		item := &entity.Role{
			ID:          id.Int,
			Name:        name.String,
			Description: description.String,
			CreatedBy:   createdBy.Int,
			UpdatedBy:   updatedBy.Int,
			DeletedBy:   deletedBy.Int,
			CreatedAt:   createdAt.Time.Local(),
			UpdatedAt:   updatedAt.Time.Local(),
			DeletedAt:   deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *RoleConn) Get(id int32) (*entity.Role, error) {
	var name pgtype.Text
	var description pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Role

	query := roleSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&name,
		&description,
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

	item = &entity.Role{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		CreatedBy:   createdBy.Int,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   updatedAt.Time.Local(),
		DeletedAt:   deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *RoleConn) Update(e *entity.Role) (int32, error) {
	query := `UPDATE role SET 
				name = $1, 
				description = $2, 
				updated_by = $3, 
				updated_at = $4 
			  WHERE id = $5`
	_, err := con.conn.Exec(context.Background(), query,
		e.Name,
		e.Description,
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
func (con *RoleConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE role SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *RoleConn) HardDelete(id int32) error {
	query := "DELETE FROM role WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
