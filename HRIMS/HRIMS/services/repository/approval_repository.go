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

// ApprovalConn initializes DB connection for approval
type ApprovalConn struct {
	conn *pgxpool.Pool
}

// NewApproval connects to DB
func NewApproval() *ApprovalConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &ApprovalConn{
		conn: conn,
	}
}

// Create inserts a new approval record
func (con *ApprovalConn) Create(e *entity.Approval) (int32, error) {
	var id int32
	query := `INSERT INTO approval 
		(request_id, approved_by, status, comment, created_by, created_at) 
		VALUES($1, $2, $3, $4, $5, $6) 
		RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.RequestID,
		e.ApprovedBy,
		e.Status,
		e.Comment,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// approvalSelectQuery builds common SELECT query
func approvalSelectQuery() string {
	return `SELECT
				id,
				request_id,
				approved_by,
				status,
				comment,
				approved_at,
				created_by,
				created_at,
				updated_at,
				deleted_at
			FROM approval WHERE deleted_at IS NULL`
}

// List returns all approvals
func (con *ApprovalConn) List() ([]*entity.Approval, error) {
	var id pgtype.Int4
	var requestID pgtype.Int4
	var approvedBy pgtype.Int4
	var status pgtype.Text
	var comment pgtype.Text
	var approvedAt pgtype.Timestamp
	var createdBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Approval

	query := approvalSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Approval{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&requestID,
			&approvedBy,
			&status,
			&comment,
			&approvedAt,
			&createdBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Approval{}, err
		}
		item := &entity.Approval{
			ID:         id.Int,
			RequestID:  requestID.Int,
			ApprovedBy: approvedBy.Int,
			Status:     status.String,
			Comment:    comment.String,
			ApprovedAt: approvedAt.Time,
			CreatedBy:  createdBy.Int,
			CreatedAt:  createdAt.Time.Local(),
			UpdatedAt:  updatedAt.Time.Local(),
			DeletedAt:  deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get returns single approval by ID
func (con *ApprovalConn) Get(id int32) (*entity.Approval, error) {
	var requestID pgtype.Int4
	var approvedBy pgtype.Int4
	var status pgtype.Text
	var comment pgtype.Text
	var approvedAt pgtype.Timestamp
	var createdBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Approval

	query := approvalSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&requestID,
		&approvedBy,
		&status,
		&comment,
		&approvedAt,
		&createdBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return item, err
	}

	item = &entity.Approval{
		ID:         id,
		RequestID:  requestID.Int,
		ApprovedBy: approvedBy.Int,
		Status:     status.String,
		Comment:    comment.String,
		ApprovedAt: approvedAt.Time,
		CreatedBy:  createdBy.Int,
		CreatedAt:  createdAt.Time.Local(),
		UpdatedAt:  updatedAt.Time.Local(),
		DeletedAt:  deletedAt.Time.Local(),
	}
	return item, err
}

// Update updates an approval
func (con *ApprovalConn) Update(e *entity.Approval) (int32, error) {
	query := `UPDATE approval SET 
				status = $1, 
				comment = $2, 
				updated_at = $3 
			  WHERE id = $4`
	_, err := con.conn.Exec(context.Background(), query,
		e.Status,
		e.Comment,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete marks an approval as deleted
func (con *ApprovalConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE approval SET deleted_at = $1 WHERE id = $2"
	_, err := con.conn.Exec(context.Background(), query, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete permanently deletes an approval
func (con *ApprovalConn) HardDelete(id int32) error {
	query := "DELETE FROM approval WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
