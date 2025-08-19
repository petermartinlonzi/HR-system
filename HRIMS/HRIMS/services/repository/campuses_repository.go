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

// CampusesConn initializes DB connection for campuses
type CampusesConn struct {
	conn *pgxpool.Pool
}

// NewCampusesConn connects to DB
func NewCampusesConn() *CampusesConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &CampusesConn{
		conn: conn,
	}
}

// Create inserts new campus record
func (con *CampusesConn) Create(e *entity.Campuses) (int32, error) {
	var id int32
	query := `INSERT INTO campuses
		(name, description, resvd1, resvd2, resvd3, resvd4, resvd5, created_by, created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.Name,
		e.Description,
		e.Resvd1,
		e.Resvd2,
		e.Resvd3,
		e.Resvd4,
		e.Resvd5,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// campusesSelectQuery builds SELECT query
func campusesSelectQuery() string {
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
			FROM campuses WHERE deleted_at IS NULL`
}

// List returns all campuses
func (con *CampusesConn) List() ([]*entity.Campuses, error) {
	var (
		id, createdBy, updatedBy, deletedBy pgtype.Int4
		name, description                  pgtype.Text
		createdAt, updatedAt, deletedAt     pgtype.Timestamp
	)

	var items []*entity.Campuses

	rows, err := con.conn.Query(context.Background(), campusesSelectQuery())
	if err != nil {
		return []*entity.Campuses{}, err
	}
	defer rows.Close()

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
			return []*entity.Campuses{}, err
		}
		item := &entity.Campuses{
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

// Get returns single campus by ID
func (con *CampusesConn) Get(id int32) (*entity.Campuses, error) {
	var (
		name, description                  pgtype.Text
		createdBy, updatedBy, deletedBy     pgtype.Int4
		createdAt, updatedAt, deletedAt     pgtype.Timestamp
	)

	var item *entity.Campuses

	query := campusesSelectQuery() + ` AND id = $1`
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

	item = &entity.Campuses{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   updatedAt.Time.Local(),
		DeletedAt:   deletedAt.Time.Local(),
	}

	return item, err
}

// Update updates single campus record
func (con *CampusesConn) Update(e *entity.Campuses) (int32, error) {
	query := `UPDATE campuses SET
				name=$1,
				description=$2,
				updated_by=$8,
				updated_at=$9
			  WHERE id=$10`
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

// SoftDelete marks a campus as deleted
func (con *CampusesConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE campuses SET deleted_by=$1, deleted_at=$2 WHERE id=$3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently deletes a campus
func (con *CampusesConn) HardDelete(id int32) error {
	query := "DELETE FROM campuses WHERE id=$1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
