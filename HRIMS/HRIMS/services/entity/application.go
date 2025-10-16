package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Application struct {
	ID          int32
	ApplicantID int32
	JobID       int32
	Resume      string
	Status      string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Constructor
func NewApplication(applicantID, jobID, createdBy int32, resume string) (*Application, error) {
	entity := &Application{
		ApplicantID: applicantID,
		JobID:       jobID,
		Resume:      resume,
		Status:      "Pending", // default status
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewApplication()
	if err != nil {
		log.Errorf("error validating new Application entity %v", err)
		return &Application{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create application
func (r *Application) ValidateNewApplication() error {
	if r.ApplicantID <= 0 {
		return errors.New("error validating Application entity, applicantID field required")
	}
	if r.JobID <= 0 {
		return errors.New("error validating Application entity, jobID field required")
	}
	if r.Resume == "" {
		return errors.New("error validating Application entity, resume field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Application entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update application
func (r *Application) ValidateUpdateApplication() error {
	if r.ID <= 0 {
		return errors.New("error validating Application entity, id field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Application entity, updatedBy field required")
	}
	if r.Status == "" {
		return errors.New("error validating Application entity, status field required")
	}
	return nil
}
