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

// DirectorateConn Initializes connection to DB
type DirectorateConn struct {
	conn *pgxpool.Pool
}

// NewDirectorate Connects to DB
func NewDirectorate() *DirectorateConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &DirectorateConn{
		conn: conn,
	}
}

// Create Inserts new directorate record to DB
func (con *DirectorateConn) Create(e *entity.Directorate) (int32, error) {
	var id int32
	query := `INSERT INTO directorates (name, description, created_by, created_at)
	VALUES($1, $2, $3, $4) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Description,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// directorateSelectQuery returns base SELECT query
func directorateSelectQuery() string {
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
			FROM directorates WHERE deleted_at IS NULL`
}

// List Lists all directorate records
func (con *DirectorateConn) List() ([]*entity.Directorate, error) {
	var id, createdBy, updatedBy, deletedBy pgtype.Int4
	var name, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.Directorate

	query := directorateSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Directorate{}, err
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
			return []*entity.Directorate{}, err
		}

		item := &entity.Directorate{
			ID:          id.Int,
			Name:        name.String,
			Description: description.String,
			CreatedBy:   createdBy.Int,
			UpdatedBy:   updatedBy.Int,
			DeletedBy:   deletedBy.Int,
			CreatedAt:   createdAt.Time.Local(),
			UpdatedAt:   &updatedAt.Time,
			DeletedAt:   &deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single directorate record by ID
func (con *DirectorateConn) Get(id int32) (*entity.Directorate, error) {
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var name, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.Directorate

	query := directorateSelectQuery() + ` AND id = $1`
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

	item = &entity.Directorate{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		CreatedBy:   createdBy.Int,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   &updatedAt.Time,
		DeletedAt:   &deletedAt.Time,
	}
	return item, err
}

// Update Updates single directorate record
func (con *DirectorateConn) Update(e *entity.Directorate) (int32, error) {
	query := `UPDATE directorates SET
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
	return e.ID, nil
}

// SoftDelete Softly delete directorate record
func (con *DirectorateConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE directorates SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete directorate record
func (con *DirectorateConn) HardDelete(id int32) error {
	query := "DELETE FROM directorates WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
