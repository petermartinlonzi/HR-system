package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type InternshipApplication struct {
	ID          int32
	StudentID   int32
	DepartmentID int32
	Resume      string
	Status      string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	AppliedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewInternshipApplication(studentID, departmentID, createdBy int32, resume string) (*InternshipApplication, error) {
	entity := &InternshipApplication{
		StudentID:    studentID,
		DepartmentID: departmentID,
		Resume:       resume,
		Status:       "Pending",
		CreatedBy:    createdBy,
		AppliedAt:    time.Now(),
	}

	err := entity.ValidateNewInternshipApplication()
	if err != nil {
		log.Errorf("error validating new InternshipApplication entity %v", err)
		return &InternshipApplication{}, err
	}

	return entity, nil
}

func (r *InternshipApplication) ValidateNewInternshipApplication() error {
	if r.StudentID <= 0 {
		return errors.New("error validating InternshipApplication entity, studentID required")
	}
	if r.DepartmentID <= 0 {
		return errors.New("error validating InternshipApplication entity, departmentID required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating InternshipApplication entity, createdBy required")
	}
	return nil
}

func (r *InternshipApplication) ValidateUpdateInternshipApplication() error {
	if r.ID <= 0 {
		return errors.New("error validating InternshipApplication entity, id required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating InternshipApplication entity, updatedBy required")
	}
	return nil
}
