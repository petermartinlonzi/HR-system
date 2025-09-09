package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type TrainingProgramRepository struct {
	db *sql.DB
}

func NewTrainingProgramRepository(db *sql.DB) *TrainingProgramRepository {
	return &TrainingProgramRepository{db: db}
}

func (r *TrainingProgramRepository) Create(tp *entity.TrainingProgram) error {
	query := `
		INSERT INTO training_programs
		(title, description, department_id, start_date, end_date, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id`
	return r.db.QueryRow(query, tp.Title, tp.Description, tp.DepartmentID, tp.StartDate, tp.EndDate, tp.CreatedBy).Scan(&tp.ID)
}

func (r *TrainingProgramRepository) List() ([]*entity.TrainingProgram, error) {
	query := `
		SELECT id, title, description, start_date, end_date, department_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM training_programs
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching training programs: %v", err)
		return nil, err
	}
	defer rows.Close()

	var programs []*entity.TrainingProgram
	for rows.Next() {
		tp := &entity.TrainingProgram{}
		err := rows.Scan(&tp.ID, &tp.Title, &tp.Description, &tp.StartDate, &tp.EndDate, &tp.DepartmentID, &tp.CreatedBy, &tp.UpdatedBy, &tp.DeletedBy, &tp.CreatedAt, &tp.UpdatedAt, &tp.DeletedAt)
		if err != nil {
			return nil, err
		}
		programs = append(programs, tp)
	}
	return programs, nil
}

func (r *TrainingProgramRepository) Get(id int32) (*entity.TrainingProgram, error) {
	query := `
		SELECT id, title, description, start_date, end_date, department_id, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM training_programs
		WHERE id = $1 AND deleted_at IS NULL`
	tp := &entity.TrainingProgram{}
	err := r.db.QueryRow(query, id).Scan(&tp.ID, &tp.Title, &tp.Description, &tp.StartDate, &tp.EndDate, &tp.DepartmentID, &tp.CreatedBy, &tp.UpdatedBy, &tp.DeletedBy, &tp.CreatedAt, &tp.UpdatedAt, &tp.DeletedAt)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func (r *TrainingProgramRepository) Update(tp *entity.TrainingProgram) error {
	query := `
		UPDATE training_programs
		SET title = $1, description = $2, start_date = $3, end_date = $4, updated_by = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, tp.Title, tp.Description, tp.StartDate, tp.EndDate, tp.UpdatedBy, tp.ID)
	return err
}

func (r *TrainingProgramRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE training_programs
		SET deleted_by = $1, deleted_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, deletedBy, id)
	return err
}

func (r *TrainingProgramRepository) HardDelete(id int32) error {
	_, err := r.db.Exec("DELETE FROM training_programs WHERE id=$1", id)
	return err
}
