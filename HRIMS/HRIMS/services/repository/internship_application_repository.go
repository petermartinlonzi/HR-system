package repository

import (
	"database/sql"
	"errors"
	"time"
	"training-backend/services/entity"
	"training-backend/package/log"
)

// InternshipApplicationsRepository handles CRUD operations for internship applications
type InternshipApplicationsRepository struct {
	db *sql.DB
}

// NewInternshipApplicationsRepository initializes the repository
func NewInternshipApplicationsRepository(db *sql.DB) *InternshipApplicationsRepository {
	return &InternshipApplicationsRepository{db: db}
}

// Create inserts a new internship application into the database
func (r *InternshipApplicationsRepository) Create(application *entity.InternshipApplication) (*entity.InternshipApplication, error) {
	query := `
		INSERT INTO internship_applications 
			(student_id, department_id, resume, status, created_by, applied_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, applied_at;
	`

	err := r.db.QueryRow(query,
		application.StudentID,
		application.DepartmentID,
		application.Resume,
		application.Status,
		application.CreatedBy,
		application.AppliedAt,
	).Scan(&application.ID, &application.AppliedAt)

	if err != nil {
		log.Errorf("failed to create internship application: %v", err)
		return nil, err
	}

	return application, nil
}

// GetByID fetches a single internship application by ID
func (r *InternshipApplicationsRepository) GetByID(id int32) (*entity.InternshipApplication, error) {
	query := `
		SELECT id, student_id, department_id, resume, status, created_by, updated_by, deleted_by,
		       applied_at, updated_at, deleted_at
		FROM internship_applications
		WHERE id = $1 AND deleted_at IS NULL;
	`

	application := &entity.InternshipApplication{}
	err := r.db.QueryRow(query, id).Scan(
		&application.ID,
		&application.StudentID,
		&application.DepartmentID,
		&application.Resume,
		&application.Status,
		&application.CreatedBy,
		&application.UpdatedBy,
		&application.DeletedBy,
		&application.AppliedAt,
		&application.UpdatedAt,
		&application.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Errorf("failed to get internship application by id: %v", err)
		return nil, err
	}

	return application, nil
}

// GetAll retrieves all internship applications
func (r *InternshipApplicationsRepository) GetAll() ([]*entity.InternshipApplication, error) {
	query := `
		SELECT id, student_id, department_id, resume, status, created_by, updated_by, deleted_by,
		       applied_at, updated_at, deleted_at
		FROM internship_applications
		WHERE deleted_at IS NULL
		ORDER BY applied_at DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("failed to fetch internship applications: %v", err)
		return nil, err
	}
	defer rows.Close()

	var applications []*entity.InternshipApplication
	for rows.Next() {
		application := &entity.InternshipApplication{}
		err := rows.Scan(
			&application.ID,
			&application.StudentID,
			&application.DepartmentID,
			&application.Resume,
			&application.Status,
			&application.CreatedBy,
			&application.UpdatedBy,
			&application.DeletedBy,
			&application.AppliedAt,
			&application.UpdatedAt,
			&application.DeletedAt,
		)
		if err != nil {
			log.Errorf("failed to scan internship application: %v", err)
			return nil, err
		}
		applications = append(applications, application)
	}

	return applications, nil
}

// Update modifies an existing internship application
func (r *InternshipApplicationsRepository) Update(application *entity.InternshipApplication) error {
	if err := application.ValidateUpdateInternshipApplication(); err != nil {
		return err
	}

	query := `
		UPDATE internship_applications
		SET status = $1, resume = $2, updated_by = $3, updated_at = $4
		WHERE id = $5 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query,
		application.Status,
		application.Resume,
		application.UpdatedBy,
		time.Now(),
		application.ID,
	)

	if err != nil {
		log.Errorf("failed to update internship application: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("internship application not found or already deleted")
	}

	return nil
}

// Delete performs a soft delete on an internship application
func (r *InternshipApplicationsRepository) Delete(id, deletedBy int32) error {
	query := `
		UPDATE internship_applications
		SET deleted_by = $1, deleted_at = $2
		WHERE id = $3 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query, deletedBy, time.Now(), id)
	if err != nil {
		log.Errorf("failed to delete internship application: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("internship application not found or already deleted")
	}

	return nil
}
