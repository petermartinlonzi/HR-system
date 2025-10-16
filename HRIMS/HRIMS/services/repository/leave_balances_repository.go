package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type LeaveBalanceRepository struct {
	db *sql.DB
}

// Constructor
func NewLeaveBalanceRepository(db *sql.DB) *LeaveBalanceRepository {
	return &LeaveBalanceRepository{db: db}
}

// Create a new leave balance record
func (r *LeaveBalanceRepository) Create(leaveBalance *entity.LeaveBalance) error {
	query := `
		INSERT INTO leave_balances 
		(user_id, leave_type, year, total_entitled, used_days, remaining_days, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW()) 
		RETURNING id`

	err := r.db.QueryRow(query,
		leaveBalance.UserID,
		leaveBalance.LeaveType,
		leaveBalance.Year,
		leaveBalance.TotalEntitled,
		leaveBalance.UsedDays,
		leaveBalance.RemainingDays,
		leaveBalance.CreatedBy,
	).Scan(&leaveBalance.ID)

	if err != nil {
		log.Errorf("error inserting leave balance: %v", err)
		return err
	}

	return nil
}

// List all active leave balances
func (r *LeaveBalanceRepository) List() ([]*entity.LeaveBalance, error) {
	query := `
		SELECT id, user_id, leave_type, year, total_entitled, used_days, remaining_days, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM leave_balances
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching leave balances: %v", err)
		return nil, err
	}
	defer rows.Close()

	var balances []*entity.LeaveBalance
	for rows.Next() {
		lb := &entity.LeaveBalance{}
		err := rows.Scan(
			&lb.ID,
			&lb.UserID,
			&lb.LeaveType,
			&lb.Year,
			&lb.TotalEntitled,
			&lb.UsedDays,
			&lb.RemainingDays,
			&lb.CreatedBy,
			&lb.UpdatedBy,
			&lb.DeletedBy,
			&lb.CreatedAt,
			&lb.UpdatedAt,
			&lb.DeletedAt,
		)
		if err != nil {
			log.Errorf("error scanning leave balance row: %v", err)
			return nil, err
		}
		balances = append(balances, lb)
	}

	return balances, nil
}

// Get a single leave balance by ID
func (r *LeaveBalanceRepository) Get(id int32) (*entity.LeaveBalance, error) {
	query := `
		SELECT id, user_id, leave_type, year, total_entitled, used_days, remaining_days, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM leave_balances
		WHERE id = $1 AND deleted_at IS NULL`

	lb := &entity.LeaveBalance{}
	err := r.db.QueryRow(query, id).Scan(
		&lb.ID,
		&lb.UserID,
		&lb.LeaveType,
		&lb.Year,
		&lb.TotalEntitled,
		&lb.UsedDays,
		&lb.RemainingDays,
		&lb.CreatedBy,
		&lb.UpdatedBy,
		&lb.DeletedBy,
		&lb.CreatedAt,
		&lb.UpdatedAt,
		&lb.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("error fetching leave balance by id: %v", err)
		return nil, err
	}

	return lb, nil
}

// Update leave balance details
func (r *LeaveBalanceRepository) Update(leaveBalance *entity.LeaveBalance) error {
	query := `
		UPDATE leave_balances
		SET leave_type = $1, total_entitled = $2, used_days = $3, remaining_days = $4, updated_by = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL`

	_, err := r.db.Exec(query,
		leaveBalance.LeaveType,
		leaveBalance.TotalEntitled,
		leaveBalance.UsedDays,
		leaveBalance.RemainingDays,
		leaveBalance.UpdatedBy,
		leaveBalance.ID,
	)

	if err != nil {
		log.Errorf("error updating leave balance: %v", err)
		return err
	}
	return nil
}

// Soft delete a leave balance (mark as deleted)
func (r *LeaveBalanceRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE leave_balances
		SET deleted_by = $1, deleted_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`

	_, err := r.db.Exec(query, deletedBy, id)
	if err != nil {
		log.Errorf("error soft deleting leave balance: %v", err)
		return err
	}
	return nil
}

// Hard delete a leave balance (permanently remove)
func (r *LeaveBalanceRepository) HardDelete(id int32) error {
	query := `DELETE FROM leave_balances WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Errorf("error hard deleting leave balance: %v", err)
		return err
	}
	return nil
}
