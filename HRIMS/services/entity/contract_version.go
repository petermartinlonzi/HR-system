package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Constract_version struct {
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

func NewConstract_version(name, description string, createdBy int32) (*Constract_version, error) {
	entity := &Constract_version{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewConstract_version()
	if err != nil {
		log.Errorf("error validating new Constract_version entity %v", err)
		return &Constract_version{}, err
	}

	return entity, nil
}

func (r *Constract_version) ValidateNewConstract_version() error {
	if r.Name == "" {
		return errors.New("error validating Constract_version entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Constract_version entity, createdBy field required")
	}
	return nil
}

func (r *Constract_version) ValidateUpdateConstract_version() error {
	if r.ID <= 0 {
		return errors.New("error validating Constract_version entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Constract_version entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Constract_version entity, updatedBy field required")
	}
	return nil
}
