package repository

import (
	"database/sql"
	"errors"
	"time"
	"training-backend/services/entity"
	"training-backend/package/log"
)

// JobAdvertsRepository handles CRUD operations for job adverts
type JobAdvertsRepository struct {
	db *sql.DB
}

// NewJobAdvertsRepository initializes the repository
func NewJobAdvertsRepository(db *sql.DB) *JobAdvertsRepository {
	return &JobAdvertsRepository{db: db}
}

// Create inserts a new job advert into the database
func (r *JobAdvertsRepository) Create(advert *entity.JobAdvert) (*entity.JobAdvert, error) {
	query := `
		INSERT INTO job_adverts 
			(title, description, department_id, posted_by, deadline, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at;
	`

	err := r.db.QueryRow(query,
		advert.Title,
		advert.Description,
		advert.DepartmentID,
		advert.PostedBy,
		advert.Deadline,
		advert.CreatedBy,
		time.Now(),
	).Scan(&advert.ID, &advert.CreatedAt)

	if err != nil {
		log.Errorf("failed to create job advert: %v", err)
		return nil, err
	}

	return advert, nil
}

// GetByID fetches a job advert by ID
func (r *JobAdvertsRepository) GetByID(id int32) (*entity.JobAdvert, error) {
	query := `
		SELECT id, title, description, department_id, posted_by, deadline,
		       created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM job_adverts
		WHERE id = $1 AND deleted_at IS NULL;
	`

	advert := &entity.JobAdvert{}
	err := r.db.QueryRow(query, id).Scan(
		&advert.ID,
		&advert.Title,
		&advert.Description,
		&advert.DepartmentID,
		&advert.PostedBy,
		&advert.Deadline,
		&advert.CreatedBy,
		&advert.UpdatedBy,
		&advert.DeletedBy,
		&advert.CreatedAt,
		&advert.UpdatedAt,
		&advert.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Errorf("failed to fetch job advert by id: %v", err)
		return nil, err
	}

	return advert, nil
}

// GetAll retrieves all job adverts
func (r *JobAdvertsRepository) GetAll() ([]*entity.JobAdvert, error) {
	query := `
		SELECT id, title, description, department_id, posted_by, deadline,
		       created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM job_adverts
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("failed to fetch job adverts: %v", err)
		return nil, err
	}
	defer rows.Close()

	var adverts []*entity.JobAdvert
	for rows.Next() {
		advert := &entity.JobAdvert{}
		err := rows.Scan(
			&advert.ID,
			&advert.Title,
			&advert.Description,
			&advert.DepartmentID,
			&advert.PostedBy,
			&advert.Deadline,
			&advert.CreatedBy,
			&advert.UpdatedBy,
			&advert.DeletedBy,
			&advert.CreatedAt,
			&advert.UpdatedAt,
			&advert.DeletedAt,
		)
		if err != nil {
			log.Errorf("failed to scan job advert: %v", err)
			return nil, err
		}
		adverts = append(adverts, advert)
	}

	return adverts, nil
}

// Update modifies an existing job advert
func (r *JobAdvertsRepository) Update(advert *entity.JobAdvert) error {
	if err := advert.ValidateUpdateJobAdvert(); err != nil {
		return err
	}

	query := `
		UPDATE job_adverts
		SET title = $1, description = $2, department_id = $3, deadline = $4,
		    updated_by = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query,
		advert.Title,
		advert.Description,
		advert.DepartmentID,
		advert.Deadline,
		advert.UpdatedBy,
		time.Now(),
		advert.ID,
	)

	if err != nil {
		log.Errorf("failed to update job advert: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("job advert not found or already deleted")
	}

	return nil
}

// Delete performs a soft delete on a job advert
func (r *JobAdvertsRepository) Delete(id, deletedBy int32) error {
	query := `
		UPDATE job_adverts
		SET deleted_by = $1, deleted_at = $2
		WHERE id = $3 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query, deletedBy, time.Now(), id)
	if err != nil {
		log.Errorf("failed to delete job advert: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("job advert not found or already deleted")
	}

	return nil
}
