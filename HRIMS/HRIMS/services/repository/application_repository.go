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

// ApplicationConn initializes connection to DB
type ApplicationConn struct {
	conn *pgxpool.Pool
}

// NewApplication connects to DB
func NewApplication() *ApplicationConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &ApplicationConn{
		conn: conn,
	}
}

// Create inserts new record into application table
func (con *ApplicationConn) Create(e *entity.Application) (int32, error) {
	var id int32
	query := `INSERT INTO application 
		(applicant_id, job_id, resume, status, created_by, created_at) 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.ApplicantID,
		e.JobID,
		e.Resume,
		e.Status,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

func applicationSelectQuery() string {
	return `SELECT
				id,
				applicant_id,
				job_id,
				resume,
				status,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM application WHERE deleted_at IS NULL`
}

// List lists all applications
func (con *ApplicationConn) List() ([]*entity.Application, error) {
	var id pgtype.Int4
	var applicantID pgtype.Int4
	var jobID pgtype.Int4
	var resume pgtype.Text
	var status pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Application

	query := applicationSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Application{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&applicantID,
			&jobID,
			&resume,
			&status,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Application{}, err
		}

		item := &entity.Application{
			ID:          id.Int,
			ApplicantID: applicantID.Int,
			JobID:       jobID.Int,
			Resume:      resume.String,
			Status:      status.String,
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

// Get gets single application by ID
func (con *ApplicationConn) Get(id int32) (*entity.Application, error) {
	var applicantID pgtype.Int4
	var jobID pgtype.Int4
	var resume pgtype.Text
	var status pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Application

	query := applicationSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&applicantID,
		&jobID,
		&resume,
		&status,
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

	item = &entity.Application{
		ID:          id,
		ApplicantID: applicantID.Int,
		JobID:       jobID.Int,
		Resume:      resume.String,
		Status:      status.String,
		CreatedBy:   createdBy.Int,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   updatedAt.Time.Local(),
		DeletedAt:   deletedAt.Time.Local(),
	}
	return item, err
}

// Update updates application record
func (con *ApplicationConn) Update(e *entity.Application) (int32, error) {
	query := `UPDATE application SET 
				applicant_id = $1,
				job_id = $2,
				resume = $3,
				status = $4,
				updated_by = $5,
				updated_at = $6
			  WHERE id = $7`

	_, err := con.conn.Exec(context.Background(), query,
		e.ApplicantID,
		e.JobID,
		e.Resume,
		e.Status,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)

	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// SoftDelete marks record as deleted
func (con *ApplicationConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE application SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently deletes record
func (con *ApplicationConn) HardDelete(id int32) error {
	query := "DELETE FROM application WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
