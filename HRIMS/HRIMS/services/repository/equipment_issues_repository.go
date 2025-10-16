package repository

import (
	"database/sql"
	"errors"
	"time"
	"training-backend/services/entity"
	"training-backend/package/log"
)

// EquipmentIssuesRepository defines the repository structure
type EquipmentIssuesRepository struct {
	db *sql.DB
}

// NewEquipmentIssuesRepository creates a new repository instance
func NewEquipmentIssuesRepository(db *sql.DB) *EquipmentIssuesRepository {
	return &EquipmentIssuesRepository{db: db}
}

// Create inserts a new EquipmentIssue into the database
func (r *EquipmentIssuesRepository) Create(issue *entity.EquipmentIssue) (*entity.EquipmentIssue, error) {
	query := `
		INSERT INTO equipment_issues (equipment_id, issued_to, issue_date, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at;
	`

	err := r.db.QueryRow(query,
		issue.EquipmentID,
		issue.IssuedTo,
		issue.IssueDate,
		issue.CreatedBy,
		time.Now(),
	).Scan(&issue.ID, &issue.CreatedAt)

	if err != nil {
		log.Errorf("failed to create equipment issue: %v", err)
		return nil, err
	}

	return issue, nil
}

// GetByID retrieves an EquipmentIssue by its ID
func (r *EquipmentIssuesRepository) GetByID(id int32) (*entity.EquipmentIssue, error) {
	query := `
		SELECT id, equipment_id, issued_to, issue_date, return_date, returned_condition,
		       created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM equipment_issues
		WHERE id = $1 AND deleted_at IS NULL;
	`

	issue := &entity.EquipmentIssue{}
	err := r.db.QueryRow(query, id).Scan(
		&issue.ID,
		&issue.EquipmentID,
		&issue.IssuedTo,
		&issue.IssueDate,
		&issue.ReturnDate,
		&issue.ReturnedCondition,
		&issue.CreatedBy,
		&issue.UpdatedBy,
		&issue.DeletedBy,
		&issue.CreatedAt,
		&issue.UpdatedAt,
		&issue.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Errorf("failed to get equipment issue by id: %v", err)
		return nil, err
	}

	return issue, nil
}

// GetAll retrieves all equipment issues that are not deleted
func (r *EquipmentIssuesRepository) GetAll() ([]*entity.EquipmentIssue, error) {
	query := `
		SELECT id, equipment_id, issued_to, issue_date, return_date, returned_condition,
		       created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM equipment_issues
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("failed to fetch equipment issues: %v", err)
		return nil, err
	}
	defer rows.Close()

	var issues []*entity.EquipmentIssue
	for rows.Next() {
		issue := &entity.EquipmentIssue{}
		err := rows.Scan(
			&issue.ID,
			&issue.EquipmentID,
			&issue.IssuedTo,
			&issue.IssueDate,
			&issue.ReturnDate,
			&issue.ReturnedCondition,
			&issue.CreatedBy,
			&issue.UpdatedBy,
			&issue.DeletedBy,
			&issue.CreatedAt,
			&issue.UpdatedAt,
			&issue.DeletedAt,
		)
		if err != nil {
			log.Errorf("failed to scan equipment issue: %v", err)
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, nil
}

// Update updates an existing EquipmentIssue
func (r *EquipmentIssuesRepository) Update(issue *entity.EquipmentIssue) error {
	if err := issue.ValidateUpdateEquipmentIssue(); err != nil {
		return err
	}

	query := `
		UPDATE equipment_issues
		SET return_date = $1, returned_condition = $2, updated_by = $3, updated_at = $4
		WHERE id = $5 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query,
		issue.ReturnDate,
		issue.ReturnedCondition,
		issue.UpdatedBy,
		time.Now(),
		issue.ID,
	)
	if err != nil {
		log.Errorf("failed to update equipment issue: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("equipment issue not found or already deleted")
	}

	return nil
}

// Delete performs a soft delete on an EquipmentIssue
func (r *EquipmentIssuesRepository) Delete(id, deletedBy int32) error {
	query := `
		UPDATE equipment_issues
		SET deleted_by = $1, deleted_at = $2
		WHERE id = $3 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query, deletedBy, time.Now(), id)
	if err != nil {
		log.Errorf("failed to delete equipment issue: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("equipment issue not found or already deleted")
	}

	return nil
}
