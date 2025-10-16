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

// NotificationConn Initializes connection to DB
type NotificationConn struct {
	conn *pgxpool.Pool
}

// NewNotification Connects to DB
func NewNotification() *NotificationConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &NotificationConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *NotificationConn) Create(e *entity.Notification) (int32, error) {
	var id int32
	query := `INSERT INTO notification (user_id, message, is_read, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.UserID,
		e.Message,
		e.IsRead,
		e.CreatedBy,
		time.Now().Local()).Scan(&id)
	return id, err
}

// CheckNotification Checks if notification exists for user_id
func (con *NotificationConn) CheckNotification(userID int32) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM notification WHERE user_id = $1)"
	err := con.conn.QueryRow(context.Background(), query, userID).Scan(&exists)
	return exists, err
}

func notificationSelectQuery() string {
	return `SELECT
				id,
				user_id,
				message,
				is_read,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM notification WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *NotificationConn) List() ([]*entity.Notification, error) {
	var id, userID, createdBy, updatedBy, deletedBy pgtype.Int4
	var message pgtype.Text
	var isRead pgtype.Bool
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.Notification
	query := notificationSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Notification{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id, &userID, &message, &isRead,
			&createdBy, &updatedBy, &deletedBy,
			&createdAt, &updatedAt, &deletedAt,
		); err != nil {
			return []*entity.Notification{}, err
		}
		item := &entity.Notification{
			ID:        id.Int,
			UserID:    userID.Int,
			Message:   message.String,
			IsRead:    isRead.Bool,
			CreatedBy: createdBy.Int,
			UpdatedBy: updatedBy.Int,
			DeletedBy: deletedBy.Int,
			CreatedAt: createdAt.Time.Local(),
			UpdatedAt: updatedAt.Time.Local(),
			DeletedAt: deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *NotificationConn) Get(id int32) (*entity.Notification, error) {
	var userID, createdBy, updatedBy, deletedBy pgtype.Int4
	var message pgtype.Text
	var isRead pgtype.Bool
	var createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.Notification
	query := notificationSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id, &userID, &message, &isRead,
		&createdBy, &updatedBy, &deletedBy,
		&createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return item, err
	}

	item = &entity.Notification{
		ID:        id,
		UserID:    userID.Int,
		Message:   message.String,
		IsRead:    isRead.Bool,
		CreatedBy: createdBy.Int,
		UpdatedBy: updatedBy.Int,
		DeletedBy: deletedBy.Int,
		CreatedAt: createdAt.Time.Local(),
		UpdatedAt: updatedAt.Time.Local(),
		DeletedAt: deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *NotificationConn) Update(e *entity.Notification) (int32, error) {
	query := `UPDATE notification SET 
				message=$1, is_read=$2, updated_by=$3, updated_at=$4 WHERE id=$5`
	_, err := con.conn.Exec(context.Background(), query,
		e.Message, e.IsRead, e.UpdatedBy, time.Now().Local(), e.ID)
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *NotificationConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE notification SET deleted_by=$1, deleted_at=$2 WHERE id=$3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete single record by ID
func (con *NotificationConn) HardDelete(id int32) error {
	query := "DELETE FROM notification WHERE id=$1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
