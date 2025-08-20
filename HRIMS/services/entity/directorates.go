package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Directorate struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

func NewDirectorate(name, description string, createdBy int32) (*Directorate, error) {
	entity := &Directorate{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDirectorate()
	if err != nil {
		log.Errorf("error validating new Directorate entity %v", err)
		return &Directorate{}, err
	}

	return entity, nil
}

func (r *Directorate) ValidateNewDirectorate() error {
	if r.Name == "" {
		return errors.New("error validating Directorate entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Directorate entity, createdBy field required")
	}
	return nil
}

func (r *Directorate) ValidateUpdateDirectorate() error {
	if r.ID <= 0 {
		return errors.New("error validating Directorate entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Directorate entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Directorate entity, updatedBy field required")
	}
	return nil
}
