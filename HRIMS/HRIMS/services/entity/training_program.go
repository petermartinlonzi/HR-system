package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type TrainingProgram struct {
	ID          int32
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	DepartmentID int32

	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewTrainingProgram(title, description string, departmentID, createdBy int32) (*TrainingProgram, error) {
	entity := &TrainingProgram{
		Title:        title,
		Description:  description,
		DepartmentID: departmentID,
		CreatedBy:    createdBy,
	}

	err := entity.ValidateNewTrainingProgram()
	if err != nil {
		log.Errorf("error validating new TrainingProgram entity %v", err)
		return &TrainingProgram{}, err
	}

	return entity, nil
}

func (r *TrainingProgram) ValidateNewTrainingProgram() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.DepartmentID <= 0 {
		return errors.New("departmentID is required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("createdBy is required")
	}
	return nil
}

func (r *TrainingProgram) ValidateUpdateTrainingProgram() error {
	if r.ID <= 0 {
		return errors.New("id is required for update")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("updatedBy is required")
	}
	return nil
}
