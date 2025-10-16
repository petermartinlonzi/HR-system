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

// SportsParticipantConn Initializes connection to DB
type SportsParticipantConn struct {
	conn *pgxpool.Pool
}

// NewSportsParticipant Connects to DB
func NewSportsParticipant() *SportsParticipantConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &SportsParticipantConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *SportsParticipantConn) Create(e *entity.SportsParticipant) (int32, error) {
	var id int32
	query := `INSERT INTO sports_participants 
				(event_id, employee_id, team_name, performance_notes, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.EventID,
		e.EmployeeID,
		e.TeamName,
		e.PerformanceNotes,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckSportsParticipant Checks if record exists in DB by event_id and employee_id
func (con *SportsParticipantConn) CheckSportsParticipant(eventID, employeeID int32) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM sports_participants WHERE event_id = $1 AND employee_id = $2)"
	err := con.conn.QueryRow(context.Background(), query, eventID, employeeID).Scan(&exists)
	return exists, err
}

func sportsParticipantSelectQuery() string {
	return `SELECT
				id,
				event_id,
				employee_id,
				team_name,
				performance_notes,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM sports_participants WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *SportsParticipantConn) List() ([]*entity.SportsParticipant, error) {
	var id pgtype.Int4
	var eventID pgtype.Int4
	var employeeID pgtype.Int4
	var teamName pgtype.Text
	var performanceNotes pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.SportsParticipant

	query := sportsParticipantSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.SportsParticipant{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&eventID,
			&employeeID,
			&teamName,
			&performanceNotes,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.SportsParticipant{}, err
		}

		item := &entity.SportsParticipant{
			ID:               id.Int,
			EventID:          eventID.Int,
			EmployeeID:       employeeID.Int,
			TeamName:         teamName.String,
			PerformanceNotes: performanceNotes.String,
			CreatedBy:        createdBy.Int,
			UpdatedBy:        updatedBy.Int,
			DeletedBy:        deletedBy.Int,
			CreatedAt:        createdAt.Time.Local(),
			UpdatedAt:        updatedAt.Time.Local(),
			DeletedAt:        deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *SportsParticipantConn) Get(id int32) (*entity.SportsParticipant, error) {
	var eventID pgtype.Int4
	var employeeID pgtype.Int4
	var teamName pgtype.Text
	var performanceNotes pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.SportsParticipant

	query := sportsParticipantSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&eventID,
		&employeeID,
		&teamName,
		&performanceNotes,
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

	item = &entity.SportsParticipant{
		ID:               id,
		EventID:          eventID.Int,
		EmployeeID:       employeeID.Int,
		TeamName:         teamName.String,
		PerformanceNotes: performanceNotes.String,
		CreatedBy:        createdBy.Int,
		UpdatedBy:        updatedBy.Int,
		DeletedBy:        deletedBy.Int,
		CreatedAt:        createdAt.Time.Local(),
		UpdatedAt:        updatedAt.Time.Local(),
		DeletedAt:        deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *SportsParticipantConn) Update(e *entity.SportsParticipant) (int32, error) {
	query := `UPDATE sports_participants SET 
				event_id = $1,
				employee_id = $2,
				team_name = $3,
				performance_notes = $4,
				updated_by = $5,
				updated_at = $6
			  WHERE id = $7`
	_, err := con.conn.Exec(context.Background(), query,
		e.EventID,
		e.EmployeeID,
		e.TeamName,
		e.PerformanceNotes,
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
func (con *SportsParticipantConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE sports_participants SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *SportsParticipantConn) HardDelete(id int32) error {
	query := "DELETE FROM sports_participants WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
