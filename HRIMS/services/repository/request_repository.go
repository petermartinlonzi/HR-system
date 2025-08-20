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

// RequestConn Initializes connection to DB
type RequestConn struct {
	conn *pgxpool.Pool
}

// NewRequest Connects to DB
func NewRequest() *RequestConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &RequestConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *RequestConn) Create(e *entity.Request) (int32, error) {
	var id int32
	query := `INSERT INTO request 
				(officer_id, title, content, status, submitted_at, position_id, department_id, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.OfficerID,
		e.Title,
		e.Content,
		e.Status,
		e.SubmittedAt,
		e.PositionID,
		e.DepartmentID,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckRequest Checks if record exists in DB by title
func (con *RequestConn) CheckRequest(title string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM request WHERE title = $1)"
	err := con.conn.QueryRow(context.Background(), query, title).Scan(&exists)
	return exists, err
}

func requestSelectQuery() string {
	return `SELECT
				id,
				officer_id,
				title,
				content,
				status,
				submitted_at,
				position_id,
				department_id,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM request WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *RequestConn) List() ([]*entity.Request, error) {
	var id pgtype.Int4
	var officerID pgtype.Int4
	var title pgtype.Text
	var content pgtype.Text
	var status pgtype.Text
	var submittedAt pgtype.Timestamp
	var positionID pgtype.Int4
	var departmentID pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Request

	query := requestSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Request{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&officerID,
			&title,
			&content,
			&status,
			&submittedAt,
			&positionID,
			&departmentID,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Request{}, err
		}

		var posIDPtr *int32
		if positionID.Status == pgtype.Present {
			posIDPtr = &positionID.Int
		}
		var deptIDPtr *int32
		if departmentID.Status == pgtype.Present {
			deptIDPtr = &departmentID.Int
		}

		item := &entity.Request{
			ID:           id.Int,
			OfficerID:    officerID.Int,
			Title:        title.String,
			Content:      content.String,
			Status:       status.String,
			SubmittedAt:  submittedAt.Time.Local(),
			PositionID:   posIDPtr,
			DepartmentID: deptIDPtr,
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
func (con *RequestConn) Get(id int32) (*entity.Request, error) {
	var officerID pgtype.Int4
	var title pgtype.Text
	var content pgtype.Text
	var status pgtype.Text
	var submittedAt pgtype.Timestamp
	var positionID pgtype.Int4
	var departmentID pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Request

	query := requestSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&officerID,
		&title,
		&content,
		&status,
		&submittedAt,
		&positionID,
		&departmentID,
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

	var posIDPtr *int32
	if positionID.Status == pgtype.Present {
		posIDPtr = &positionID.Int
	}
	var deptIDPtr *int32
	if departmentID.Status == pgtype.Present {
		deptIDPtr = &departmentID.Int
	}

	item = &entity.Request{
		ID:           id,
		OfficerID:    officerID.Int,
		Title:        title.String,
		Content:      content.String,
		Status:       status.String,
		SubmittedAt:  submittedAt.Time.Local(),
		PositionID:   posIDPtr,
		DepartmentID: deptIDPtr,
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
func (con *RequestConn) Update(e *entity.Request) (int32, error) {
	query := `UPDATE request SET 
				officer_id = $1,
				title = $2, 
				content = $3,
				status = $4,
				position_id = $5,
				department_id = $6,
				updated_by = $7, 
				updated_at = $8 
			  WHERE id = $9`
	_, err := con.conn.Exec(context.Background(), query,
		e.OfficerID,
		e.Title,
		e.Content,
		e.Status,
		e.PositionID,
		e.DepartmentID,
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
func (con *RequestConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE request SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *RequestConn) HardDelete(id int32) error {
	query := "DELETE FROM request WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
