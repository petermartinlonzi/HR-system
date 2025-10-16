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

// AuditLogsConn initializes DB connection for audit_logs
type AuditLogsConn struct {
	conn *pgxpool.Pool
}

// NewAuditLogs connects to DB
func NewAuditLogs() *AuditLogsConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &AuditLogsConn{
		conn: conn,
	}
}

// Create inserts a new audit log record
func (con *AuditLogsConn) Create(e *entity.AuditLog) (int32, error) {
	var id int32
	query := `INSERT INTO audit_logs
		(user_id, action, table_name, record_id, details, created_by, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.UserID,
		e.Action,
		e.TableName,
		e.RecordID,
		e.Details,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// auditLogsSelectQuery builds common SELECT query
func auditLogsSelectQuery() string {
	return `SELECT
				id,
				user_id,
				action,
				table_name,
				record_id,
				details,
				created_by,
				created_at,
				updated_at,
				deleted_at
			FROM audit_logs WHERE deleted_at IS NULL`
}

// List returns all audit logs
func (con *AuditLogsConn) List() ([]*entity.AuditLog, error) {
	var id pgtype.Int4
	var userID pgtype.Int4
	var action pgtype.Text
	var tableName pgtype.Text
	var recordID pgtype.Int4
	var details pgtype.Text
	var createdBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.AuditLog

	query := auditLogsSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.AuditLog{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&userID,
			&action,
			&tableName,
			&recordID,
			&details,
			&createdBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.AuditLog{}, err
		}
		item := &entity.AuditLog{
			ID:        id.Int,
			UserID:    userID.Int,
			Action:    action.String,
			TableName: tableName.String,
			RecordID:  recordID.Int,
			Details:   details.String,
			CreatedBy: createdBy.Int,
			CreatedAt: createdAt.Time.Local(),
			UpdatedAt: updatedAt.Time.Local(),
			DeletedAt: deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get returns single audit log by ID
func (con *AuditLogsConn) Get(id int32) (*entity.AuditLog, error) {
	var userID pgtype.Int4
	var action pgtype.Text
	var tableName pgtype.Text
	var recordID pgtype.Int4
	var details pgtype.Text
	var createdBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.AuditLog

	query := auditLogsSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&userID,
		&action,
		&tableName,
		&recordID,
		&details,
		&createdBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return item, err
	}

	item = &entity.AuditLog{
		ID:        id,
		UserID:    userID.Int,
		Action:    action.String,
		TableName: tableName.String,
		RecordID:  recordID.Int,
		Details:   details.String,
		CreatedBy: createdBy.Int,
		CreatedAt: createdAt.Time.Local(),
		UpdatedAt: updatedAt.Time.Local(),
		DeletedAt: deletedAt.Time.Local(),
	}
	return item, err
}

// Update updates an audit log
func (con *AuditLogsConn) Update(e *entity.AuditLog) (int32, error) {
	query := `UPDATE audit_logs SET 
				action = $1,
				table_name = $2,
				record_id = $3,
				details = $4,
				updated_at = $5
			  WHERE id = $6`
	_, err := con.conn.Exec(context.Background(), query,
		e.Action,
		e.TableName,
		e.RecordID,
		e.Details,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete marks an audit log as deleted
func (con *AuditLogsConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE audit_logs SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete permanently deletes an audit log
func (con *AuditLogsConn) HardDelete(id int32) error {
	query := "DELETE FROM audit_logs WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
