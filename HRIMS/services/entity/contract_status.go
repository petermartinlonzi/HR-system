package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Constract_status struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewConstract_status(name, description string, createdBy int32) (*Constract_status, error) {
	entity := &Constract_status{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewConstract_status()
	if err != nil {
		log.Errorf("error validating new Constract_status entity %v", err)
		return &Constract_status{}, err
	}

	return entity, nil
}

func (r *Constract_status) ValidateNewConstract_status() error {
	if r.Name == "" {
		return errors.New("error validating Constract_status entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Constract_status entity, createdBy field required")
	}
	return nil
}

func (r *Constract_status) ValidateUpdateConstract_status() error {
	if r.ID <= 0 {
		return errors.New("error validating Constract_status entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Constract_status entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Constract_status entity, updatedBy field required")
	}
	return nil
}
