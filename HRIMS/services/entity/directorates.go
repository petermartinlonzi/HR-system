package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Directorates struct {
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

func NewDirectorates(name, description string, createdBy int32) (*Directorates, error) {
	entity := &Directorates{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDirectorates()
	if err != nil {
		log.Errorf("error validating new Directorates entity %v", err)
		return &Directorates{}, err
	}

	return entity, nil
}

func (r *Directorates) ValidateNewDirectorates() error {
	if r.Name == "" {
		return errors.New("error validating Directorates entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Directorates entity, createdBy field required")
	}
	return nil
}

func (r *Directorates) ValidateUpdateDirectorates() error {
	if r.ID <= 0 {
		return errors.New("error validating Directorates entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Directorates entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Directorates entity, updatedBy field required")
	}
	return nil
}
