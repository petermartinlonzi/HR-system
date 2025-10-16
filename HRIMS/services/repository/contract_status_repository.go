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

// ContractStatusConn initializes connection to DB
type ContractStatusConn struct {
	conn *pgxpool.Pool
}

// NewContractStatus connects to DB
func NewContractStatus() *ContractStatusConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &ContractStatusConn{
		conn: conn,
	}
}

// Create inserts new ContractStatus record to DB
func (con *ContractStatusConn) Create(e *entity.ContractStatus) (int32, error) {
	var id int32
	query := `INSERT INTO contract_status
	(status_name, description, created_by, created_at)
	VALUES($1,$2,$3,$4) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.StatusName,
		e.Description,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// contractStatusSelectQuery returns base SELECT query
func contractStatusSelectQuery() string {
	return `SELECT
				id,
				status_name,
				description,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM contract_status WHERE deleted_at IS NULL`
}

// List lists all contract status records
func (con *ContractStatusConn) List() ([]*entity.ContractStatus, error) {
	var id, createdBy, updatedBy, deletedBy pgtype.Int4
	var statusName, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.ContractStatus

	query := contractStatusSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.ContractStatus{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&statusName,
			&description,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.ContractStatus{}, err
		}

		item := &entity.ContractStatus{
			ID:         id.Int,
			StatusName: statusName.String,
			Description: description.String,
			CreatedBy:  createdBy.Int,
			UpdatedBy:  updatedBy.Int,
			DeletedBy:  deletedBy.Int,
			CreatedAt:  createdAt.Time.Local(),
			UpdatedAt:  &updatedAt.Time,
			DeletedAt:  &deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get gets single contract status by ID
func (con *ContractStatusConn) Get(id int32) (*entity.ContractStatus, error) {
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var statusName, description pgtype.Text
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.ContractStatus

	query := contractStatusSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&statusName,
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

	item = &entity.ContractStatus{
		ID:         id,
		StatusName: statusName.String,
		Description: description.String,
		CreatedBy:  createdBy.Int,
		UpdatedBy:  updatedBy.Int,
		DeletedBy:  deletedBy.Int,
		CreatedAt:  createdAt.Time.Local(),
		UpdatedAt:  &updatedAt.Time,
		DeletedAt:  &deletedAt.Time,
	}
	return item, err
}

// Update updates single contract status record
func (con *ContractStatusConn) Update(e *entity.ContractStatus) (int32, error) {
	query := `UPDATE contract_status SET
				status_name = $1,
				description = $2,
				updated_by = $3,
				updated_at = $4
			WHERE id = $5`
	_, err := con.conn.Exec(context.Background(), query,
		e.StatusName,
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

// SoftDelete softly delete contract status record
func (con *ContractStatusConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE contract_status SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently delete contract status record
func (con *ContractStatusConn) HardDelete(id int32) error {
	query := "DELETE FROM contract_status WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
