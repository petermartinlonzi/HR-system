package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type StudentTrainingEnrollmentRepository struct {
	db *sql.DB
}

func NewStudentTrainingEnrollmentRepository(db *sql.DB) *StudentTrainingEnrollmentRepository {
	return &StudentTrainingEnrollmentRepository{db: db}
}

func (r *StudentTrainingEnrollmentRepository) Create(enrollment *entity.StudentTrainingEnrollment) error {
	query := `
		INSERT INTO student_training_enrollments
		(student_id, training_program_id, status, created_by, enrolled_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id`
	return r.db.QueryRow(query,
		enrollment.StudentID,
		enrollment.TrainingProgramID,
		enrollment.Status,
		enrollment.CreatedBy,
	).Scan(&enrollment.ID)
}

func (r *StudentTrainingEnrollmentRepository) List() ([]*entity.StudentTrainingEnrollment, error) {
	query := `
		SELECT id, student_id, training_program_id, status, created_by, updated_by, deleted_by, enrolled_at, updated_at, deleted_at
		FROM student_training_enrollments
		WHERE deleted_at IS NULL
		ORDER BY enrolled_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching student enrollments: %v", err)
		return nil, err
	}
	defer rows.Close()

	var enrollments []*entity.StudentTrainingEnrollment
	for rows.Next() {
		en := &entity.StudentTrainingEnrollment{}
		err := rows.Scan(&en.ID, &en.StudentID, &en.TrainingProgramID, &en.Status, &en.CreatedBy, &en.UpdatedBy, &en.DeletedBy, &en.EnrolledAt, &en.UpdatedAt, &en.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, en)
	}
	return enrollments, nil
}

func (r *StudentTrainingEnrollmentRepository) Get(id int32) (*entity.StudentTrainingEnrollment, error) {
	query := `
		SELECT id, student_id, training_program_id, status, created_by, updated_by, deleted_by, enrolled_at, updated_at, deleted_at
		FROM student_training_enrollments
		WHERE id = $1 AND deleted_at IS NULL`
	en := &entity.StudentTrainingEnrollment{}
	err := r.db.QueryRow(query, id).Scan(&en.ID, &en.StudentID, &en.TrainingProgramID, &en.Status, &en.CreatedBy, &en.UpdatedBy, &en.DeletedBy, &en.EnrolledAt, &en.UpdatedAt, &en.DeletedAt)
	if err != nil {
		return nil, err
	}
	return en, nil
}

func (r *StudentTrainingEnrollmentRepository) Update(en *entity.StudentTrainingEnrollment) error {
	query := `
		UPDATE student_training_enrollments
		SET status = $1, updated_by = $2, updated_at = NOW()
		WHERE id = $3 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, en.Status, en.UpdatedBy, en.ID)
	return err
}

func (r *StudentTrainingEnrollmentRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE student_training_enrollments
		SET deleted_by = $1, deleted_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, deletedBy, id)
	return err
}

func (r *StudentTrainingEnrollmentRepository) HardDelete(id int32) error {
	_, err := r.db.Exec("DELETE FROM student_training_enrollments WHERE id=$1", id)
	return err
}
