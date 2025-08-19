package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type StudentTrainingEnrollment struct {
	ID               int32
	StudentID        int32
	TrainingProgramID int32
	EnrolledAt       time.Time
	Status           string

	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewStudentTrainingEnrollment(studentID, trainingProgramID, createdBy int32, status string) (*StudentTrainingEnrollment, error) {
	entity := &StudentTrainingEnrollment{
		StudentID:        studentID,
		TrainingProgramID: trainingProgramID,
		Status:           status,
		CreatedBy:        createdBy,
	}

	err := entity.ValidateNewStudentTrainingEnrollment()
	if err != nil {
		log.Errorf("error validating new StudentTrainingEnrollment entity %v", err)
		return &StudentTrainingEnrollment{}, err
	}

	return entity, nil
}

func (r *StudentTrainingEnrollment) ValidateNewStudentTrainingEnrollment() error {
	if r.StudentID <= 0 {
		return errors.New("studentID is required")
	}
	if r.TrainingProgramID <= 0 {
		return errors.New("trainingProgramID is required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("createdBy is required")
	}
	return nil
}

func (r *StudentTrainingEnrollment) ValidateUpdateStudentTrainingEnrollment() error {
	if r.ID <= 0 {
		return errors.New("id is required for update")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("updatedBy is required")
	}
	return nil
}
