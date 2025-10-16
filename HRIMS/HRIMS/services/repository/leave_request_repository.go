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

// LeaveRequestConn Initializes connection to DB
type LeaveRequestConn struct {
	conn *pgxpool.Pool
}

// NewLeaveRequest Connects to DB
func NewLeaveRequest() *LeaveRequestConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &LeaveRequestConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *LeaveRequestConn) Create(e *entity.LeaveRequest) (int32, error) {
	var id int32
	query := `INSERT INTO leave_requests 
		(user_id, leave_type, reason, start_date, end_date, status, applied_at, reviewed_by, created_by, created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.UserID,
		e.LeaveType,
		e.Reason,
		e.StartDate,
		e.EndDate,
		e.Status,
		e.AppliedAt,
		e.ReviewedBy,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckLeaveRequest Checks if leave request exists for a user on specific date
func (con *LeaveRequestConn) CheckLeaveRequest(userID int32, startDate, endDate time.Time) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM leave_requests WHERE user_id=$1 AND start_date=$2 AND end_date=$3)"
	err := con.conn.QueryRow(context.Background(), query, userID, startDate, endDate).Scan(&exists)
	return exists, err
}

func leaveRequestSelectQuery() string {
	return `SELECT
				id,
				user_id,
				leave_type,
				reason,
				start_date,
				end_date,
				total_days,
				status,
				applied_at,
				reviewed_by,
				reviewed_at,
				comment,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM leave_requests WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *LeaveRequestConn) List() ([]*entity.LeaveRequest, error) {
	var id, userID, totalDays, reviewedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var leaveType, reason, status, comment pgtype.Text
	var startDate, endDate, appliedAt, reviewedAt, createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.LeaveRequest
	query := leaveRequestSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.LeaveRequest{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id, &userID, &leaveType, &reason, &startDate, &endDate, &totalDays, &status,
			&appliedAt, &reviewedBy, &reviewedAt, &comment,
			&createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
		); err != nil {
			return []*entity.LeaveRequest{}, err
		}

		item := &entity.LeaveRequest{
			ID:         id.Int,
			UserID:     userID.Int,
			LeaveType:  leaveType.String,
			Reason:     reason.String,
			StartDate:  startDate.Time.Local(),
			EndDate:    endDate.Time.Local(),
			TotalDays:  totalDays.Int,
			Status:     status.String,
			AppliedAt:  appliedAt.Time.Local(),
			ReviewedBy: reviewedBy.Int,
			ReviewedAt: reviewedAt.Time.Local(),
			Comment:    comment.String,
			CreatedBy:  createdBy.Int,
			UpdatedBy:  updatedBy.Int,
			DeletedBy:  deletedBy.Int,
			CreatedAt:  createdAt.Time.Local(),
			UpdatedAt:  updatedAt.Time.Local(),
			DeletedAt:  deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *LeaveRequestConn) Get(id int32) (*entity.LeaveRequest, error) {
	var userID, totalDays, reviewedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var leaveType, reason, status, comment pgtype.Text
	var startDate, endDate, appliedAt, reviewedAt, createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.LeaveRequest
	query := leaveRequestSelectQuery() + ` AND id=$1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id, &userID, &leaveType, &reason, &startDate, &endDate, &totalDays, &status,
		&appliedAt, &reviewedBy, &reviewedAt, &comment,
		&createdBy, &updatedBy, &deletedBy, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return item, err
	}

	item = &entity.LeaveRequest{
		ID:         id,
		UserID:     userID.Int,
		LeaveType:  leaveType.String,
		Reason:     reason.String,
		StartDate:  startDate.Time.Local(),
		EndDate:    endDate.Time.Local(),
		TotalDays:  totalDays.Int,
		Status:     status.String,
		AppliedAt:  appliedAt.Time.Local(),
		ReviewedBy: reviewedBy.Int,
		ReviewedAt: reviewedAt.Time.Local(),
		Comment:    comment.String,
		CreatedBy:  createdBy.Int,
		UpdatedBy:  updatedBy.Int,
		DeletedBy:  deletedBy.Int,
		CreatedAt:  createdAt.Time.Local(),
		UpdatedAt:  updatedAt.Time.Local(),
		DeletedAt:  deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *LeaveRequestConn) Update(e *entity.LeaveRequest) (int32, error) {
	query := `UPDATE leave_requests SET 
				leave_type=$1, reason=$2, start_date=$3, end_date=$4, status=$5, reviewed_by=$6, reviewed_at=$7, comment=$8, updated_by=$9, updated_at=$10
				WHERE id=$11`
	_, err := con.conn.Exec(context.Background(), query,
		e.LeaveType,
		e.Reason,
		e.StartDate,
		e.EndDate,
		e.Status,
		e.ReviewedBy,
		e.ReviewedAt,
		e.Comment,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *LeaveRequestConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE leave_requests SET deleted_by=$1, deleted_at=$2 WHERE id=$3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete Permanently delete single record by ID
func (con *LeaveRequestConn) HardDelete(id int32) error {
	query := "DELETE FROM leave_requests WHERE id=$1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
