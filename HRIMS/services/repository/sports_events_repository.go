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

// SportsEventConn Initializes connection to DB
type SportsEventConn struct {
	conn *pgxpool.Pool
}

// NewSportsEvent Connects to DB
func NewSportsEvent() *SportsEventConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &SportsEventConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *SportsEventConn) Create(e *entity.SportsEvent) (int32, error) {
	var id int32
	query := `INSERT INTO sports_event 
				(event_name, description, location, event_date, start_time, end_time, organized_by, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.EventName,
		e.Description,
		e.Location,
		e.EventDate,
		e.StartTime,
		e.EndTime,
		e.OrganizedBy,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckSportsEvent Checks if record exists in DB by event_name
func (con *SportsEventConn) CheckSportsEvent(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM sports_event WHERE event_name = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

func sportsEventSelectQuery() string {
	return `SELECT
				id,
				event_name,
				description,
				location,
				event_date,
				start_time,
				end_time,
				organized_by,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM sports_event WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *SportsEventConn) List() ([]*entity.SportsEvent, error) {
	var id pgtype.Int4
	var eventName pgtype.Text
	var description pgtype.Text
	var location pgtype.Text
	var eventDate pgtype.Timestamp
	var startTime pgtype.Timestamp
	var endTime pgtype.Timestamp
	var organizedBy pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.SportsEvent

	query := sportsEventSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.SportsEvent{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&eventName,
			&description,
			&location,
			&eventDate,
			&startTime,
			&endTime,
			&organizedBy,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.SportsEvent{}, err
		}

		var startPtr, endPtr *time.Time
		if startTime.Status == pgtype.Present {
			startPtr = &startTime.Time
		}
		if endTime.Status == pgtype.Present {
			endPtr = &endTime.Time
		}

		item := &entity.SportsEvent{
			ID:          id.Int,
			EventName:   eventName.String,
			Description: description.String,
			Location:    location.String,
			EventDate:   eventDate.Time.Local(),
			StartTime:   startPtr,
			EndTime:     endPtr,
			OrganizedBy: organizedBy.Int,
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
func (con *SportsEventConn) Get(id int32) (*entity.SportsEvent, error) {
	var eventName pgtype.Text
	var description pgtype.Text
	var location pgtype.Text
	var eventDate pgtype.Timestamp
	var startTime pgtype.Timestamp
	var endTime pgtype.Timestamp
	var organizedBy pgtype.Int4
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.SportsEvent

	query := sportsEventSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&eventName,
		&description,
		&location,
		&eventDate,
		&startTime,
		&endTime,
		&organizedBy,
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

	var startPtr, endPtr *time.Time
	if startTime.Status == pgtype.Present {
		startPtr = &startTime.Time
	}
	if endTime.Status == pgtype.Present {
		endPtr = &endTime.Time
	}

	item = &entity.SportsEvent{
		ID:          id,
		EventName:   eventName.String,
		Description: description.String,
		Location:    location.String,
		EventDate:   eventDate.Time.Local(),
		StartTime:   startPtr,
		EndTime:     endPtr,
		OrganizedBy: organizedBy.Int,
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
func (con *SportsEventConn) Update(e *entity.SportsEvent) (int32, error) {
	query := `UPDATE sports_event SET 
				event_name = $1,
				description = $2,
				location = $3,
				event_date = $4,
				start_time = $5,
				end_time = $6,
				organized_by = $7,
				updated_by = $8,
				updated_at = $9
			  WHERE id = $10`
	_, err := con.conn.Exec(context.Background(), query,
		e.EventName,
		e.Description,
		e.Location,
		e.EventDate,
		e.StartTime,
		e.EndTime,
		e.OrganizedBy,
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
func (con *SportsEventConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE sports_event SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *SportsEventConn) HardDelete(id int32) error {
	query := "DELETE FROM sports_event WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
