package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Drivers struct {
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

func NewDrivers(name, description string, createdBy int32) (*Drivers, error) {
	entity := &Drivers{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDrivers()
	if err != nil {
		log.Errorf("error validating new Drivers entity %v", err)
		return &Drivers{}, err
	}

	return entity, nil
}

func (r *Drivers) ValidateNewDrivers() error {
	if r.Name == "" {
		return errors.New("error validating Drivers entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Drivers entity, createdBy field required")
	}
	return nil
}

func (r *Drivers) ValidateUpdateDrivers() error {
	if r.ID <= 0 {
		return errors.New("error validating Drivers entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Drivers entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Drivers entity, updatedBy field required")
	}
	return nil
}
