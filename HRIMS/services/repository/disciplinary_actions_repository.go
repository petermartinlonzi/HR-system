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

// DisciplinaryConn Initializes connection to DB
type DisciplinaryConn struct {
	conn *pgxpool.Pool
}

// NewDisciplinaryAction Connects to DB
func NewDisciplinaryAction() *DisciplinaryConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &DisciplinaryConn{
		conn: conn,
	}
}

// Create Inserts new disciplinary record to DB
func (con *DisciplinaryConn) Create(e *entity.DisciplinaryAction) (int32, error) {
	var id int32
	query := `INSERT INTO disciplinary_actions 
	(employee_id, reported_by, incident_date, description, action_taken, action_date, duration, status, created_by, created_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.EmployeeID,
		e.ReportedBy,
		e.IncidentDate,
		e.Description,
		e.ActionTaken,
		e.ActionDate,
		e.Duration,
		e.Status,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// disciplinarySelectQuery returns base SELECT query
func disciplinarySelectQuery() string {
	return `SELECT
				id,
				employee_id,
				reported_by,
				incident_date,
				description,
				action_taken,
				action_date,
				duration,
				status,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM disciplinary_actions WHERE deleted_at IS NULL`
}

// List Lists all disciplinary records
func (con *DisciplinaryConn) List() ([]*entity.DisciplinaryAction, error) {
	var id, employeeID, reportedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var description, actionTaken, status pgtype.Text
	var incidentDate, actionDate, createdAt, updatedAt, deletedAt pgtype.Timestamp
	var duration pgtype.Int4

	var items []*entity.DisciplinaryAction

	query := disciplinarySelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.DisciplinaryAction{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&employeeID,
			&reportedBy,
			&incidentDate,
			&description,
			&actionTaken,
			&actionDate,
			&duration,
			&status,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.DisciplinaryAction{}, err
		}

		item := &entity.DisciplinaryAction{
			ID:          id.Int,
			EmployeeID:  employeeID.Int,
			ReportedBy:  reportedBy.Int,
			IncidentDate: incidentDate.Time.Local(),
			Description: description.String,
			ActionTaken: actionTaken.String,
			ActionDate:  &actionDate.Time,
			Duration:    &duration.Int,  // *int32 sasa
			Status:      status.String,
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

// Get Gets single disciplinary record by ID
func (con *DisciplinaryConn) Get(id int32) (*entity.DisciplinaryAction, error) {
	var employeeID, reportedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var description, actionTaken, status pgtype.Text
	var incidentDate, actionDate, createdAt, updatedAt, deletedAt pgtype.Timestamp
	var duration pgtype.Int4

	var item *entity.DisciplinaryAction

	query := disciplinarySelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&employeeID,
		&reportedBy,
		&incidentDate,
		&description,
		&actionTaken,
		&actionDate,
		&duration,
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

	item = &entity.DisciplinaryAction{
		ID:          id,
		EmployeeID:  employeeID.Int,
		ReportedBy:  reportedBy.Int,
		IncidentDate: incidentDate.Time.Local(),
		Description: description.String,
		ActionTaken: actionTaken.String,
		ActionDate:  &actionDate.Time,
		Duration:    &duration.Int,  // *int32 sasa
		Status:      status.String,
		CreatedBy:   createdBy.Int,
		UpdatedBy:   updatedBy.Int,
		DeletedBy:   deletedBy.Int,
		CreatedAt:   createdAt.Time.Local(),
		UpdatedAt:   &updatedAt.Time,
		DeletedAt:   &deletedAt.Time,
	}
	return item, err
}

// Update Updates single disciplinary record
func (con *DisciplinaryConn) Update(e *entity.DisciplinaryAction) (int32, error) {
	query := `UPDATE disciplinary_actions SET 
				employee_id = $1,
				reported_by = $2,
				incident_date = $3,
				description = $4,
				action_taken = $5,
				action_date = $6,
				duration = $7,
				status = $8,
				updated_by = $9,
				updated_at = $10
			WHERE id = $11`
	_, err := con.conn.Exec(context.Background(), query,
		e.EmployeeID,
		e.ReportedBy,
		e.IncidentDate,
		e.Description,
		e.ActionTaken,
		e.ActionDate,
		e.Duration,
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

// SoftDelete Softly delete disciplinary record
func (con *DisciplinaryConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE disciplinary_actions SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete disciplinary record
func (con *DisciplinaryConn) HardDelete(id int32) error {
	query := "DELETE FROM disciplinary_actions WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
